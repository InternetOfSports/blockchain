package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetTeamTrainer = "set_team_trainer"

var _ sdk.Msg = &MsgSetTeamTrainer{}

func NewMsgSetTeamTrainer(creator string, team string, trainer string) *MsgSetTeamTrainer {
	return &MsgSetTeamTrainer{
		Creator: creator,
		Team:    team,
		Trainer: trainer,
	}
}

func (msg *MsgSetTeamTrainer) Route() string {
	return RouterKey
}

func (msg *MsgSetTeamTrainer) Type() string {
	return TypeMsgSetTeamTrainer
}

func (msg *MsgSetTeamTrainer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetTeamTrainer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetTeamTrainer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
