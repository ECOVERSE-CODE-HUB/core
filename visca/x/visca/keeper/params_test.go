package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "visca/testutil/keeper"
	"visca/x/visca/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ViscaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
