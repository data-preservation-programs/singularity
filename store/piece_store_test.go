package store

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	blocks "github.com/ipfs/go-block-format"
	"github.com/multiformats/go-varint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestReadAt2(t *testing.T) {
	db := database.OpenInMemory()
	defer database.DropAll(db)

	car := model.Car{
		Header: []byte("car-Header"),
	}
	assert.NoError(t, db.Create(&car).Error)

	block1 := blocks.NewBlock([]byte("block-data"))
	varint1 := uint64(block1.Cid().ByteLen() + len(block1.RawData()))
	carBlock1 := model.CarBlock{
		CarID:          car.ID,
		CID:            block1.Cid().String(),
		CarOffset:      10,
		CarBlockLength: varint1 + uint64(varint.UvarintSize(varint1)),
		Varint:         varint1,
	}
	assert.NoError(t, db.Create(&carBlock1).Error)

	// Create an instance of the PieceReader
	pieceReader := NewPieceReader2(db, &car)
	mockStore := &mockBlockstore{}
	mockStore.On("Get", mock.Anything, block1.Cid()).Return(block1, nil)
	pieceReader.blockStore = mockStore

	// Test the ReadAt method
	buf := make([]byte, 10)
	n, err := pieceReader.ReadAt(buf, 0)
	assert.NoError(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, "car-Header", string(buf))

	buf = make([]byte, 10)
	n, err = pieceReader.ReadAt(buf, 5)
	assert.NoError(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "eader", string(buf[:n]))

	buf = make([]byte, 100)
	n, err = pieceReader.ReadAt(buf, 10)
	assert.NoError(t, err)
	assert.Equal(t, 45, n)
	assert.Equal(t, "block-data", string(buf[35:45]))
}
