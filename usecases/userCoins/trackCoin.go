package userCoins

import (
	"context"
	"errors"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/frozenTime"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
	"gorm.io/gorm"
)

func (u usecases) TrackCoin(ctx context.Context, req dto.TrackUntrackCoinRequest) (dao.UserCoin, error) {
	coin, err := u.getCreateCoin(ctx, req)
	if err != nil {
		return dao.UserCoin{}, err
	}

	userCoin, err := u.persistents.Repositories.UserCoins.Create(ctx, dao.UserCoin{
		CoinID:    coin.ID,
		UserID:    req.UserSignIn.ID,
		CreatedAt: frozenTime.Now(ctx),
		UpdatedAt: frozenTime.Now(ctx),
		CreatedBy: req.UserSignIn.Email,
		UpdatedBy: req.UserSignIn.Email,
	})
	if err != nil {
		if sqliteOrm.IsUniqueConstraintError(err) {
			return dao.UserCoin{}, consts.ErrCoinHasBeenAdded
		}

		return dao.UserCoin{}, err
	}

	userCoin.Coin = coin

	return userCoin, nil
}

func (u usecases) getCreateCoin(ctx context.Context, req dto.TrackUntrackCoinRequest) (dao.Coin, error) {
	coin, err := u.persistents.Repositories.Coins.GetByCode(ctx, req.CoinSymbol)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dao.Coin{}, err
	}

	if coin.ID != 0 {
		return coin, nil
	}

	assets, err := u.persistents.Webclients.Coincap.GetAllAssets(ctx, dto.AllAssetsCoincapRequest{
		Search: req.CoinSymbol,
		Limit:  1,
	})
	if err != nil {
		return dao.Coin{}, err
	}

	asset := assets.Data.GetBySymbol(req.CoinSymbol)
	if asset.ID == "" {
		return dao.Coin{}, consts.ErrCoinSymbolMissing
	}

	coin, err = u.persistents.Repositories.Coins.Create(ctx, asset.ToCoin(ctx, req.UserSignIn.Email))
	if err != nil {
		return dao.Coin{}, err
	}

	return coin, err
}
