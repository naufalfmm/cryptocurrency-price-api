package sqliteOrm

import "github.com/naufalfmm/cryptocurrency-price-api/utils/orm"

type QueryModifier func(o orm.Orm) orm.Orm
