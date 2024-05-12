package webclients

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/webclients/coincap"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/mockLogger"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Init(t *testing.T) {
	t.Run("If no error, it will return the repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		conf := config.EnvConfig{
			CoincapBasePath: "/test/",
		}
		log := mockLogger.NewMockLogger(ctrl)

		c, _ := coincap.Init(conf.CoincapBasePath, log)

		expWcl := Webclients{
			Coincap: c,
		}

		wcl, err := Init(&conf, log)

		assert.Nil(t, err)
		assert.Equal(t, expWcl, wcl)
	})
}
