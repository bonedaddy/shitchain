package keeper

import (
	"context"

	"github.com/bonedaddy/shitchain/x/shitchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowThread(goCtx context.Context, req *types.QueryShowThreadRequest) (*types.QueryShowThreadResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryShowThreadResponse{}, nil
}