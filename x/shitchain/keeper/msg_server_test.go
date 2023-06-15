package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/bonedaddy/shitchain/testutil/keeper"
	"github.com/bonedaddy/shitchain/x/shitchain/keeper"
	"github.com/bonedaddy/shitchain/x/shitchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ShitchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
