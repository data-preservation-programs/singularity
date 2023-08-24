package cliutil

import (
	"time"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"testing"
)

type Widget struct {
	ID      int       `cli:"normal"`
	Name    string    `cli:"normal"`
	Cost    float64   `cli:"normal"`
	Time    time.Time `cli:"normal"`
	CID     []byte    `cli:"normal"`
	Verbose bool      `cli:"verbose"`
	Hidden  string
}

func TestPrintSingleObject(t *testing.T) {
	c, err := cid.Decode("QmPK1s3pNYLi9ERiq3BDxKa4XosgWwFRQUydHUtz4YgpqB")
	require.NoError(t, err)
	widget := Widget{
		ID:      1,
		Name:    "Example",
		Cost:    3.14,
		Time:    time.Time{},
		CID:     c.Bytes(),
		Verbose: true,
		Hidden:  "hidden",
	}
	PrintToConsole(widget, false, true)
	PrintToConsole(widget, false, false)
	PrintToConsole([]Widget{widget}, false, true)
	PrintToConsole([]Widget{widget}, false, false)
}
