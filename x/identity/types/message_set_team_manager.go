package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetTeamManager = "set_team_manager"

var _ sdk.Msg = &MsgSetTeamManager{}

func NewMsgSetTeamManager(creator string, team string, manager string) *MsgSetTeamManager {
	return &MsgSetTeamManager{
		Creator: creator,
		Team:    team,
		Manager: manager,
	}
}

func (msg *MsgSetTeamManager) Route() string {
	return RouterKey
}

func (msg *MsgSetTeamManager) Type() string {
	return TypeMsgSetTeamManager
}

func (msg *MsgSetTeamManager) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetTeamManager) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetTeamManager) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
