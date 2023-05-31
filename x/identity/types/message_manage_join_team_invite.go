package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgManageJoinTeamInvite = "manage_join_team_invite"

var _ sdk.Msg = &MsgManageJoinTeamInvite{}

func NewMsgManageJoinTeamInvite(creator string, team string, accept bool) *MsgManageJoinTeamInvite {
	return &MsgManageJoinTeamInvite{
		Creator: creator,
		Team:    team,
		Accept:  accept,
	}
}

func (msg *MsgManageJoinTeamInvite) Route() string {
	return RouterKey
}

func (msg *MsgManageJoinTeamInvite) Type() string {
	return TypeMsgManageJoinTeamInvite
}

func (msg *MsgManageJoinTeamInvite) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgManageJoinTeamInvite) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgManageJoinTeamInvite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
