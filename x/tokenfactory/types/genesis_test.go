package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/osmosis-labs/osmosis/v13/x/tokenfactory/types"
)

func TestGenesisState_Validate(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "different admin from creator",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "osmo1ft6e5esdtdegnvcr3djd3ftk4kwpcr6jrx5fj9",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "empty admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "no admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
					},
				},
			},
			valid: true,
		},
		{
			desc: "invalid admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "moose",
						},
					},
				},
			},
			valid: false,
		},
		{
			desc: "multiple denoms",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/litecoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicate denoms",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			valid: false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
