package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/helper"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/validator"
)

type TrackUntrackCoinRequest struct {
	CoinSymbol string `json:"coin_symbol"`

	UserSignIn UserSignIn `json:"-"`
}

func (req *TrackUntrackCoinRequest) FromGinContext(gc *gin.Context) error {
	if err := gc.ShouldBindJSON(req); err != nil {
		return validator.NewFromValidationErrors(*req, err)
	}

	userHeader, _ := gc.Get(consts.XUserHeader)
	req.UserSignIn = userHeader.(UserSignIn)

	return nil
}

type (
	TrackCoinResponse struct {
		TrackID             uint64    `json:"track_id"`
		Code                string    `json:"code"`
		Name                string    `json:"name"`
		LatestPrice         float64   `json:"latest_price"`
		LatestPriceCurrency string    `json:"latest_price_currency"`
		AddedAt             time.Time `json:"added_at"`
		AddedBy             string    `json:"added_by"`
	}

	TrackCoinResponses []TrackCoinResponse
)

func NewTrackCoinResponse(uc dao.UserCoin) TrackCoinResponse {
	return TrackCoinResponse{
		TrackID:             uc.ID,
		Code:                uc.Coin.Code,
		Name:                uc.Coin.Name,
		LatestPrice:         uc.Coin.LatestPrice,
		LatestPriceCurrency: helper.DefaultIfEmpty(uc.Coin.LatestPriceCurrency, consts.UsdCurrency),
		AddedAt:             uc.CreatedAt,
		AddedBy:             uc.CreatedBy,
	}
}

func NewTrackCoinResponses(ucs []dao.UserCoin) TrackCoinResponses {
	resps := make(TrackCoinResponses, len(ucs))
	for i, uc := range ucs {
		resps[i] = NewTrackCoinResponse(uc)
	}

	return resps
}

type GetAllRequest struct {
	UserID uint64
}

type GetAllTrackRequest struct {
	Currency   string
	UserSignIn UserSignIn `json:"-"`
}

func (req *GetAllTrackRequest) FromGinContext(gc *gin.Context) error {
	userHeader, _ := gc.Get(consts.XUserHeader)
	userSignIn := userHeader.(UserSignIn)

	*req = GetAllTrackRequest{
		Currency:   gc.Request.Header.Get(consts.XCurrencyHeader),
		UserSignIn: userSignIn,
	}

	return nil
}
