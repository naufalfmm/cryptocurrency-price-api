package dao

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/utils/frozenTime"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/helper"
)

type (
	AllAssetCoin struct {
		ID       string `json:"id"`
		Symbol   string `json:"symbol"`
		Name     string `json:"name"`
		PriceUSD string `json:"priceUsd"`
	}

	AllAssetCoins []AllAssetCoin

	AllAsset struct {
		Data AllAssetCoins `json:"data"`
	}
)

func (aacs AllAssetCoins) GetBySymbol(symbol string) AllAssetCoin {
	for _, aac := range aacs {
		if aac.Symbol == symbol {
			return aac
		}
	}

	return AllAssetCoin{}
}

func (c AllAssetCoin) ToCoin(ctx context.Context, email string) Coin {
	return Coin{
		Code:        c.Symbol,
		CoincapID:   c.ID,
		Name:        c.Name,
		LatestPrice: helper.DefaultConvertFloat64(c.PriceUSD),
		CreatedAt:   frozenTime.Now(ctx),
		UpdatedAt:   frozenTime.Now(ctx),
		CreatedBy:   email,
		UpdatedBy:   email,
	}
}
