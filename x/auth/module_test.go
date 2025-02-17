package auth_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	//abcitypes "github.com/tendermint/tendermint/abci/types"
	dabcitypes "github.com/dojimanetwork/dojimamint/abci/types"
	//tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dtmproto "github.com/dojimanetwork/dojimamint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, dtmproto.Header{})

	app.InitChain(
		dabcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	acc := app.AccountKeeper.GetAccount(ctx, types.NewModuleAddress(types.FeeCollectorName))
	require.NotNil(t, acc)
}
