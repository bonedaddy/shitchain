package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBroadcastPost = "broadcast_post"

var _ sdk.Msg = &MsgBroadcastPost{}

func NewMsgBroadcastPost(creator string, post *Post) *MsgBroadcastPost {
	return &MsgBroadcastPost{
		Creator: creator,
		Post:    post,
	}
}

func (msg *MsgBroadcastPost) Route() string {
	return RouterKey
}

func (msg *MsgBroadcastPost) Type() string {
	return TypeMsgBroadcastPost
}

func (msg *MsgBroadcastPost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBroadcastPost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBroadcastPost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
