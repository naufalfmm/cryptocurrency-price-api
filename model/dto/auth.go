package dto

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/frozenTime"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/validator"
)

type (
	UserSignIn struct {
		ID    uint64 `json:"id"`
		Email string `json:"email,omitempty"`
	}

	LoginClaims struct {
		*jwt.StandardClaims
		UserLogin UserSignIn `json:"user"`
	}
)

func (u UserSignIn) CreatedBy() string {
	return u.Email
}

func (c *LoginClaims) Valid() error {
	jwt.TimeFunc = time.Now

	return c.StandardClaims.Valid()
}

func (c *LoginClaims) SetExp(exp time.Duration) token.Claims {
	if c.StandardClaims == nil {
		c.StandardClaims = &jwt.StandardClaims{}
	}
	c.ExpiresAt = time.Now().Add(exp).Unix()

	return c
}

type SignUpRequest struct {
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

func (req *SignUpRequest) FromGinContext(gc *gin.Context) error {
	if err := gc.ShouldBindJSON(req); err != nil {
		return validator.NewFromValidationErrors(*req, err)
	}

	return nil
}

func (req SignUpRequest) ToUser(ctx context.Context) dao.User {
	return dao.User{
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: frozenTime.Now(ctx),
		UpdatedAt: frozenTime.Now(ctx),
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}
}

type SignUpResponse struct {
	UserSignIn
}

func NewSignUpResponse(u dao.User) SignUpResponse {
	return SignUpResponse{
		UserSignIn: UserSignIn{
			ID:    u.ID,
			Email: u.Email,
		},
	}
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (req *SignInRequest) FromGinContext(gc *gin.Context) error {
	if err := gc.BindJSON(req); err != nil {
		return validator.NewFromValidationErrors(*req, err)
	}

	return nil
}

type SignInResponse struct {
	JWT        string     `json:"jwt"`
	UserSignIn UserSignIn `json:"user"`
}

func NewSignInResponse(ul dao.UserSignIn) SignInResponse {
	return SignInResponse{
		JWT: ul.Token,
		UserSignIn: UserSignIn{
			ID:    ul.User.ID,
			Email: ul.User.Email,
		},
	}
}
