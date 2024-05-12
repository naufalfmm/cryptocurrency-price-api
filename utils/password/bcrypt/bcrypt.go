package bcrypt

import (
	"github.com/naufalfmm/cryptocurrency-price-api/utils/password"
	"golang.org/x/crypto/bcrypt"
)

type bcryptPassword struct {
	config config
}

func NewBcrypt(confs ...BcryptConfig) (password.Password, error) {
	config := config{}

	for _, conf := range confs {
		conf(&config)
	}

	return &bcryptPassword{
		config: config,
	}, nil
}

func (b *bcryptPassword) Generate(password string) (string, error) {
	bytesGenPass, err := bcrypt.GenerateFromPassword([]byte(password), b.config.cost)
	if err != nil {
		return "", err
	}

	return string(bytesGenPass), nil
}

func (b *bcryptPassword) Check(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
