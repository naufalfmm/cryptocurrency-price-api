package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
)

func (c Controllers) SignUp(gc *gin.Context) {
	var req dto.SignUpRequest
	if err := req.FromGinContext(gc); err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data:    err,
		})

		return
	}

	user, err := c.Usecases.Auth.SignUp(gc.Request.Context(), req)
	if err != nil {
		c.buildErrorSignUp(gc, err)
		return
	}

	gc.JSON(http.StatusOK, dto.Default{
		Ok:      true,
		Message: "Success",
		Data:    dto.NewSignUpResponse(user),
	})
}

func (c Controllers) buildErrorSignUp(gc *gin.Context, err error) {
	statusCode := http.StatusInternalServerError

	if errors.Is(err, consts.ErrEmailHasBeenUsed) {
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
