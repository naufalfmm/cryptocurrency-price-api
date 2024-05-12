package userCoins

import (
	"context"
	"errors"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"gorm.io/gorm"
)

func (u usecases) UntrackCoin(ctx context.Context, req dto.TrackUntrackCoinRequest) error {
	coin, err := u.persistents.Repositories.Coins.GetByCode(ctx, req.CoinSymbol)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return consts.ErrCoinMissing
		}

		return err
	}

	uc, err := u.persistents.Repositories.UserCoins.Get(ctx, req.UserSignIn.ID, coin.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return consts.ErrCoinTrackMissing
		}

		return err
	}

	if err := u.persistents.Repositories.UserCoins.DeleteByID(ctx, uc.ID, req.UserSignIn.CreatedBy()); err != nil {
		return err
	}

	return nil
}
