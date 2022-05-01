package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/PanGan21/loan/testutil/keeper"
	"github.com/PanGan21/loan/x/loan/keeper"
	"github.com/PanGan21/loan/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LoanKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
