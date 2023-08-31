package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/fxamacker/cbor/v2"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
)

// DownloadHandler fetches the metadata for a specified piece from a remote API and downloads the associated content.
// The content is then saved to a .car file in the specified directory.
//
// Parameters:
// - ctx: The context for the operation.
// - piece: The identifier of the content piece to be downloaded.
// - api: The base URL of the API from which metadata is to be fetched.
// - config: A map containing configuration settings for various storage types.
// - outDir: The directory where the downloaded content should be saved.
// - concurrency: The number of concurrent operations allowed during the download.
//
// Returns:
// - An error, if any occurred during the download process.
func DownloadHandler(ctx *cli.Context,
	piece string,
	api string,
	config map[string]string,
	outDir string,
	concurrency int,
) error {
	req, err := http.NewRequestWithContext(ctx.Context, http.MethodGet, api+"/piece/metadata/"+piece, nil)
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

	var pieceMetadata contentprovider.PieceMetadata
	err = cbor.NewDecoder(resp.Body).Decode(&pieceMetadata)
	if err != nil {
		return errors.Wrap(err, "failed to decode metadata")
	}

	for i, storage := range pieceMetadata.Storages {
		cfg := make(map[string]string)
		backend, ok := storagesystem.BackendMap[storage.Type]
		if !ok {
			return errors.Newf("storage type %s is not supported", storage.Type)
		}

		prefix := storage.Type + "-"
		provider := storage.Config["provider"]
		if provider == "" {
			provider = config[prefix+"provider"]
		}
		providerOptions, err := underscore.Find(backend.ProviderOptions, func(providerOption storagesystem.ProviderOptions) bool {
			return providerOption.Provider == provider
		})
		if err != nil {
			return errors.Newf("provider '%s' is not supported", provider)
		}

		for _, option := range providerOptions.Options {
			if option.Default != nil {
				cfg[option.Name] = fmt.Sprintf("%v", option.Default)
			}
		}

		for key, value := range storage.Config {
			cfg[key] = value
		}

		for key, value := range config {
			if strings.HasPrefix(key, prefix) {
				trimmed := strings.TrimPrefix(key, prefix)
				snake := strings.ReplaceAll(trimmed, "-", "_")
				cfg[snake] = value
			}
		}

		pieceMetadata.Storages[i].Config = cfg
	}

	pieceReader, err := store.NewPieceReader(ctx.Context, pieceMetadata.Car, pieceMetadata.Storages, pieceMetadata.CarBlocks, pieceMetadata.Files)
	if err != nil {
		return errors.Wrap(err, "failed to create piece reader")
	}
	defer pieceReader.Close()

	return download(ctx, pieceReader, filepath.Join(outDir, piece+".car"), concurrency)
}

// download concurrently fetches content from a given PieceReader, saving it to the specified output path.
// The content is divided into parts and downloaded concurrently based on the provided concurrency level.
//
// Parameters:
// - cctx: The context for the operation, allowing for cancellation.
// - reader: The PieceReader providing the content.
// - outPath: The path where the downloaded content should be saved.
// - concurrency: The number of concurrent download tasks.
//
// Returns:
// - An error, if any occurred during the download process.
func download(cctx *cli.Context, reader *store.PieceReader, outPath string, concurrency int) error {
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
		return errors.WithStack(err)
	}

	var wg sync.WaitGroup
	partSize := size / int64(concurrency)

	ctx, cancel := context.WithCancel(cctx.Context)
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
			clonedReader := reader.Clone()
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
			if !cctx.Bool("quiet") {
				_, _ = cctx.App.Writer.Write([]byte(fmt.Sprintf("[Thread %d] Downloading part %d - %d\n", i, end, start)))
			}
			for {
				if ctx.Err() != nil {
					return
				}
				n, err := reader.Read(buffer)
				if err != nil && !errors.Is(err, io.EOF) {
					select {
					case <-ctx.Done():
					case errChan <- err:
					}
					cancel()
					return
				}
				if errors.Is(err, io.EOF) {
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
			if !cctx.Bool("quiet") {
				_, _ = cctx.App.Writer.Write([]byte(fmt.Sprintf("[Thread %d] Completed\n", i)))
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
		if !cctx.Bool("quiet") {
			_, _ = cctx.App.Writer.Write([]byte("Download Complete\n"))
		}
		return file.Close()
	case err := <-errChan:
		file.Close()
		return errors.WithStack(err)
	}
}
