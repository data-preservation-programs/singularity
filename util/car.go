package util

import (
	"github.com/cockroachdb/errors"
	"github.com/ipfs/go-cid"
	cbornode "github.com/ipfs/go-ipld-cbor"
	"github.com/ipld/go-car"
	"github.com/multiformats/go-varint"
)

// GenerateCarHeader generates the CAR (Content Addressable aRchive) format header
// based on the given root CID (Content Identifier).
//
// Parameters:
//   - root: The root CID of the MerkleDag that the CAR file represents.
//
// Returns:
//   - []byte: The byte representation of the CAR header.
//   - error: An error that can occur during the header generation process, or nil if successful.
func GenerateCarHeader(root cid.Cid) ([]byte, error) {
	header := car.CarHeader{
		Roots:   []cid.Cid{root},
		Version: 1,
	}

	headerBytes, err := cbornode.DumpObject(&header)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dump header")
	}
	headerBytesVarint := varint.ToUvarint(uint64(len(headerBytes)))
	headerBytes = append(headerBytesVarint, headerBytes...)
	return headerBytes, nil
}
