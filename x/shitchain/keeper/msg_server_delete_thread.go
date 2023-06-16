package keeper

import (
	"context"

	"github.com/bonedaddy/shitchain/x/shitchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteThread(goCtx context.Context, msg *types.MsgDeleteThread) (*types.MsgDeleteThreadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteThreadResponse{}, nil
}
