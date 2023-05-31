package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterParticipant{}, "identity/RegisterParticipant", nil)
	cdc.RegisterConcrete(&MsgCreateTeam{}, "identity/CreateTeam", nil)
	cdc.RegisterConcrete(&MsgSetTeamManager{}, "identity/SetTeamManager", nil)
	cdc.RegisterConcrete(&MsgSetTeamTrainer{}, "identity/SetTeamTrainer", nil)
	cdc.RegisterConcrete(&MsgRequestJoinTeam{}, "identity/RequestJoinTeam", nil)
	cdc.RegisterConcrete(&MsgManageJoinTeamRequest{}, "identity/ManageJoinTeamRequest", nil)
	cdc.RegisterConcrete(&MsgInviteParticipantToJoinTeam{}, "identity/InviteParticipantToJoinTeam", nil)
	cdc.RegisterConcrete(&MsgManageJoinTeamInvite{}, "identity/ManageJoinTeamInvite", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterParticipant{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTeam{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetTeamManager{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetTeamTrainer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestJoinTeam{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgManageJoinTeamRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInviteParticipantToJoinTeam{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgManageJoinTeamInvite{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
