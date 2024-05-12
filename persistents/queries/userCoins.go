package queries

import (
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
)

var GetAllUserCoins sqliteOrm.QueryModifier = func(o orm.Orm) orm.Orm {
	return o.
		InnerJoins("Coin")
}
