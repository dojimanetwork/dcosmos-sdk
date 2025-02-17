package tmservice

import (
	"context"

	//ctypes "github.com/tendermint/tendermint/rpc/core/types"
	dtypes "github.com/dojimanetwork/dojimamint/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/client"
)

func getNodeStatus(ctx context.Context, clientCtx client.Context) (*dtypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &dtypes.ResultStatus{}, err
	}
	return node.Status(ctx)
}
