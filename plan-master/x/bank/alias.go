package bank

import (
	"github.com/plan-crypto/node/x/bank/keeper"
)

var (
	NewKeeper = keeper.NewKeeper
)

type (
	Keeper = keeper.Keeper
)
