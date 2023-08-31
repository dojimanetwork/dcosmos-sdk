package baseapp_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	dtproto "github.com/dojimanetwork/dojimamint/proto/tendermint/types"
	"github.com/stretchr/testify/require"
)

func TestValidateBlockParams(t *testing.T) {
	testCases := []struct {
		arg       interface{}
		expectErr bool
	}{
		{nil, true},
		{&dtproto.BlockParams{}, true},
		{dtproto.BlockParams{}, true},
		{dtproto.BlockParams{MaxBytes: -1, MaxGas: -1}, true},
		{dtproto.BlockParams{MaxBytes: 2000000, MaxGas: -5}, true},
		{dtproto.BlockParams{MaxBytes: 2000000, MaxGas: 300000}, false},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expectErr, baseapp.ValidateBlockParams(tc.arg) != nil)
	}
}

func TestValidateEvidenceParams(t *testing.T) {
	testCases := []struct {
		arg       interface{}
		expectErr bool
	}{
		{nil, true},
		{&dtproto.EvidenceParams{}, true},
		{dtproto.EvidenceParams{}, true},
		{dtproto.EvidenceParams{MaxAgeNumBlocks: -1, MaxAgeDuration: 18004000, MaxBytes: 5000000}, true},
		{dtproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: -1, MaxBytes: 5000000}, true},
		{dtproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: 18004000, MaxBytes: -1}, true},
		{dtproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: 18004000, MaxBytes: 5000000}, false},
		{dtproto.EvidenceParams{MaxAgeNumBlocks: 360000, MaxAgeDuration: 18004000, MaxBytes: 0}, false},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expectErr, baseapp.ValidateEvidenceParams(tc.arg) != nil)
	}
}

func TestValidateValidatorParams(t *testing.T) {
	testCases := []struct {
		arg       interface{}
		expectErr bool
	}{
		{nil, true},
		{&dtproto.ValidatorParams{}, true},
		{dtproto.ValidatorParams{}, true},
		{dtproto.ValidatorParams{PubKeyTypes: []string{}}, true},
		{dtproto.ValidatorParams{PubKeyTypes: []string{"secp256k1"}}, false},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expectErr, baseapp.ValidateValidatorParams(tc.arg) != nil)
	}
}
