package persistents

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/webclients"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/mockLogger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/mockOrm"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Init(t *testing.T) {
	t.Run("If no error, it will return the repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		conf := config.EnvConfig{
			CoincapBasePath: "/test/",
		}
		orm := mockOrm.NewMockOrm(ctrl)
		log := mockLogger.NewMockLogger(ctrl)

		db := db.DB{
			Orm: orm,
		}

		w, _ := webclients.Init(&conf, log)
		r, _ := repositories.Init(&db, log)

		expPer := Persistents{
			Repositories: r,
			Webclients:   w,
		}

		per, err := Init(&db, log, &conf)

		assert.Nil(t, err)
		assert.Equal(t, expPer, per)
	})
}
