package tmservice

import (
	"context"

	dtypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/client"
)

func getNodeStatus(ctx context.Context, clientCtx client.Context) (*dtypes.ResultStatus, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &dtypes.ResultStatus{}, err
	}
	return node.Status(ctx)
}
