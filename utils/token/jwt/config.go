package jwt

import "time"

type (
	config struct {
		expires   time.Duration
		publicKey string
		alg       string
	}

	JwtConfig func(c *config)
)

func WithExpires(exp time.Duration) JwtConfig {
	return func(c *config) {
		c.expires = exp
	}
}

func WithHS256(publicKey string) JwtConfig {
	return func(c *config) {
		c.alg = "HS256"
		c.publicKey = publicKey
	}
}
