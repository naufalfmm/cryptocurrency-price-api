package users

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
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

		expRepo := repositories{
			orm: &db,
			log: log,
		}

		repo, err := Init(&db, log)

		assert.Nil(t, err)
		assert.Equal(t, &expRepo, repo)
	})
}

var (
	errAny = errors.New("any error")
)

type mock struct {
	ctrl *gomock.Controller
	ctx  context.Context

	orm *mockOrm.MockOrm
	db  *db.DB
	log *mockLogger.MockLogger

	repositories Repositories
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.orm = mockOrm.NewMockOrm(mock.ctrl)
	mock.log = mockLogger.NewMockLogger(mock.ctrl)

	mock.db = &db.DB{
		Orm: mock.orm,
	}

	mock.repositories = &repositories{
		orm: mock.db,
		log: mock.log,
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	mock.ctx = req.Context()

	return mock
}

func (m *mock) Finish() {
	m.ctrl.Finish()
}
