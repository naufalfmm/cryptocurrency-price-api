package userCoins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/queries"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/frozenTime"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/helper"
)

func (u usecases) updatePrice(ctx context.Context, coins []dao.Coin, userSignin dto.UserSignIn) {
	coinHistories := make([]dao.CoinHistory, len(coins))
	for i, coin := range coins {
		coinHistories[i] = dao.CoinHistory{
			CoinID:      coin.ID,
			LatestPrice: coin.LatestPrice,
			CreatedAt:   frozenTime.Now(ctx),
			CreatedBy:   userSignin.CreatedBy(),
		}
	}

	u.o.StartTransaction(ctx)
	defer u.o.RollbackTransaction(ctx)

	if err := u.persistents.Repositories.Coins.UpdatePrices(ctx, coins); err != nil {
		u.log.Error(ctx, consts.UpdatePriceLogMessage).Err(err).Any("coins", coins).Send()
	}

	if _, err := u.persistents.Repositories.CoinHistories.BulkCreate(ctx, coinHistories); err != nil {
		u.log.Error(ctx, consts.UpdatePriceLogMessage).Err(err).Any("coins", coins).Send()
	}

	u.o.CommitTransaction(ctx)
}

func (u usecases) convertCurrency(ctx context.Context, userCoins []dao.UserCoin, toCurrency string) ([]dao.UserCoin, error) {
	if toCurrency == "" {
		return userCoins, nil
	}

	ratesResp, err := u.persistents.Webclients.Coincap.GetAllRates(ctx)
	if err != nil {
		return userCoins, nil
	}

	currRate := ratesResp.Data.GetBySymbol(toCurrency)
	if currRate.Symbol == "" {
		return userCoins, nil
	}

	for i, userCoin := range userCoins {
		userCoins[i].Coin.LatestPrice = userCoin.Coin.LatestPrice / helper.DefaultConvertFloat64(currRate.RateUSD, 1)
		userCoins[i].Coin.LatestPriceCurrency = currRate.Symbol
	}

	return userCoins, nil
}

func (u usecases) GetAllTrack(ctx context.Context, req dto.GetAllTrackRequest) ([]dao.UserCoin, error) {
	ucs, err := u.persistents.Repositories.UserCoins.GetAll(ctx, dto.GetAllRequest{
		UserID: req.UserSignIn.ID,
	}, queries.GetAllUserCoins)
	if err != nil {
		return nil, err
	}

	if u.conf.CoincapPriceSyncMode {
		ucs, err := u.convertCurrency(ctx, ucs, req.Currency)
		if err != nil {
			return nil, err
		}

		return ucs, nil
	}

	ccReq := dto.AllAssetsCoincapRequest{}
	for i, uc := range ucs {
		if i == 0 {
			ccReq.IDs = uc.Coin.CoincapID
			continue
		}

		ccReq.IDs += ("," + uc.Coin.CoincapID)
	}

	assetResp, err := u.persistents.Webclients.Coincap.GetAllAssets(ctx, ccReq)
	if err != nil {
		return ucs, nil
	}

	assetPriceMap := make(map[string]string)
	for _, asset := range assetResp.Data {
		assetPriceMap[asset.ID] = asset.PriceUSD
	}

	updatedCoins := make([]dao.Coin, len(ucs))
	for i, uc := range ucs {
		if _, ok := assetPriceMap[uc.Coin.CoincapID]; !ok {
			continue
		}

		ucs[i].Coin.LatestPrice = helper.DefaultConvertFloat64(assetPriceMap[uc.Coin.CoincapID], 0)
		updatedCoins[i] = dao.Coin{
			CoincapID:   ucs[i].Coin.CoincapID,
			LatestPrice: ucs[i].Coin.LatestPrice,
			UpdatedAt:   frozenTime.Now(ctx),
			UpdatedBy:   req.UserSignIn.CreatedBy(),
		}
	}

	go u.updatePrice(context.Background(), updatedCoins, req.UserSignIn)

	ucs, err = u.convertCurrency(ctx, ucs, req.Currency)
	if err != nil {
		return nil, err
	}

	return ucs, nil
}
