package vcon_test

import (
	"testing"

	keepertest "github.com/ipgit2860/vcon/testutil/keeper"
	"github.com/ipgit2860/vcon/testutil/nullify"
	"github.com/ipgit2860/vcon/x/vcon"
	"github.com/ipgit2860/vcon/x/vcon/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VconKeeper(t)
	vcon.InitGenesis(ctx, *k, genesisState)
	got := vcon.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
