package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
)

func (m middlewares) VerifyToken() gin.HandlerFunc {
	return func(gc *gin.Context) {
		tokenAuth := gc.Request.Header.Get("Authorization")
		if tokenAuth == "" {
			gc.JSON(http.StatusUnauthorized, dto.Default{
				Ok:      false,
				Message: consts.ErrInvalidToken.Error(),
				Data:    consts.ErrInvalidToken,
			})
			return
		}

		token := ""
		tokenSplit := strings.Split(tokenAuth, " ")
		if len(tokenSplit) > 1 {
			token = tokenSplit[1]
		}

		tokenData, err := m.jwt.Decoder.DecodeToken(token, &dto.LoginClaims{})
		if err != nil {
			gc.JSON(http.StatusUnauthorized, dto.Default{
				Ok:      false,
				Message: consts.ErrInvalidToken.Error(),
				Data:    consts.ErrInvalidToken,
			})
			return
		}

		userTokenData := tokenData.(*dto.LoginClaims)

		gc.Set(consts.XUserHeader, userTokenData.UserLogin)
		gc.Next()
	}
}
