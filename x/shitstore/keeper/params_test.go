package keeper_test

import (
	"testing"

	testkeeper "github.com/bonedaddy/shitchain/testutil/keeper"
	"github.com/bonedaddy/shitchain/x/shitstore/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ShitstoreKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
