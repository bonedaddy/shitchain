package simulation

import (
	"math/rand"

	"github.com/bonedaddy/shitchain/x/shitstore/keeper"
	"github.com/bonedaddy/shitchain/x/shitstore/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgBroadcastPost(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBroadcastPost{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the BroadcastPost simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BroadcastPost simulation not implemented"), nil, nil
	}
}
