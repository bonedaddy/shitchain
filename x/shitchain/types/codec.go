package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateThread{}, "shitchain/CreateThread", nil)
	cdc.RegisterConcrete(&MsgUpdateThread{}, "shitchain/UpdateThread", nil)
	cdc.RegisterConcrete(&MsgDeleteThread{}, "shitchain/DeleteThread", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateThread{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateThread{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteThread{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
