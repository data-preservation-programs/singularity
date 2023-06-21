package cliutil

import (
	"github.com/ipfs/go-cid"
	"testing"
)

type Widget struct {
	ID    int
	Name  string
	Cost  float64
	Added string
	CID   []byte
}

func TestPrintSingleObject(t *testing.T) {
	c := cid.MustParse("QmPK1s3pNYLi9ERiq3BDxKa4XosgWwFRQUydHUtz4YgpqB")
	widget := Widget{ID: 1, Name: "Example", Cost: 3.14, Added: "2023-05-08", CID: c.Bytes()}
	PrintToConsole(widget, false, nil)
	PrintToConsole([]Widget{widget}, false, nil)
}
