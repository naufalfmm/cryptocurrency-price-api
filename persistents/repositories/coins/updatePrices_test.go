package coins

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Test_repositories_UpdatePrices(t *testing.T) {
	var (
		coins = []dao.Coin{
			{
				CoincapID: "ethereum",
			},
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "coincap_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"latest_price": gorm.Expr("excluded.latest_price"),
				"updated_at":   gorm.Expr("excluded.updated_at"),
				"updated_by":   gorm.Expr("excluded.updated_by"),
			}),
		}).Return(mock.orm)
		mock.orm.EXPECT().Create(&coins).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(nil)

		err := mock.repositories.UpdatePrices(mock.ctx, coins)

		assert.Nil(t, err)
	})

	t.Run("If error exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "coincap_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"latest_price": gorm.Expr("excluded.latest_price"),
				"updated_at":   gorm.Expr("excluded.updated_at"),
				"updated_by":   gorm.Expr("excluded.updated_by"),
			}),
		}).Return(mock.orm)
		mock.orm.EXPECT().Create(&coins).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "update-coin-prices").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Send()

		err := mock.repositories.UpdatePrices(mock.ctx, coins)

		assert.Equal(t, errAny, err)
	})
}
