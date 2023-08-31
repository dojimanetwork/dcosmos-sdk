package tmservice

import (
	"context"

	dtmproto "github.com/dojimanetwork/dojimamint/proto/tendermint/types"
	dtypes "github.com/dojimanetwork/dojimamint/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/client"
)

func getBlock(ctx context.Context, clientCtx client.Context, height *int64) (*dtypes.ResultBlock, error) {
	// get the node
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	return node.Block(ctx, height)
}

func GetProtoBlock(ctx context.Context, clientCtx client.Context, height *int64) (dtmproto.BlockID, *dtmproto.Block, error) {
	block, err := getBlock(ctx, clientCtx, height)
	if err != nil {
		return dtmproto.BlockID{}, nil, err
	}
	protoBlock, err := block.Block.ToProto()
	if err != nil {
		return dtmproto.BlockID{}, nil, err
	}
	protoBlockId := block.BlockID.ToProto()

	return protoBlockId, protoBlock, nil
}
