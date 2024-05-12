package repositories

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/coinHistories"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/coins"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/userCoins"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/users"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/mockLogger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/mockOrm"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Init(t *testing.T) {
	t.Run("If no error, it will return the repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orm := mockOrm.NewMockOrm(ctrl)
		log := mockLogger.NewMockLogger(ctrl)

		db := db.DB{
			Orm: orm,
		}

		u, _ := users.Init(&db, log)

		c, _ := coins.Init(&db, log)

		uc, _ := userCoins.Init(&db, log)

		ch, _ := coinHistories.Init(&db, log)

		expRepo := Repositories{
			Users:         u,
			Coins:         c,
			UserCoins:     uc,
			CoinHistories: ch,
		}

		repo, err := Init(&db, log)

		assert.Nil(t, err)
		assert.Equal(t, &expRepo, repo)
	})
}
