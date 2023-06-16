package keeper

import (
	"context"

	"github.com/bonedaddy/shitchain/x/shitchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateThread(goCtx context.Context, msg *types.MsgUpdateThread) (*types.MsgUpdateThreadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateThreadResponse{}, nil
}
