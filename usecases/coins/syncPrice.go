package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/frozenTime"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/helper"
)

func (u usecases) SyncPrice(ctx context.Context, req dto.SyncCoinPriceRequest) {
	coins, err := u.persistents.Repositories.Coins.GetByCoincapIDs(ctx, req.ToCoincapIDs())
	if err != nil {
		u.log.Error(ctx, consts.SyncPriceLogMessage).Err(err).Any(consts.ReqLogKey, req)
	}

	updatedCoins := make([]dao.Coin, len(coins))
	coinHistories := make([]dao.CoinHistory, len(coins))
	for i, coin := range coins {
		coinPrice, ok := req.CoinPriceMap[coin.CoincapID]
		if !ok {
			continue
		}

		updatedCoins[i] = dao.Coin{
			CoincapID:   coin.CoincapID,
			LatestPrice: helper.DefaultConvertFloat64(coinPrice, 0),
			UpdatedAt:   frozenTime.Now(ctx),
			UpdatedBy:   consts.SystemCreatedBy,
		}

		coinHistories[i] = dao.CoinHistory{
			CoinID:      coin.ID,
			LatestPrice: helper.DefaultConvertFloat64(coinPrice, 0),
			CreatedAt:   frozenTime.Now(ctx),
			CreatedBy:   consts.SystemCreatedBy,
		}
	}

	if err := u.persistents.Repositories.Coins.UpdatePrices(ctx, updatedCoins); err != nil {
		u.log.Error(ctx, consts.SyncPriceLogMessage).Err(err).Any(consts.ReqLogKey, req)
	}

	if _, err := u.persistents.Repositories.CoinHistories.BulkCreate(ctx, coinHistories); err != nil {
		u.log.Error(ctx, consts.SyncPriceLogMessage).Err(err).Any(consts.ReqLogKey, req)
	}
}
