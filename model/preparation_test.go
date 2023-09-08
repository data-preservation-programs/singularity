package model

import (
	"testing"

	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestFile_FileName(t *testing.T) {
	tests := map[string]string{
		"test.txt":      "test.txt",
		"/test.txt":     "test.txt",
		"test/test.txt": "test.txt",
	}
	for path, name := range tests {
		t.Run(path, func(t *testing.T) {
			file := File{Path: path}
			if file.FileName() != name {
				t.Errorf("expected %s, got %s", name, file.FileName())
			}
		})
	}
}

func TestCarBlock_BlockLength(t *testing.T) {
	carBlock := CarBlock{
		CarBlockLength: 100,
		RawBlock:       []byte("test"),
	}
	require.EqualValues(t, 4, carBlock.BlockLength())

	carBlock = CarBlock{
		CarBlockLength: 100,
		Varint:         []byte("test"),
		CID:            CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))),
	}
	require.EqualValues(t, 100-4-36, carBlock.BlockLength())
}
