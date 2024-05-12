package coincap

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/mockLogger"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Init(t *testing.T) {
	t.Run("If no error, it will return the repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		basePath := "/test/"
		log := mockLogger.NewMockLogger(ctrl)

		expRepo := coincap{
			basePath: basePath,
			log:      log,
		}

		repo, err := Init(basePath, log)

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

	basePath string
	log      *mockLogger.MockLogger

	coincap Coincap
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.basePath = "/test/"
	mock.log = mockLogger.NewMockLogger(mock.ctrl)

	mock.coincap = &coincap{
		basePath: mock.basePath,
		log:      mock.log,
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	mock.ctx = req.Context()

	return mock
}

func (m *mock) Finish() {
	m.ctrl.Finish()
}
