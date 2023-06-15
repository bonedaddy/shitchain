package shitstore_test

import (
	"testing"

	keepertest "github.com/bonedaddy/shitchain/testutil/keeper"
	"github.com/bonedaddy/shitchain/testutil/nullify"
	"github.com/bonedaddy/shitchain/x/shitstore"
	"github.com/bonedaddy/shitchain/x/shitstore/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ShitstoreKeeper(t)
	shitstore.InitGenesis(ctx, *k, genesisState)
	got := shitstore.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
