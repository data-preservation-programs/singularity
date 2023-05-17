package proposal110

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
	"io"
	"unicode/utf8"
)

var lengthBufClientDealProposal = []byte{130}

func (t *ClientDealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufClientDealProposal); err != nil {
		return err
	}

	// t.Proposal (market.DealProposal) (struct)
	if err := t.Proposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ClientSignature (crypto.Signature) (struct)
	if err := t.ClientSignature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *ClientDealProposal) UnmarshalCBOR(r io.Reader) error {
	*t = ClientDealProposal{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Proposal (market.DealProposal) (struct)

	{

		if err := t.Proposal.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Proposal: %w", err)
		}

	}
	// t.ClientSignature (crypto.Signature) (struct)

	{

		if err := t.ClientSignature.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ClientSignature: %w", err)
		}

	}
	return nil
}

var lengthBufDealProposal = []byte{139}

func (t *DealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDealProposal); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.PieceCID (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	// t.PieceSize (abi.PaddedPieceSize) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PieceSize)); err != nil {
		return err
	}

	// t.VerifiedDeal (bool) (bool)
	if err := cbg.WriteBool(w, t.VerifiedDeal); err != nil {
		return err
	}

	// t.Client (address.Address) (struct)
	if err := t.Client.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Provider (address.Address) (struct)
	if err := t.Provider.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Label (market.DealLabel) (struct)
	if err := t.Label.MarshalCBOR(w); err != nil {
		return err
	}

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if t.StartEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.StartEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.StartEpoch-1)); err != nil {
			return err
		}
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if t.EndEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.EndEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.EndEpoch-1)); err != nil {
			return err
		}
	}

	// t.StoragePricePerEpoch (big.Int) (struct)
	if err := t.StoragePricePerEpoch.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ProviderCollateral (big.Int) (struct)
	if err := t.ProviderCollateral.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ClientCollateral (big.Int) (struct)
	if err := t.ClientCollateral.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealProposal) UnmarshalCBOR(r io.Reader) error {
	*t = DealProposal{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 11 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PieceCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
		}

		t.PieceCID = c

	}
	// t.PieceSize (abi.PaddedPieceSize) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PieceSize = abi.PaddedPieceSize(extra)

	}
	// t.VerifiedDeal (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.VerifiedDeal = false
	case 21:
		t.VerifiedDeal = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.Client (address.Address) (struct)

	{

		if err := t.Client.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Client: %w", err)
		}

	}
	// t.Provider (address.Address) (struct)

	{

		if err := t.Provider.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Provider: %w", err)
		}

	}
	// t.Label (market.DealLabel) (struct)

	{

		if err := t.Label.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Label: %w", err)
		}

	}
	// t.StartEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.StartEpoch = abi.ChainEpoch(extraI)
	}
	// t.EndEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.EndEpoch = abi.ChainEpoch(extraI)
	}
	// t.StoragePricePerEpoch (big.Int) (struct)

	{

		if err := t.StoragePricePerEpoch.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.StoragePricePerEpoch: %w", err)
		}

	}
	// t.ProviderCollateral (big.Int) (struct)

	{

		if err := t.ProviderCollateral.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ProviderCollateral: %w", err)
		}

	}
	// t.ClientCollateral (big.Int) (struct)

	{

		if err := t.ClientCollateral.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ClientCollateral: %w", err)
		}

	}
	return nil
}

func (label *DealLabel) MarshalCBOR(w io.Writer) error {
	scratch := make([]byte, 9)

	// nil *DealLabel counts as EmptyLabel
	// on chain structures should never have a pointer to a DealLabel but the case is included for completeness
	if label == nil {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, 0); err != nil {
			return err
		}
		_, err := io.WriteString(w, string(""))
		return err
	}
	if len(label.bs) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("label is too long to marshal (%d), max allowed (%d)", len(label.bs), cbg.ByteArrayMaxLen)
	}

	majorType := byte(cbg.MajByteString)
	if label.IsString() {
		majorType = cbg.MajTextString
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, majorType, uint64(len(label.bs))); err != nil {
		return err
	}
	_, err := w.Write(label.bs)
	return err
}

func (label *DealLabel) UnmarshalCBOR(br io.Reader) error {
	if label == nil {
		return xerrors.Errorf("cannot unmarshal into nil pointer")
	}

	// reset fields
	label.bs = nil

	scratch := make([]byte, 8)

	maj, length, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajTextString && maj != cbg.MajByteString {
		return fmt.Errorf("unexpected major tag (%d) when unmarshaling DealLabel: only textString (%d) or byteString (%d) expected", maj, cbg.MajTextString, cbg.MajByteString)
	}
	if length > cbg.ByteArrayMaxLen {
		return fmt.Errorf("label was too long (%d), max allowed (%d)", length, cbg.ByteArrayMaxLen)
	}
	buf := make([]byte, length)
	_, err = io.ReadAtLeast(br, buf, int(length))
	if err != nil {
		return err
	}
	label.bs = buf
	label.notString = maj != cbg.MajTextString
	if !label.notString && !utf8.ValidString(string(buf)) {
		return fmt.Errorf("label string not valid utf8")
	}

	return nil
}

func (label DealLabel) IsString() bool {
	return !label.notString
}
