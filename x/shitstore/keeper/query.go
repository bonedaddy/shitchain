package keeper

import (
	"github.com/bonedaddy/shitchain/x/shitstore/types"
)

var _ types.QueryServer = Keeper{}
