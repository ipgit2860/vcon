package keeper_test

import (
	"testing"

	testkeeper "github.com/ipgit2860/vcon/testutil/keeper"
	"github.com/ipgit2860/vcon/x/vcon/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VconKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
