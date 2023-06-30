package handler

import (
	"context"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/fxamacker/cbor/v2"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func DownloadHandler(ctx context.Context,
	piece string,
	api string,
	meta map[string]string,
	outDir string,
	concurrency int,
) error {
	resolver := datasource.DefaultHandlerResolver{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api+"/piece/metadata/"+piece, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.Header.Add("Accept", "application/cbor")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to get metadata")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("failed to get metadata: %s", resp.Status)
	}

	var pieceMetadata service.PieceMetadata
	err = cbor.NewDecoder(resp.Body).Decode(&pieceMetadata)
	if err != nil {
		return errors.Wrap(err, "failed to decode metadata")
	}

	pieceMetadata.Source.Metadata = meta
	pieceReader, err := store.NewPieceReader(ctx, pieceMetadata.Car, pieceMetadata.Source, pieceMetadata.CarBlocks, pieceMetadata.Items, resolver)
	if err != nil {
		return errors.Wrap(err, "failed to create piece reader")
	}
	defer pieceReader.Close()

	return download(ctx, pieceReader, filepath.Join(outDir, piece+".car"), concurrency)
}

func download(ctx context.Context, reader *store.PieceReader, outPath string, concurrency int) error {
	size, err := reader.Seek(0, io.SeekEnd)
	if err != nil {
		return errors.New("failed to seek to end of piece")
	}
	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New("failed to seek to start of piece")
	}

	file, err := os.Create(outPath)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	partSize := size / int64(concurrency)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errChan := make(chan error, 1)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			start := int64(i) * partSize
			end := start + partSize

			// Adjust for the last part
			if i == concurrency-1 {
				end = size
			}

			// Clone the reader
			clonedReader := reader.Clone(ctx)
			defer clonedReader.Close()

			// Seek to the start position
			_, err := clonedReader.Seek(start, io.SeekStart)
			if err != nil {
				select {
				case <-ctx.Done():
				case errChan <- err:
				}
				cancel()
				return
			}

			// Read the part into a buffer
			reader := io.LimitReader(clonedReader, end-start)
			buffer := make([]byte, 4096)
			for {
				n, err := reader.Read(buffer)
				if err != nil && !errors.Is(err, io.EOF) {
					select {
					case <-ctx.Done():
					case errChan <- err:
					}
					cancel()
					return
				}
				if n == 0 {
					break
				}
				if _, err := file.WriteAt(buffer[:n], start); err != nil {
					select {
					case <-ctx.Done():
					case errChan <- err:
					}
					cancel()
					return
				}
				start += int64(n)
			}
		}(i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return file.Close()
	case err := <-errChan:
		file.Close()
		return err
	}
}
