package codec

import (
	dtmcrypto "github.com/dojimanetwork/dojimamint/crypto"
	tmcrypto "github.com/dojimanetwork/dojimamint/crypto"
	"github.com/dojimanetwork/dojimamint/crypto/encoding"
	dencoding "github.com/dojimanetwork/dojimamint/crypto/encoding"
	dtmprotocrypto "github.com/dojimanetwork/dojimamint/proto/dojimamint/crypto"
	tmprotocrypto "github.com/dojimanetwork/dojimamint/proto/dojimamint/crypto"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// FromTmProtoPublicKey converts a TM's tmprotocrypto.PublicKey into our own PubKey.
func FromTmProtoPublicKey(protoPk tmprotocrypto.PublicKey) (cryptotypes.PubKey, error) {
	switch protoPk := protoPk.Sum.(type) {
	case *tmprotocrypto.PublicKey_Ed25519:
		return &ed25519.PubKey{
			Key: protoPk.Ed25519,
		}, nil
	case *tmprotocrypto.PublicKey_Secp256K1:
		return &secp256k1.PubKey{
			Key: protoPk.Secp256K1,
		}, nil
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot convert %v from Tendermint public key", protoPk)
	}
}

// FromTmProtoPublicKey converts a TM's tmprotocrypto.PublicKey into our own PubKey.
func FromDTmProtoPublicKey(protoPk dtmprotocrypto.PublicKey) (cryptotypes.PubKey, error) {
	switch protoPk := protoPk.Sum.(type) {
	case *dtmprotocrypto.PublicKey_Ed25519:
		return &ed25519.PubKey{
			Key: protoPk.Ed25519,
		}, nil
	case *dtmprotocrypto.PublicKey_Secp256K1:
		return &secp256k1.PubKey{
			Key: protoPk.Secp256K1,
		}, nil
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot convert %v from Tendermint public key", protoPk)
	}
}

// ToTmProtoPublicKey converts our own PubKey to TM's tmprotocrypto.PublicKey.
func ToTmProtoPublicKey(pk cryptotypes.PubKey) (tmprotocrypto.PublicKey, error) {
	switch pk := pk.(type) {
	case *ed25519.PubKey:
		return tmprotocrypto.PublicKey{
			Sum: &tmprotocrypto.PublicKey_Ed25519{
				Ed25519: pk.Key,
			},
		}, nil
	case *secp256k1.PubKey:
		return tmprotocrypto.PublicKey{
			Sum: &tmprotocrypto.PublicKey_Secp256K1{
				Secp256K1: pk.Key,
			},
		}, nil
	default:
		return tmprotocrypto.PublicKey{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot convert %v to Tendermint public key", pk)
	}
}

// ToTmProtoPublicKey converts our own PubKey to TM's tmprotocrypto.PublicKey.
func ToDTmProtoPublicKey(pk cryptotypes.PubKey) (dtmprotocrypto.PublicKey, error) {
	switch pk := pk.(type) {
	case *ed25519.PubKey:
		return dtmprotocrypto.PublicKey{
			Sum: &dtmprotocrypto.PublicKey_Ed25519{
				Ed25519: pk.Key,
			},
		}, nil
	case *secp256k1.PubKey:
		return dtmprotocrypto.PublicKey{
			Sum: &dtmprotocrypto.PublicKey_Secp256K1{
				Secp256K1: pk.Key,
			},
		}, nil
	default:
		return dtmprotocrypto.PublicKey{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot convert %v to Tendermint public key", pk)
	}
}

// FromTmPubKeyInterface converts TM's tmcrypto.PubKey to our own PubKey.
func FromTmPubKeyInterface(tmPk tmcrypto.PubKey) (cryptotypes.PubKey, error) {
	tmProtoPk, err := encoding.PubKeyToProto(tmPk)
	if err != nil {
		return nil, err
	}

	return FromTmProtoPublicKey(tmProtoPk)
}

func FromDTmPubKeyInterface(dtmPk dtmcrypto.PubKey) (cryptotypes.PubKey, error) {
	tmProtoPk, err := dencoding.PubKeyToProto(dtmPk)
	if err != nil {
		return nil, err
	}

	return FromDTmProtoPublicKey(tmProtoPk)
}

// ToTmPubKeyInterface converts our own PubKey to TM's tmcrypto.PubKey.
func ToTmPubKeyInterface(pk cryptotypes.PubKey) (tmcrypto.PubKey, error) {
	tmProtoPk, err := ToTmProtoPublicKey(pk)
	if err != nil {
		return nil, err
	}

	return encoding.PubKeyFromProto(tmProtoPk)
}

// ToTmPubKeyInterface converts our own PubKey to TM's tmcrypto.PubKey.
func ToDTmPubKeyInterface(pk cryptotypes.PubKey) (dtmcrypto.PubKey, error) {
	tmProtoPk, err := ToDTmProtoPublicKey(pk)
	if err != nil {
		return nil, err
	}

	return dencoding.PubKeyFromProto(tmProtoPk)
}
