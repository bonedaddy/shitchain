package shitchain

import (
	"math/rand"

	"github.com/bonedaddy/shitchain/testutil/sample"
	shitchainsimulation "github.com/bonedaddy/shitchain/x/shitchain/simulation"
	"github.com/bonedaddy/shitchain/x/shitchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = shitchainsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateThread = "op_weight_msg_create_thread"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateThread int = 100

	opWeightMsgUpdateThread = "op_weight_msg_update_thread"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateThread int = 100

	opWeightMsgDeleteThread = "op_weight_msg_delete_thread"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteThread int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	shitchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&shitchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateThread int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateThread, &weightMsgCreateThread, nil,
		func(_ *rand.Rand) {
			weightMsgCreateThread = defaultWeightMsgCreateThread
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateThread,
		shitchainsimulation.SimulateMsgCreateThread(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateThread int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateThread, &weightMsgUpdateThread, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateThread = defaultWeightMsgUpdateThread
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateThread,
		shitchainsimulation.SimulateMsgUpdateThread(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteThread int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteThread, &weightMsgDeleteThread, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteThread = defaultWeightMsgDeleteThread
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteThread,
		shitchainsimulation.SimulateMsgDeleteThread(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateThread,
			defaultWeightMsgCreateThread,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				shitchainsimulation.SimulateMsgCreateThread(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateThread,
			defaultWeightMsgUpdateThread,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				shitchainsimulation.SimulateMsgUpdateThread(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteThread,
			defaultWeightMsgDeleteThread,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				shitchainsimulation.SimulateMsgDeleteThread(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
