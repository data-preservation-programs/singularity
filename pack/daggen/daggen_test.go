package daggen

import (
	"context"
	"fmt"
	"github.com/data-preservation-programs/singularity/model"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestDagGen(t *testing.T) {
	tempDir := t.TempDir()
	itemsChan := make(chan model.Item)
	c := cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes()
	cid.MustParse(c)
	go func() {
		itemsChan <- model.Item{
			CID: cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes(),
			Path: "a/b/c.txt",
			Size: 200,
			Offset: 0,
			Length: 100,
		}
		itemsChan <- model.Item{
			CID: cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes(),
			Path: "a/b/c.txt",
			Size: 200,
			Offset: 100,
			Length: 100,
		}
		itemsChan <- model.Item{
			CID: cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes(),
			Path: "a/b/d.txt",
			Size: 100,
			Length: 100,
		}
		itemsChan <- model.Item{
			CID: cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes(),
			Path: "a/e.txt",
			Size: 100,
			Length: 100,
		}
		itemsChan <- model.Item{
			CID: cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes(),
			Path: "a/f/g.txt",
			Size: 100,
			Length: 100,
		}
		itemsChan <- model.Item{
			CID: cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes(),
			Path: "a/f/h.txt",
			Size: 100,
			Length: 100,
		}
		itemsChan <- model.Item{
			CID: cid.MustParse("bafkreifiltefabyfkrw6zp7gsxqhcrqzchwebpafi32aezetxfhml2nk4m").Bytes(),
			Path: "i.txt",
			Size: 100,
			Length: 100,
		}
		for i := 0; i < 1000000; i++ {
			blk := blocks.NewBlock([]byte(strconv.Itoa(i)))
			itemsChan <- model.Item{
				CID: blk.Cid().Bytes(),
				Path: "z/" + strconv.Itoa(i) + ".txt",
				Size: 100,
				Length: 100,
			}
		}
		close(itemsChan)
	}()
	result, err := DagGen(context.Background(), nil, itemsChan, tempDir, 2<<30, 1<<30)
	result.Close()
	fmt.Println(result.Cars[0].PieceCID.ToCid().String())
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
