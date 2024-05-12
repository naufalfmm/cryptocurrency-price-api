package jwt

import "github.com/naufalfmm/cryptocurrency-price-api/utils/token"

type JWT struct {
	Encoder token.Encoder
	Decoder token.Decoder

	conf config
}

func NewJWT(configs ...JwtConfig) (JWT, error) {
	c := config{}
	for _, conf := range configs {
		conf(&c)
	}

	encoder := NewEncoder(c.publicKey, c.alg, c.expires)
	decoder := NewDecoder(c.publicKey)

	return JWT{
		Encoder: encoder,
		Decoder: decoder,

		conf: c,
	}, nil
}
