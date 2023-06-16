package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateThread = "create_thread"

var _ sdk.Msg = &MsgCreateThread{}

func NewMsgCreateThread(creator string, title string, body string) *MsgCreateThread {
	return &MsgCreateThread{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgCreateThread) Route() string {
	return RouterKey
}

func (msg *MsgCreateThread) Type() string {
	return TypeMsgCreateThread
}

func (msg *MsgCreateThread) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateThread) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateThread) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
