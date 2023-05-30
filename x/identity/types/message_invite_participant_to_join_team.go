package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInviteParticipantToJoinTeam = "invite_participant_to_join_team"

var _ sdk.Msg = &MsgInviteParticipantToJoinTeam{}

func NewMsgInviteParticipantToJoinTeam(creator string, team string, participant string) *MsgInviteParticipantToJoinTeam {
	return &MsgInviteParticipantToJoinTeam{
		Creator:     creator,
		Team:        team,
		Participant: participant,
	}
}

func (msg *MsgInviteParticipantToJoinTeam) Route() string {
	return RouterKey
}

func (msg *MsgInviteParticipantToJoinTeam) Type() string {
	return TypeMsgInviteParticipantToJoinTeam
}

func (msg *MsgInviteParticipantToJoinTeam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInviteParticipantToJoinTeam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInviteParticipantToJoinTeam) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
