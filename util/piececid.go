package util

import (
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

const (
	// CommPv2 CIDs currently used by PDP contracts are raw CIDs using
	// Filecoin unsealed commitment multihash with 35-byte digest payload.
	commPv2Codec       = cid.Raw
	commPv2MhCode      = 0x1011
	commPv2MhDigestLen = 35
)

// IsAcceptedPieceCID returns true for legacy CommP CIDs and CommPv2 CIDs.
func IsAcceptedPieceCID(pieceCID cid.Cid) bool {
	if pieceCID.Type() == cid.FilCommitmentUnsealed {
		return true
	}
	if pieceCID.Type() != commPv2Codec {
		return false
	}

	decoded, err := mh.Decode(pieceCID.Hash())
	if err != nil {
		return false
	}
	return decoded.Code == commPv2MhCode && decoded.Length == commPv2MhDigestLen
}
