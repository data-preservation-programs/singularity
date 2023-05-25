package daggen

import (
	"bytes"
	"context"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	bstore "github.com/ipfs/go-ipfs-blockstore"
	cbor "github.com/ipfs/go-ipld-cbor"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/ipld/go-car"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"io"
	"os"
	"path"
	"sort"
	"strings"
)

func init() {
	uio.DefaultShardWidth = 1024
}

type DagGen struct {
	DB           *gorm.DB
	SourceID     uint32
	DatasetID    uint32
	ChunkID      uint32
	TargetLength int64
	PieceSize    int64
	OutputDir    string
	offset       int64
	filePath     string
	header       []byte
	rootCID      cid.Cid
	calc         *commp.Calc
	writer       io.Writer
	closer       io.Closer
	carBlockIDs  []uint64
}

func (r *DagGen) Digest() error {
	if r.calc == nil {
		return nil
	}
	rawCommp, rawPieceSize, err := r.calc.Digest()
	if rawPieceSize < uint64(r.PieceSize) {
		rawCommp, err = commp.PadCommP(rawCommp, rawPieceSize, uint64(r.PieceSize))
		if err != nil {
			return errors.Wrap(err, "failed to pad commP")
		}

		rawPieceSize = uint64(r.PieceSize)
	}

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommp)
	if err != nil {
		return errors.Wrap(err, "failed to convert commp to cid")
	}

	if r.closer != nil {
		err = r.closer.Close()
		if err != nil {
			return errors.Wrap(err, "failed to close car file")
		}
	}
	carModel := model.Car{
		PieceCID:  commCid.Bytes(),
		PieceSize: int64(rawPieceSize),
		RootCID:   r.rootCID.Bytes(),
		FileSize:  r.offset,
		FilePath:  r.filePath,
		DatasetID: r.DatasetID,
		SourceID:  &r.SourceID,
		ChunkID:   &r.ChunkID,
		Header:    r.header,
	}

	err = r.DB.Create(&carModel).Error
	if err != nil {
		return errors.Wrap(err, "failed to create car")
	}

	err = r.DB.Model(&model.CarBlock{}).Where("id in ?", r.carBlockIDs).
		Update("car_id", carModel.ID).Error
	if err != nil {
		return errors.Wrap(err, "failed to update car blocks")
	}
	r.carBlockIDs = nil

	return nil
}

func (r *DagGen) Write(blk blocks.Block) error {
	if r.offset > r.TargetLength {
		err := r.Digest()
		if err != nil {
			return errors.Wrap(err, "failed to digest")
		}
		r.offset = 0
		r.calc = nil
		r.writer = nil
		r.closer = nil
	}

	if r.writer == nil {
		r.calc = &commp.Calc{}
		r.writer = r.calc
		if r.OutputDir != "" {
			filename := uuid.NewString() + ".car"
			r.filePath = path.Join(r.OutputDir, filename)
			file, err := os.Create(r.filePath)
			if err != nil {
				return errors.Wrap(err, "failed to create car file")
			}
			r.writer = io.MultiWriter(r.calc, file)
			r.closer = file
		}
	}

	if r.offset == 0 {
		r.rootCID = blk.Cid()
		header := car.CarHeader{
			Roots:   []cid.Cid{blk.Cid()},
			Version: 1,
		}
		headerBytes, err := cbor.DumpObject(&header)
		if err != nil {
			return errors.Wrap(err, "failed to marshal car header")
		}
		headerBytesVarint := varint.ToUvarint(uint64(len(headerBytes)))

		r.header = append(headerBytesVarint, headerBytes...)
		n, err := io.Copy(r.writer, bytes.NewReader(r.header))
		if err != nil {
			return errors.Wrap(err, "failed to write car header")
		}

		r.offset += n
	}

	written, err := pack.WriteCarBlock(r.writer, blk)
	if err != nil {
		return errors.Wrap(err, "failed to write car block")
	}
	carBlock := model.CarBlock{
		CID:            blk.Cid().Bytes(),
		CarOffset:      r.offset,
		CarBlockLength: int32(written),
		Varint:         varint.ToUvarint(uint64(len(blk.RawData()) + blk.Cid().ByteLen())),
		RawBlock:       blk.RawData(),
	}
	err = r.DB.Create(&carBlock).Error
	if err != nil {
		return errors.Wrap(err, "failed to create car block")
	}
	r.carBlockIDs = append(r.carBlockIDs, carBlock.ID)
	r.offset += written
	return nil
}

func (r *DagGen) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := r.DB.WithContext(ctx).Model(&model.Item{}).Where("source_id = ?", r.SourceID).Order("path asc").Rows()
	if err != nil {
		return errors.Wrap(err, "failed to query items")
	}
	defer rows.Close()
	stack := []*stackEntry{
		{
			Name:    "",
			Entries: make(map[string]*entry),
		},
	}
	for rows.Next() {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		var item model.Item
		err = r.DB.ScanRows(rows, &item)
		if err != nil {
			return errors.Wrap(err, "failed to scan item")
		}

		segments := strings.Split(item.Path, "/")
		fileName := segments[len(segments)-1]
		dirSegments := []string{""}
		dirSegments = append(dirSegments, segments[:len(segments)-1]...)
		// Exit the directory stack
		for i := 0; i < len(stack); i++ {
			if len(dirSegments) <= i || stack[i].Name != dirSegments[i] {
				var err error
				stack, err = r.exitDirectory(ctx, stack, i)
				if err != nil {
					cancel()
					return err
				}
				break
			}
		}
		// Enter the new directory stack
		for i := len(stack); i < len(dirSegments); i++ {
			stack = append(stack, &stackEntry{
				Name:    dirSegments[i],
				Entries: make(map[string]*entry),
			})
		}
		// Add the item to the stack
		if item.Offset == 0 && item.Size == item.Length {
			stack[len(stack)-1].Entries[fileName] = &entry{
				CID:  item.CID.ToCid(),
				Size: uint64(item.Size),
			}
		} else {
			newItemPart := itemPart{
				CID:    item.CID.ToCid(),
				Size:   item.Size,
				Offset: item.Offset,
				Length: item.Length,
			}
			if _, ok := stack[len(stack)-1].Entries[fileName]; !ok {
				stack[len(stack)-1].Entries[fileName] = &entry{
					ItemParts: []itemPart{newItemPart},
				}
			} else {
				stack[len(stack)-1].Entries[fileName].ItemParts = append(
					stack[len(stack)-1].Entries[fileName].ItemParts,
					newItemPart)
			}
		}
	}

	stack, err = r.exitDirectory(ctx, stack, 0)
	if err != nil {
		cancel()
		return err
	}
	return nil
}

func (r *DagGen) exitDirectory(ctx context.Context, stack []*stackEntry, length int) ([]*stackEntry, error) {
	for len(stack) > length {
		bs := bstore.NewBlockstore(datastore.NewMapDatastore())
		dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
		dir := uio.NewDirectory(dagServ)
		dir.SetCidBuilder(merkledag.V1CidPrefix())
		keys := make([]string, 0, len(stack[len(stack)-1].Entries))
		for k := range stack[len(stack)-1].Entries {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, name := range keys {
			e := stack[len(stack)-1].Entries[name]
			if e.CID.Defined() {
				node := NewDummyNode(e.Size, e.CID)
				err := dir.AddChild(ctx, name, node)
				if err != nil {
					return nil, errors.Wrap(err, "failed to add child to directory")
				}
			} else {
				var links []pack.Link
				for _, part := range e.ItemParts {
					links = append(links, pack.Link{
						Link: format.Link{
							Name: "",
							Size: uint64(part.Length),
							Cid:  part.CID,
						},
						ChunkSize: uint64(part.Length),
					})
				}
				blks, rootNode, _ := pack.AssembleItem(links)
				for _, blk := range blks {
					err := bs.Put(ctx, blk)
					if err != nil {
						return nil, errors.Wrap(err, "failed to put block")
					}
				}
				err := dagServ.Add(ctx, rootNode)
				if err != nil {
					return nil, errors.Wrap(err, "failed to add root node")
				}
				err = dir.AddChild(ctx, name, rootNode)
				if err != nil {
					return nil, errors.Wrap(err, "failed to add child to directory")
				}
			}
		}
		dirNode, err := dir.GetNode()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get directory node")
		}
		err = dagServ.Add(ctx, dirNode)
		if err != nil {
			return nil, errors.Wrap(err, "failed to add directory node")
		}
		bs.HashOnRead(false)
		allKeysChan, err := bs.AllKeysChan(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get all keys")
		}
		var allKeys []cid.Cid
		for key := range allKeysChan {
			allKeys = append(allKeys, key)
		}
		slices.SortFunc(allKeys, func(i, j cid.Cid) bool {
			return i.KeyString() < j.KeyString()
		})
		for _, key := range allKeys {
			blk, err := bs.Get(ctx, key)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get block")
			}
			err = r.Write(blk)
			if err != nil {
				return nil, errors.Wrap(err, "failed to write block")
			}
		}
		stack = stack[:len(stack)-1]
	}
	return stack, nil
}

type stackEntry struct {
	Name    string
	Entries map[string]*entry
}

type entry struct {
	CID       cid.Cid
	Size      uint64
	ItemParts []itemPart
}

type itemPart struct {
	CID    cid.Cid
	Size   int64
	Offset int64
	Length int64
}
