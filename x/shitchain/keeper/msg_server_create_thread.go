package keeper

import (
	"context"

	"github.com/bonedaddy/shitchain/x/shitchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateThread(goCtx context.Context, msg *types.MsgCreateThread) (*types.MsgCreateThreadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateThreadResponse{}, nil
}
