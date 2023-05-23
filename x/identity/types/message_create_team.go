package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateTeam = "create_team"

var _ sdk.Msg = &MsgCreateTeam{}

func NewMsgCreateTeam(creator string, name string) *MsgCreateTeam {
	return &MsgCreateTeam{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgCreateTeam) Route() string {
	return RouterKey
}

func (msg *MsgCreateTeam) Type() string {
	return TypeMsgCreateTeam
}

func (msg *MsgCreateTeam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTeam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTeam) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
