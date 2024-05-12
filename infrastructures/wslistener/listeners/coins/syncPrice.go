package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/listener/ws"
)

func (l Listeners) SyncPrice(wsc ws.Context) {
	var req dto.SyncCoinPriceRequest
	if err := req.FromWsContext(wsc); err != nil {
		l.log.Error(context.Background(), "sync-price-bind-error").Err(err).Send()
	}

	l.Usecases.Coins.SyncPrice(context.Background(), req)
}
