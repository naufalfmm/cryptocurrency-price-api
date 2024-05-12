package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
)

func (c *Controllers) SignIn(gc *gin.Context) {
	var req dto.SignInRequest
	if err := req.FromGinContext(gc); err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data:    err,
		})

		return
	}

	resp, err := c.Usecases.Auth.SignIn(gc.Request.Context(), req)
	if err != nil {
		c.buildErrorSignIn(gc, err)
		return
	}

	gc.JSON(http.StatusOK, dto.Default{
		Ok:      true,
		Message: "Success",
		Data:    dto.NewSignInResponse(resp),
	})
}

func (c Controllers) buildErrorSignIn(gc *gin.Context, err error) {
	statusCode := http.StatusInternalServerError

	if errors.Is(err, consts.ErrWrongPassword) ||
		errors.Is(err, consts.ErrEmailMissing) {
		statusCode = http.StatusForbidden
	}

	gc.JSON(statusCode, dto.Default{
		Ok:      false,
		Message: err.Error(),
		Data: dto.ErrorData{
			Error: err.Error(),
		},
	})
}
