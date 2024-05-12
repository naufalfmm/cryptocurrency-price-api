package userCoins

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
)

func (c Controllers) GetAllTrack(gc *gin.Context) {
	var req dto.GetAllTrackRequest
	if err := req.FromGinContext(gc); err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data:    err,
		})

		return
	}

	resp, err := c.Usecases.UserCoins.GetAllTrack(gc.Request.Context(), req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data: dto.ErrorData{
				Error: err.Error(),
			},
		})
		return
	}

	gc.JSON(http.StatusOK, dto.Default{
		Ok:      true,
		Message: "Success",
		Data: dto.ItemData{
			Items: dto.NewTrackCoinResponses(resp),
		},
	})
}
