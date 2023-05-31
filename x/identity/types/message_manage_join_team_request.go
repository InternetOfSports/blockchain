package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgManageJoinTeamRequest = "manage_join_team_request"

var _ sdk.Msg = &MsgManageJoinTeamRequest{}

func NewMsgManageJoinTeamRequest(creator string, team string, participant string, accept bool) *MsgManageJoinTeamRequest {
	return &MsgManageJoinTeamRequest{
		Creator:     creator,
		Team:        team,
		Participant: participant,
		Accept:      accept,
	}
}

func (msg *MsgManageJoinTeamRequest) Route() string {
	return RouterKey
}

func (msg *MsgManageJoinTeamRequest) Type() string {
	return TypeMsgManageJoinTeamRequest
}

func (msg *MsgManageJoinTeamRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgManageJoinTeamRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgManageJoinTeamRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
