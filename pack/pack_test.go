package pack

import (
	"bytes"
	"context"
	"crypto/rand"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/multiformats/go-varint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"testing"
)

func TestAssembleCar_LargeItems(t *testing.T) {
	ctx := context.Background()
	data := make([]byte, 10*1<<20)
	rand.Read(data)
	handler := new(MockReadHandler)
	handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(io.NopCloser(bytes.NewReader(data)), nil, nil)
	items := []model.ItemPart{
		{
			ID:     0,
			Offset: 0,
			Length: 10 * 1 << 20,
			Item: &model.Item{
				Size: 10 * 1 << 20,
			},
		},
		{
			ID:     1,
			Offset: 0,
			Length: 10 * 1 << 20,
			Item: &model.Item{
				Size: 10 * 1 << 20,
			},
		},
	}
	result, err := AssembleCar(ctx, handler, model.Dataset{}, items, "", 1<<30)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "", result.CarFilePath)
	assert.EqualValues(t, 10486756, result.CarFileSize)
	assert.EqualValues(t, 1<<30, result.PieceSize)
	assert.Len(t, result.Header, 59)
	assert.Len(t, result.ItemEncryptionStates, 0)
	assert.Len(t, result.CarBlocks, 11)
}

func TestAssembleCar_NoEncryption(t *testing.T) {
	ctx := context.Background()
	handler := new(MockReadHandler)
	handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(io.NopCloser(bytes.NewReader([]byte("hello"))), nil, nil)
	items := []model.ItemPart{
		{
			Offset: 1,
			Length: 4,
			Item: &model.Item{
				Size: 5,
			},
		},
	}
	result, err := AssembleCar(ctx, handler, model.Dataset{}, items, "", 1<<20)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "", result.CarFilePath)
	assert.EqualValues(t, 101, result.CarFileSize)
	assert.Equal(t, "baga6ea4seaqddc4kqdxmnglxhmrfxkx2cd7hl6kounifgewgb2hdthmeddfe4hy", result.PieceCID.String())
	assert.EqualValues(t, 1<<20, result.PieceSize)
	assert.Len(t, result.Header, 59)
	assert.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", result.RootCID.String())
	assert.Len(t, result.ItemEncryptionStates, 0)
	assert.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", result.ItemPartCIDs[0].String())
	assert.Len(t, result.CarBlocks, 1)
	assert.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", result.CarBlocks[0].CID.String())
	assert.EqualValues(t, 59, result.CarBlocks[0].CarOffset)
	v, _, _ := varint.FromUvarint(result.CarBlocks[0].Varint)
	assert.EqualValues(t, 41, v)
	assert.EqualValues(t, 1, result.CarBlocks[0].ItemOffset)
	assert.False(t, result.CarBlocks[0].ItemEncrypted)
}

func TestAssembleCar_WithEncryption(t *testing.T) {
	ctx := context.Background()
	handler := new(MockReadHandler)
	handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(io.NopCloser(bytes.NewReader([]byte("hello"))), nil, nil)
	items := []model.ItemPart{
		{
			Offset: 0,
			Length: 4,
			Item: &model.Item{
				Size: 5,
			},
		},
	}
	result, err := AssembleCar(ctx, handler, model.Dataset{
		EncryptionRecipients: []string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"},
	}, items, t.TempDir(), 1<<20)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.CarFilePath)
	assert.EqualValues(t, 281, result.CarFileSize)
	assert.EqualValues(t, 1<<20, result.PieceSize)
	assert.Len(t, result.Header, 59)
	assert.Len(t, result.ItemEncryptionStates, 1)
	assert.Len(t, result.ItemEncryptionStates[0], 65631)
	assert.Len(t, result.CarBlocks, 1)
	assert.EqualValues(t, 59, result.CarBlocks[0].CarOffset)
	v, _, _ := varint.FromUvarint(result.CarBlocks[0].Varint)
	assert.EqualValues(t, 220, v)
	assert.EqualValues(t, 0, result.CarBlocks[0].ItemOffset)
	assert.True(t, result.CarBlocks[0].ItemEncrypted)
}
