package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePost{}, "shitstore/CreatePost", nil)
	cdc.RegisterConcrete(&MsgUpdatePost{}, "shitstore/UpdatePost", nil)
	cdc.RegisterConcrete(&MsgDeletePost{}, "shitstore/DeletePost", nil)
	cdc.RegisterConcrete(&MsgBroadcastPost{}, "shitstore/BroadcastPost", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
		&MsgUpdatePost{},
		&MsgDeletePost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBroadcastPost{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
