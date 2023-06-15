package shitchain_test

import (
	"testing"

	keepertest "github.com/bonedaddy/shitchain/testutil/keeper"
	"github.com/bonedaddy/shitchain/testutil/nullify"
	"github.com/bonedaddy/shitchain/x/shitchain"
	"github.com/bonedaddy/shitchain/x/shitchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ShitchainKeeper(t)
	shitchain.InitGenesis(ctx, *k, genesisState)
	got := shitchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
