package v4

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"

	appparams "github.com/osmosis-labs/osmosis/v25/app/params"
)

func Prop12(ctx sdk.Context, bank bankkeeper.Keeper, distr *distrkeeper.Keeper) {
	payments := GetProp12Payments()

	total := int64(0)

	for _, payment := range payments {
		addr, err := sdk.AccAddressFromBech32(payment[0])
		if err != nil {
			panic(err)
		}

		amount, err := strconv.ParseInt(strings.TrimSpace(payment[1]), 10, 64)
		if err != nil {
			panic(err)
		}

		coins := sdk.NewCoins(sdk.NewInt64Coin(appparams.BaseCoinUnit, amount))
		if err := bank.SendCoinsFromModuleToAccount(ctx, "distribution", addr, coins); err != nil {
			panic(err)
		}
		total += amount
	}

	// deduct from the feePool tracker
	feePool := distr.GetFeePool(ctx)
	feePool.CommunityPool = feePool.CommunityPool.Sub(sdk.NewDecCoins(osmomath.NewInt64DecCoin(appparams.BaseCoinUnit, total)))
	distr.SetFeePool(ctx, feePool)
}
