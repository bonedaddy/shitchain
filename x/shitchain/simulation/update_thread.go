package simulation

import (
	"math/rand"

	"github.com/bonedaddy/shitchain/x/shitchain/keeper"
	"github.com/bonedaddy/shitchain/x/shitchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateThread(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateThread{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateThread simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateThread simulation not implemented"), nil, nil
	}
}
