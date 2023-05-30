package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestJoinTeam = "request_join_team"

var _ sdk.Msg = &MsgRequestJoinTeam{}

func NewMsgRequestJoinTeam(creator string, team string) *MsgRequestJoinTeam {
	return &MsgRequestJoinTeam{
		Creator: creator,
		Team:    team,
	}
}

func (msg *MsgRequestJoinTeam) Route() string {
	return RouterKey
}

func (msg *MsgRequestJoinTeam) Type() string {
	return TypeMsgRequestJoinTeam
}

func (msg *MsgRequestJoinTeam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestJoinTeam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestJoinTeam) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
