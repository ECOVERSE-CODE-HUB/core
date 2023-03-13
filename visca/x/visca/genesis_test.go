package visca_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "visca/testutil/keeper"
	"visca/testutil/nullify"
	"visca/x/visca"
	"visca/x/visca/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ViscaKeeper(t)
	visca.InitGenesis(ctx, *k, genesisState)
	got := visca.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
