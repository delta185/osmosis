package osmosisibctesting

import (
	"fmt"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/osmosis-labs/osmosis/v10/app"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
)

func (chain *TestChain) StoreContractCode(suite *suite.Suite) {
	osmosisApp := chain.GetOsmosisApp()

	govKeeper := osmosisApp.GovKeeper
	wasmCode, err := ioutil.ReadFile("./testdata/rate_limiter.wasm")
	suite.Require().NoError(err)

	addr := osmosisApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	src := wasmtypes.StoreCodeProposalFixture(func(p *wasmtypes.StoreCodeProposal) {
		p.RunAs = addr.String()
		p.WASMByteCode = wasmCode
	})

	// when stored
	storedProposal, err := govKeeper.SubmitProposal(chain.GetContext(), src, false)
	suite.Require().NoError(err)

	// and proposal execute
	handler := govKeeper.Router().GetRoute(storedProposal.ProposalRoute())
	err = handler(chain.GetContext(), storedProposal.GetContent())
	suite.Require().NoError(err)
}

func (chain *TestChain) InstantiateContract(suite *suite.Suite) sdk.AccAddress {
	osmosisApp := chain.GetOsmosisApp()
	transferModule := osmosisApp.AccountKeeper.GetModuleAddress(transfertypes.ModuleName)

	initMsgBz := []byte(fmt.Sprintf(`{"ibc_module": "%s", "channel_quotas": [["test", 10]]}`, transferModule))
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(osmosisApp.WasmKeeper)
	codeID := uint64(1)
	creator := osmosisApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	addr, _, err := contractKeeper.Instantiate(chain.GetContext(), codeID, creator, creator, initMsgBz, "rate limiting contract", nil)
	suite.Require().NoError(err)
	return addr
}

func (chain *TestChain) GetOsmosisApp() *app.OsmosisApp {
	return chain.App.(*app.OsmosisApp)
}
