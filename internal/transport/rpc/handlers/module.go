package handlers

import (
	"Crash-Currency-service/pkg/proto"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers",
	fx.Provide(
		NewCurrencyHandler,
		func(h *CurrencyHandler) proto.CurrencyServiceServer { return h },
	),
)
