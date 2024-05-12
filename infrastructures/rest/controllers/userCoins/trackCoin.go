package userCoins

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
)

func (c Controllers) TrackCoin(gc *gin.Context) {
	var req dto.TrackUntrackCoinRequest
	if err := req.FromGinContext(gc); err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data:    err,
		})

		return
	}

	resp, err := c.Usecases.UserCoins.TrackCoin(gc.Request.Context(), req)
	if err != nil {
		c.buildErrorTrack(gc, err)
		return
	}

	gc.JSON(http.StatusOK, dto.Default{
		Ok:      true,
		Message: "Success",
		Data:    dto.NewTrackCoinResponse(resp),
	})
}

func (c Controllers) buildErrorTrack(gc *gin.Context, err error) {
	statusCode := http.StatusInternalServerError

	if errors.Is(err, consts.ErrCoinHasBeenAdded) {
		statusCode = http.StatusBadRequest
	}

	gc.JSON(statusCode, dto.Default{
		Ok:      false,
		Message: err.Error(),
		Data: dto.ErrorData{
			Error: err.Error(),
		},
	})
}
