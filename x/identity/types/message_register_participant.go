package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterParticipant = "register_participant"

var _ sdk.Msg = &MsgRegisterParticipant{}

func NewMsgRegisterParticipant(creator string, nickname string) *MsgRegisterParticipant {
	return &MsgRegisterParticipant{
		Creator:  creator,
		Nickname: nickname,
	}
}

func (msg *MsgRegisterParticipant) Route() string {
	return RouterKey
}

func (msg *MsgRegisterParticipant) Type() string {
	return TypeMsgRegisterParticipant
}

func (msg *MsgRegisterParticipant) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterParticipant) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterParticipant) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
