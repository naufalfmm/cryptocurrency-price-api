package jwtoken

import (
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token/jwt"
)

func NewJWT(config *config.EnvConfig) (jwt.JWT, error) {
	confs := []jwt.JwtConfig{
		jwt.WithExpires(config.JwtExpires),
	}

	if config.JwtAlg == "HS256" {
		confs = append(confs, jwt.WithHS256(config.JwtAlg))
	}

	return jwt.NewJWT(confs...)
}
