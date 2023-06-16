package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateThread = "update_thread"

var _ sdk.Msg = &MsgUpdateThread{}

func NewMsgUpdateThread(creator string, title string, body string, id uint64) *MsgUpdateThread {
	return &MsgUpdateThread{
		Creator: creator,
		Title:   title,
		Body:    body,
		Id:      id,
	}
}

func (msg *MsgUpdateThread) Route() string {
	return RouterKey
}

func (msg *MsgUpdateThread) Type() string {
	return TypeMsgUpdateThread
}

func (msg *MsgUpdateThread) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateThread) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateThread) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
