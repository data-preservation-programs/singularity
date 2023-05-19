package handler

import (
	"context"
	"encoding/json"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/store"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

func DownloadHandler(piece string, api string, meta model.Metadata) error {
	resolver := datasource.DefaultHandlerResolver{}
	resp, err := http.Get(api + "/admin/api/piece/metadata/" + piece)
	if err != nil {
		return errors.Wrap(err, "failed to get metadata")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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
	pieceReader, err = pieceReader.MakeCopy(context.TODO(), 0)
	for i, _ := range pieceReader.Blocks {
		if itemBlock, ok := pieceReader.Blocks[i].(store.ItemBlock); ok {
			source := model.Source{}
			// TODO source.Type = model.Local
			source.Metadata = meta
			handler, err := resolver.Resolve(context.TODO(), source)
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
