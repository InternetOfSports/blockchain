package identity

import (
	"math/rand"

	"github.com/InternetOfSports/blockchain/testutil/sample"
	identitysimulation "github.com/InternetOfSports/blockchain/x/identity/simulation"
	"github.com/InternetOfSports/blockchain/x/identity/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = identitysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterParticipant = "op_weight_msg_register_participant"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterParticipant int = 100

	opWeightMsgCreateTeam = "op_weight_msg_create_team"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTeam int = 100

	opWeightMsgSetTeamManager = "op_weight_msg_set_team_manager"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetTeamManager int = 100

	opWeightMsgSetTeamTrainer = "op_weight_msg_set_team_trainer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetTeamTrainer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterParticipant int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterParticipant, &weightMsgRegisterParticipant, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterParticipant = defaultWeightMsgRegisterParticipant
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterParticipant,
		identitysimulation.SimulateMsgRegisterParticipant(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTeam int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTeam, &weightMsgCreateTeam, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTeam = defaultWeightMsgCreateTeam
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTeam,
		identitysimulation.SimulateMsgCreateTeam(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetTeamManager int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetTeamManager, &weightMsgSetTeamManager, nil,
		func(_ *rand.Rand) {
			weightMsgSetTeamManager = defaultWeightMsgSetTeamManager
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetTeamManager,
		identitysimulation.SimulateMsgSetTeamManager(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetTeamTrainer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetTeamTrainer, &weightMsgSetTeamTrainer, nil,
		func(_ *rand.Rand) {
			weightMsgSetTeamTrainer = defaultWeightMsgSetTeamTrainer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetTeamTrainer,
		identitysimulation.SimulateMsgSetTeamTrainer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
