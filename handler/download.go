package handler

import (
	"context"
	"encoding/json"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

func DownloadHandler(ctx context.Context, piece string, api string, meta model.Metadata) error {
	resolver := datasource.DefaultHandlerResolver{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api+"/admin/api/piece/metadata/"+piece, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to get metadata")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("failed to get metadata: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read metadata")
	}

	pieceReader, err := UnmarshalPieceReader(body)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata")
	}
	pieceReader, err = pieceReader.MakeCopy(ctx, 0)
	if err != nil {
		return errors.Wrap(err, "failed to make copy")
	}
	for i := range pieceReader.Blocks {
		if itemBlock, ok := pieceReader.Blocks[i].(store.ItemBlock); ok {
			source := model.Source{}
			// TODO source.Type = model.Local
			source.Metadata = meta
			handler, err := resolver.Resolve(ctx, source)
			if err != nil {
				return errors.Wrap(err, "failed to get handler")
			}
			pieceReader.Blocks[i] = store.ItemBlock{
				PieceOffset:   itemBlock.PieceOffset,
				SourceHandler: handler,
				Item:          itemBlock.Item,
				Meta:          itemBlock.Meta,
			}
		}
	}
	_, err = io.Copy(os.Stdout, pieceReader)
	if err != nil {
		return errors.Wrap(err, "failed to copy data")
	}
	return nil
}

type blockType struct {
	Varint *[]byte `json:"varint"`
}

func UnmarshalPieceReader(data []byte) (*store.PieceReader, error) {
	var rawBlocks []json.RawMessage
	reader := store.PieceReader{
		Blocks: make([]store.PieceBlock, 0),
	}

	// Unmarshal into temporary struct to get the raw JSON blocks
	temp := struct {
		Blocks []json.RawMessage `json:"blocks"`
		*store.PieceReader
	}{
		PieceReader: &reader,
		Blocks:      rawBlocks,
	}

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return nil, err
	}

	for _, rawBlock := range temp.Blocks {
		var blockType blockType
		err := json.Unmarshal(rawBlock, &blockType)
		if err != nil {
			return nil, err
		}

		if blockType.Varint != nil {
			var b store.RawBlock
			err := json.Unmarshal(rawBlock, &b)
			if err != nil {
				return nil, err
			}
			reader.Blocks = append(reader.Blocks, b)
		} else {
			var b store.ItemBlock
			err := json.Unmarshal(rawBlock, &b)
			if err != nil {
				return nil, err
			}
			reader.Blocks = append(reader.Blocks, b)
		}
	}
	return &reader, nil
}
