package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/ipgit2860/vcon/testutil/keeper"
	"github.com/ipgit2860/vcon/x/vcon/keeper"
	"github.com/ipgit2860/vcon/x/vcon/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VconKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
