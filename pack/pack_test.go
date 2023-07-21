//go:build !windows

package pack

import (
	"bytes"
	"context"
	"crypto/rand"
	"io"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/multiformats/go-varint"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "", result.CarResults[0].CarFilePath)
	require.EqualValues(t, 10486793, result.CarResults[0].CarFileSize)
	require.EqualValues(t, 1<<30, result.CarResults[0].PieceSize)
	require.Len(t, result.CarResults[0].Header, 59)
	require.Len(t, result.CarResults[0].CarBlocks, 12)
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
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "", result.CarResults[0].CarFilePath)
	require.EqualValues(t, 101, result.CarResults[0].CarFileSize)
	require.Equal(t, "baga6ea4seaqddc4kqdxmnglxhmrfxkx2cd7hl6kounifgewgb2hdthmeddfe4hy", result.CarResults[0].PieceCID.String())
	require.EqualValues(t, 1<<20, result.CarResults[0].PieceSize)
	require.Len(t, result.CarResults[0].Header, 59)
	require.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", result.CarResults[0].RootCID.String())
	require.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", result.ItemPartCIDs[0].String())
	require.Len(t, result.CarResults[0].CarBlocks, 1)
	require.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", result.CarResults[0].CarBlocks[0].CID.String())
	require.EqualValues(t, 59, result.CarResults[0].CarBlocks[0].CarOffset)
	v, _, _ := varint.FromUvarint(result.CarResults[0].CarBlocks[0].Varint)
	require.EqualValues(t, 41, v)
	require.EqualValues(t, 1, result.CarResults[0].CarBlocks[0].ItemOffset)
}

func TestAssembleCar_WithEncryption(t *testing.T) {
	ctx := context.Background()
	handler := new(MockReadHandler)
	handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(io.NopCloser(bytes.NewReader([]byte("hello"))), nil, nil)
	items := []model.ItemPart{
		{
			Offset: 0,
			Length: 5,
			Item: &model.Item{
				Size: 5,
			},
		},
	}
	result, err := AssembleCar(ctx, handler, model.Dataset{
		EncryptionRecipients: []string{"age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"},
	}, items, t.TempDir(), 1<<20)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotEmpty(t, result.CarResults[0].CarFilePath)
	require.EqualValues(t, 302, result.CarResults[0].CarFileSize)
	require.EqualValues(t, 1<<20, result.CarResults[0].PieceSize)
	require.Len(t, result.CarResults[0].Header, 59)
	require.Len(t, result.CarResults[0].CarBlocks, 0)
}
