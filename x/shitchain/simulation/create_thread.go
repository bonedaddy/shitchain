package simulation

import (
	"math/rand"

	"github.com/bonedaddy/shitchain/x/shitchain/keeper"
	"github.com/bonedaddy/shitchain/x/shitchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateThread(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateThread{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateThread simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateThread simulation not implemented"), nil, nil
	}
}
