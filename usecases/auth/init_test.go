package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/mockLogger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/password/mockPassword"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token/jwt"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token/mockToken"
	"github.com/stretchr/testify/assert"

	mockCoinHistory "github.com/naufalfmm/cryptocurrency-price-api/mocks/persistents/repositories/coinHistories"
	mockCoin "github.com/naufalfmm/cryptocurrency-price-api/mocks/persistents/repositories/coins"
	mockUserCoin "github.com/naufalfmm/cryptocurrency-price-api/mocks/persistents/repositories/userCoins"
	mockUser "github.com/naufalfmm/cryptocurrency-price-api/mocks/persistents/repositories/users"
)

func Test_usecases_Init(t *testing.T) {
	t.Run("If no error, it will return the repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		log := mockLogger.NewMockLogger(ctrl)
		pwd := mockPassword.NewMockPassword(ctrl)
		enc := mockToken.NewMockEncoder(ctrl)
		dec := mockToken.NewMockDecoder(ctrl)
		persist := persistents.Persistents{
			Repositories: repositories.Repositories{
				Users:         mockUser.NewMockRepositories(ctrl),
				Coins:         mockCoin.NewMockRepositories(ctrl),
				UserCoins:     mockUserCoin.NewMockRepositories(ctrl),
				CoinHistories: mockCoinHistory.NewMockRepositories(ctrl),
			},
		}

		jwt := jwt.JWT{
			Encoder: enc,
			Decoder: dec,
		}

		expUsec := usecases{
			persistents: persist,
			log:         log,
			pwd:         pwd,
			jwt:         jwt,
		}

		usec, err := Init(persist, log, pwd, jwt)

		assert.Nil(t, err)
		assert.Equal(t, &expUsec, usec)
	})
}

var (
	errAny = errors.New("any error")
)

type mock struct {
	ctrl *gomock.Controller
	ctx  context.Context

	user        *mockUser.MockRepositories
	coin        *mockCoin.MockRepositories
	userCoin    *mockUserCoin.MockRepositories
	coinHistory *mockCoinHistory.MockRepositories
	log         *mockLogger.MockLogger
	pwd         *mockPassword.MockPassword
	enc         *mockToken.MockEncoder
	dec         *mockToken.MockDecoder

	persistent persistents.Persistents
	jwt        jwt.JWT

	usecases Usecases
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.user = mockUser.NewMockRepositories(mock.ctrl)
	mock.coin = mockCoin.NewMockRepositories(mock.ctrl)
	mock.userCoin = mockUserCoin.NewMockRepositories(mock.ctrl)
	mock.coinHistory = mockCoinHistory.NewMockRepositories(mock.ctrl)
	mock.log = mockLogger.NewMockLogger(mock.ctrl)
	mock.pwd = mockPassword.NewMockPassword(mock.ctrl)
	mock.enc = mockToken.NewMockEncoder(mock.ctrl)
	mock.dec = mockToken.NewMockDecoder(mock.ctrl)

	mock.persistent = persistents.Persistents{
		Repositories: repositories.Repositories{
			Users:         mock.user,
			Coins:         mock.coin,
			UserCoins:     mock.userCoin,
			CoinHistories: mock.coinHistory,
		},
	}

	mock.jwt = jwt.JWT{
		Encoder: mock.enc,
		Decoder: mock.dec,
	}

	mock.usecases = &usecases{
		persistents: mock.persistent,
		log:         mock.log,
		pwd:         mock.pwd,
		jwt:         mock.jwt,
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	mock.ctx = req.Context()

	return mock
}

func (m *mock) Finish() {
	m.ctrl.Finish()
}
