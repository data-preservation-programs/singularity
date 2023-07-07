package inspect

import (
	"fmt"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
)

type ItemPartDetail struct {
	ID                        uint64    `json:"id"`
	ItemID                    uint64    `json:"itemId"`
	Offset                    int64     `json:"offset"`
	Length                    int64     `json:"length"`
	ItemPartCID               model.CID `json:"itemPartCid"`
	SourceID                  uint32    `json:"sourceId"`
	Path                      string    `json:"path"`
	Hash                      string    `json:"hash"`
	Size                      int64     `json:"size"`
	LastModifiedTimestampNano int64     `json:"lastModified"`
	ItemCID                   model.CID `json:"itemCid"`
	DirectoryID               uint64    `json:"directoryId"`
}

var ChunkDetailCmd = &cli.Command{
	Name:      "chunkdetail",
	Usage:     "Get details about a specific chunk",
	ArgsUsage: "<chunk_id>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		result, err := inspect.GetSourceChunkDetailHandler(
			db,
			c.Args().Get(0),
		)
		if err != nil {
			return err
		}

		if c.Bool("json") {
			cliutil.PrintToConsole(result, true, nil)
			return nil
		}

		fmt.Println("Chunk:")
		cliutil.PrintToConsole(result, false, []string{"PackingWorkerID"})
		fmt.Println("Pieces:")
		cliutil.PrintToConsole(result.Cars, false, nil)
		fmt.Println("Item Parts:")
		cliutil.PrintToConsole(underscore.Map(result.ItemParts, func(i model.ItemPart) ItemPartDetail {
			return ItemPartDetail{
				ID:                        i.ID,
				ItemID:                    i.ItemID,
				Offset:                    i.Offset,
				Length:                    i.Length,
				ItemPartCID:               i.CID,
				SourceID:                  i.Item.SourceID,
				Path:                      i.Item.Path,
				Hash:                      i.Item.Hash,
				Size:                      i.Item.Size,
				LastModifiedTimestampNano: i.Item.LastModifiedTimestampNano,
				ItemCID:                   i.Item.CID,
				DirectoryID:               *i.Item.DirectoryID,
			}
		}), false, nil)
		return nil
	},
}
