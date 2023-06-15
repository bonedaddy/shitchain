package keeper

import (
	"github.com/bonedaddy/shitchain/x/shitchain/types"
)

var _ types.QueryServer = Keeper{}
