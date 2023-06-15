package keeper

import (
	"context"

	"github.com/bonedaddy/shitchain/x/shitstore/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BroadcastPost(goCtx context.Context, msg *types.MsgBroadcastPost) (*types.MsgBroadcastPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBroadcastPostResponse{}, nil
}
