// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package proposal110

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	crypto "github.com/filecoin-project/go-state-types/crypto"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *DataRef) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{165}); err != nil {
		return err
	}

	// t.Root (cid.Cid) (struct)
	if len("Root") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Root\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Root"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Root")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.Root); err != nil {
		return xerrors.Errorf("failed to write cid field t.Root: %w", err)
	}

	// t.PieceCid (cid.Cid) (struct)
	if len("PieceCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceCid\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PieceCid"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceCid")); err != nil {
		return err
	}

	if t.PieceCid == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.PieceCid); err != nil {
			return xerrors.Errorf("failed to write cid field t.PieceCid: %w", err)
		}
	}

	// t.PieceSize (abi.UnpaddedPieceSize) (uint64)
	if len("PieceSize") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceSize\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PieceSize"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceSize")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.PieceSize)); err != nil {
		return err
	}

	// t.RawBlockSize (uint64) (uint64)
	if len("RawBlockSize") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"RawBlockSize\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RawBlockSize"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("RawBlockSize")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.RawBlockSize)); err != nil {
		return err
	}

	// t.TransferType (string) (string)
	if len("TransferType") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TransferType\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("TransferType"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TransferType")); err != nil {
		return err
	}

	if len(t.TransferType) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.TransferType was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.TransferType))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.TransferType)); err != nil {
		return err
	}
	return nil
}

func (t *DataRef) UnmarshalCBOR(r io.Reader) (err error) {
	*t = DataRef{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("DataRef: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Root (cid.Cid) (struct)
		case "Root":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Root: %w", err)
				}

				t.Root = c

			}
			// t.PieceCid (cid.Cid) (struct)
		case "PieceCid":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PieceCid: %w", err)
					}

					t.PieceCid = &c
				}

			}
			// t.PieceSize (abi.UnpaddedPieceSize) (uint64)
		case "PieceSize":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.PieceSize = abi.UnpaddedPieceSize(extra)

			}
			// t.RawBlockSize (uint64) (uint64)
		case "RawBlockSize":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.RawBlockSize = uint64(extra)

			}
			// t.TransferType (string) (string)
		case "TransferType":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.TransferType = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *Proposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.Piece (proposal110.DataRef) (struct)
	if len("Piece") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Piece\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Piece"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Piece")); err != nil {
		return err
	}

	if err := t.Piece.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.DealProposal (proposal110.ClientDealProposal) (struct)
	if len("DealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealProposal\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("DealProposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealProposal")); err != nil {
		return err
	}

	if err := t.DealProposal.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.FastRetrieval (bool) (bool)
	if len("FastRetrieval") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FastRetrieval\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("FastRetrieval"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("FastRetrieval")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.FastRetrieval); err != nil {
		return err
	}
	return nil
}

func (t *Proposal) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Proposal{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Proposal: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Piece (proposal110.DataRef) (struct)
		case "Piece":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Piece = new(DataRef)
					if err := t.Piece.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Piece pointer: %w", err)
					}
				}

			}
			// t.DealProposal (proposal110.ClientDealProposal) (struct)
		case "DealProposal":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.DealProposal = new(ClientDealProposal)
					if err := t.DealProposal.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.DealProposal pointer: %w", err)
					}
				}

			}
			// t.FastRetrieval (bool) (bool)
		case "FastRetrieval":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.FastRetrieval = false
			case 21:
				t.FastRetrieval = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ClientDealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.Proposal (proposal110.DealProposal) (struct)
	if len("Proposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Proposal\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Proposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Proposal")); err != nil {
		return err
	}

	if err := t.Proposal.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ClientSignature (crypto.Signature) (struct)
	if len("ClientSignature") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ClientSignature\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ClientSignature"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ClientSignature")); err != nil {
		return err
	}

	if err := t.ClientSignature.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *ClientDealProposal) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ClientDealProposal{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ClientDealProposal: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Proposal (proposal110.DealProposal) (struct)
		case "Proposal":

			{

				if err := t.Proposal.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Proposal: %w", err)
				}

			}
			// t.ClientSignature (crypto.Signature) (struct)
		case "ClientSignature":

			{

				if err := t.ClientSignature.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ClientSignature: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *DealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{171}); err != nil {
		return err
	}

	// t.Label (proposal110.DealLabel) (struct)
	if len("Label") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Label\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Label"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Label")); err != nil {
		return err
	}

	if err := t.Label.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Client (address.Address) (struct)
	if len("Client") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Client\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Client"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Client")); err != nil {
		return err
	}

	if err := t.Client.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if len("EndEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"EndEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("EndEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("EndEpoch")); err != nil {
		return err
	}

	if t.EndEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.EndEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.EndEpoch-1)); err != nil {
			return err
		}
	}

	// t.PieceCID (cid.Cid) (struct)
	if len("PieceCID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceCID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PieceCID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceCID")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	// t.Provider (address.Address) (struct)
	if len("Provider") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Provider\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Provider"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Provider")); err != nil {
		return err
	}

	if err := t.Provider.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.PieceSize (abi.PaddedPieceSize) (uint64)
	if len("PieceSize") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceSize\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PieceSize"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceSize")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.PieceSize)); err != nil {
		return err
	}

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if len("StartEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StartEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("StartEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StartEpoch")); err != nil {
		return err
	}

	if t.StartEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.StartEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.StartEpoch-1)); err != nil {
			return err
		}
	}

	// t.VerifiedDeal (bool) (bool)
	if len("VerifiedDeal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"VerifiedDeal\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("VerifiedDeal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("VerifiedDeal")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.VerifiedDeal); err != nil {
		return err
	}

	// t.ClientCollateral (big.Int) (struct)
	if len("ClientCollateral") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ClientCollateral\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ClientCollateral"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ClientCollateral")); err != nil {
		return err
	}

	if err := t.ClientCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ProviderCollateral (big.Int) (struct)
	if len("ProviderCollateral") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ProviderCollateral\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ProviderCollateral"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ProviderCollateral")); err != nil {
		return err
	}

	if err := t.ProviderCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.StoragePricePerEpoch (big.Int) (struct)
	if len("StoragePricePerEpoch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StoragePricePerEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("StoragePricePerEpoch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StoragePricePerEpoch")); err != nil {
		return err
	}

	if err := t.StoragePricePerEpoch.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *DealProposal) UnmarshalCBOR(r io.Reader) (err error) {
	*t = DealProposal{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("DealProposal: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Label (proposal110.DealLabel) (struct)
		case "Label":

			{

				if err := t.Label.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Label: %w", err)
				}

			}
			// t.Client (address.Address) (struct)
		case "Client":

			{

				if err := t.Client.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Client: %w", err)
				}

			}
			// t.EndEpoch (abi.ChainEpoch) (int64)
		case "EndEpoch":
			{
				maj, extra, err := cr.ReadHeader()
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
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.EndEpoch = abi.ChainEpoch(extraI)
			}
			// t.PieceCID (cid.Cid) (struct)
		case "PieceCID":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
				}

				t.PieceCID = c

			}
			// t.Provider (address.Address) (struct)
		case "Provider":

			{

				if err := t.Provider.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Provider: %w", err)
				}

			}
			// t.PieceSize (abi.PaddedPieceSize) (uint64)
		case "PieceSize":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.PieceSize = abi.PaddedPieceSize(extra)

			}
			// t.StartEpoch (abi.ChainEpoch) (int64)
		case "StartEpoch":
			{
				maj, extra, err := cr.ReadHeader()
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
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.StartEpoch = abi.ChainEpoch(extraI)
			}
			// t.VerifiedDeal (bool) (bool)
		case "VerifiedDeal":

			maj, extra, err = cr.ReadHeader()
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
			// t.ClientCollateral (big.Int) (struct)
		case "ClientCollateral":

			{

				if err := t.ClientCollateral.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ClientCollateral: %w", err)
				}

			}
			// t.ProviderCollateral (big.Int) (struct)
		case "ProviderCollateral":

			{

				if err := t.ProviderCollateral.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ProviderCollateral: %w", err)
				}

			}
			// t.StoragePricePerEpoch (big.Int) (struct)
		case "StoragePricePerEpoch":

			{

				if err := t.StoragePricePerEpoch.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.StoragePricePerEpoch: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *DealLabel) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{160}); err != nil {
		return err
	}
	return nil
}

func (t *DealLabel) UnmarshalCBOR(r io.Reader) (err error) {
	*t = DealLabel{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("DealLabel: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SignedResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.Response (proposal110.Response) (struct)
	if len("Response") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Response\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Response"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Response")); err != nil {
		return err
	}

	if err := t.Response.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Signature (crypto.Signature) (struct)
	if len("Signature") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Signature\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Signature"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Signature")); err != nil {
		return err
	}

	if err := t.Signature.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *SignedResponse) UnmarshalCBOR(r io.Reader) (err error) {
	*t = SignedResponse{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SignedResponse: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Response (proposal110.Response) (struct)
		case "Response":

			{

				if err := t.Response.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Response: %w", err)
				}

			}
			// t.Signature (crypto.Signature) (struct)
		case "Signature":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Signature = new(crypto.Signature)
					if err := t.Signature.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Signature pointer: %w", err)
					}
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *Response) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.State (uint64) (uint64)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("State")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.State)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("Message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Message\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Message")); err != nil {
		return err
	}

	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.Proposal (cid.Cid) (struct)
	if len("Proposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Proposal\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Proposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Proposal")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.PublishMessage (cid.Cid) (struct)
	if len("PublishMessage") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PublishMessage\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PublishMessage"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PublishMessage")); err != nil {
		return err
	}

	if t.PublishMessage == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.PublishMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishMessage: %w", err)
		}
	}

	return nil
}

func (t *Response) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Response{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Response: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.State (uint64) (uint64)
		case "State":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.State = uint64(extra)

			}
			// t.Message (string) (string)
		case "Message":

			{
				sval, err := cbg.ReadString(cr)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}
			// t.Proposal (cid.Cid) (struct)
		case "Proposal":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
				}

				t.Proposal = c

			}
			// t.PublishMessage (cid.Cid) (struct)
		case "PublishMessage":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PublishMessage: %w", err)
					}

					t.PublishMessage = &c
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
