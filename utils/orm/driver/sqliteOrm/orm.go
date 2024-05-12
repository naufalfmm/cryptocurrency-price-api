package sqliteOrm

import (
	"context"
	"database/sql"

	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type slite struct {
	g    *gorm.DB
	name string
}

func NewOrm(g *gorm.DB, name string) (slite, error) {
	return slite{
		g:    g,
		name: name,
	}, nil
}

func (o *slite) newImpl(g *gorm.DB) orm.Orm {
	return &slite{
		g:    g,
		name: o.name,
	}
}

func (o *slite) Gorm() *gorm.DB {
	return o.g
}

func (o *slite) Error() error {
	return o.g.Error
}

func (o *slite) Clone() orm.Orm {
	return &slite{
		g:    &(*o.g),
		name: o.name,
	}
}

func (o *slite) Ping() error {
	db, err := o.g.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

func (o *slite) AddError(err error) error {
	return o.g.AddError(err)
}

func (o *slite) Assign(attrs ...interface{}) orm.Orm {
	return o.newImpl(o.g.Assign(attrs...))
}

func (o *slite) Association(column string) orm.Association {
	return o.g.Association(column)
}

func (o *slite) Attrs(attrs ...interface{}) orm.Orm {
	return o.newImpl(o.g.Attrs(attrs...))
}

func (o *slite) AutoMigrate(dst ...interface{}) error {
	return o.g.AutoMigrate()
}

func (o *slite) Begin(opts ...*sql.TxOptions) orm.Orm {
	return o.newImpl(o.g.Begin(opts...))
}

func (o *slite) Clauses(conds ...clause.Expression) orm.Orm {
	return o.newImpl(o.g.Clauses(conds...))
}

func (o *slite) Commit() orm.Orm {
	return o.newImpl(o.g.Commit())
}

func (o *slite) Connection(fc func(tx orm.Orm) error) error {
	return o.g.Connection(func(tx *gorm.DB) error {
		return fc(o)
	})
}

func (o *slite) Count(count *int64) orm.Orm {
	return o.newImpl(o.g.Count(count))
}

func (o *slite) Create(value interface{}) orm.Orm {
	return o.newImpl(o.g.Create(value))
}

func (o *slite) CreateInBatches(value interface{}, batchSize int) orm.Orm {
	return o.newImpl(o.g.CreateInBatches(value, batchSize))
}

func (o *slite) DB() (*sql.DB, error) {
	return o.g.DB()
}

func (o *slite) Debug() orm.Orm {
	return o.newImpl(o.g.Debug())
}

func (o *slite) Delete(value interface{}, conds ...interface{}) orm.Orm {
	return o.newImpl(o.g.Delete(value, conds...))
}

func (o *slite) Distinct(args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Distinct(args...))
}

func (o *slite) Exec(sql string, values ...interface{}) orm.Orm {
	return o.newImpl(o.g.Exec(sql, values...))
}

func (o *slite) Find(dest interface{}, conds ...interface{}) orm.Orm {
	return o.newImpl(o.g.Find(dest, conds...))
}

func (o *slite) FindInBatches(dest interface{}, batchSize int, fc func(tx orm.Orm, batch int) error) orm.Orm {
	return o.newImpl(o.g.FindInBatches(dest, batchSize, func(tx *gorm.DB, batch int) error {
		return fc(o, batch)
	}))
}

func (o *slite) First(dest interface{}, conds ...interface{}) orm.Orm {
	return o.newImpl(o.g.First(dest, conds...))
}

func (o *slite) FirstOrCreate(dest interface{}, conds ...interface{}) orm.Orm {
	return o.newImpl(o.g.FirstOrCreate(dest, conds...))
}

func (o *slite) FirstOrInit(dest interface{}, conds ...interface{}) orm.Orm {
	return o.newImpl(o.g.FirstOrInit(dest, conds...))
}

func (o *slite) Get(key string) (interface{}, bool) {
	return o.g.Get(key)
}

func (o *slite) Group(name string) orm.Orm {
	return o.newImpl(o.g.Group(name))
}

func (o *slite) Having(query interface{}, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Having(query, args...))
}

func (o *slite) InnerJoins(query string, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.InnerJoins(query, args...))
}

func (o *slite) InstanceGet(key string) (interface{}, bool) {
	return o.g.InstanceGet(key)
}

func (o *slite) InstanceSet(key string, value interface{}) orm.Orm {
	return o.newImpl(o.g.InstanceSet(key, value))
}

func (o *slite) Joins(query string, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Joins(query, args...))
}

func (o *slite) Last(dest interface{}, conds ...interface{}) orm.Orm {
	return o.newImpl(o.g.Last(dest, conds...))
}

func (o *slite) Limit(limit int) orm.Orm {
	return o.newImpl(o.g.Limit(limit))
}

func (o *slite) Migrator() gorm.Migrator {
	return o.g.Migrator()
}

func (o *slite) Model(value interface{}) orm.Orm {
	return o.newImpl(o.g.Model(value))
}

func (o *slite) Not(query interface{}, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Not(query, args...))
}

func (o *slite) Offset(offset int) orm.Orm {
	return o.newImpl(o.g.Offset(offset))
}

func (o *slite) Omit(columns ...string) orm.Orm {
	return o.newImpl(o.g.Omit(columns...))
}

func (o *slite) Or(query interface{}, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Or(query, args...))
}

func (o *slite) Order(value interface{}) orm.Orm {
	return o.newImpl(o.g.Order(value))
}

func (o *slite) Pluck(column string, dest interface{}) orm.Orm {
	return o.newImpl(o.g.Pluck(column, dest))
}

func (o *slite) Preload(query string, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Preload(query, args...))
}

func (o *slite) Raw(sql string, values ...interface{}) orm.Orm {
	return o.newImpl(o.g.Raw(sql, values...))
}

func (o *slite) Rollback() orm.Orm {
	return o.newImpl(o.g.Rollback())
}

func (o *slite) RollbackTo(name string) orm.Orm {
	return o.newImpl(o.g.RollbackTo(name))
}

func (o *slite) Row() *sql.Row {
	return o.g.Row()
}

func (o *slite) Rows() (*sql.Rows, error) {
	return o.g.Rows()
}

func (o *slite) Save(value interface{}) orm.Orm {
	return o.newImpl(o.g.Save(value))
}

func (o *slite) SavePoint(name string) orm.Orm {
	return o.newImpl(o.g.SavePoint(name))
}

func (o *slite) Scan(dest interface{}) orm.Orm {
	return o.newImpl(o.g.Scan(dest))
}

func (o *slite) ScanRows(rows *sql.Rows, dest interface{}) error {
	return o.g.ScanRows(rows, dest)
}

func (o *slite) Scopes(funcs ...func(tx orm.Orm) orm.Orm) orm.Orm {
	newFuncs := make([]func(tx *gorm.DB) *gorm.DB, len(funcs))
	for i := 0; i < len(funcs); i++ {
		newFuncs[i] = func(tx *gorm.DB) *gorm.DB {
			return funcs[i](o).Gorm()
		}
	}

	return o.newImpl(o.g.Scopes(newFuncs...))
}

func (o *slite) Select(query interface{}, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Select(query, args...))
}

func (o *slite) Session(config *gorm.Session) orm.Orm {
	return o.newImpl(o.g.Session(config))
}

func (o *slite) Set(key string, value interface{}) orm.Orm {
	return o.newImpl(o.g.Set(key, value))
}

func (o *slite) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return o.g.SetupJoinTable(model, field, joinTable)
}

func (o *slite) Table(name string, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Table(name, args...))
}

func (o *slite) Take(dest interface{}, conds ...interface{}) orm.Orm {
	return o.newImpl(o.g.Take(dest, conds...))
}

func (o *slite) ToSQL(queryFn func(tx orm.Orm) orm.Orm) string {
	return o.g.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return queryFn(o).Gorm()
	})
}

func (o *slite) Transaction(fc func(tx orm.Orm) error, opts ...*sql.TxOptions) error {
	return o.g.Transaction(func(tx *gorm.DB) error {
		return fc(o)
	}, opts...)
}

func (o *slite) Unscoped() orm.Orm {
	return o.newImpl(o.g.Unscoped())
}

func (o *slite) Update(column string, value interface{}) orm.Orm {
	return o.newImpl(o.g.Update(column, value))
}

func (o *slite) UpdateColumn(column string, value interface{}) orm.Orm {
	return o.newImpl(o.g.UpdateColumn(column, value))
}

func (o *slite) UpdateColumns(values interface{}) orm.Orm {
	return o.newImpl(o.g.UpdateColumns(values))
}

func (o *slite) Updates(values interface{}) orm.Orm {
	return o.newImpl(o.g.Updates(values))
}

func (o *slite) Use(plugin gorm.Plugin) error {
	return o.g.Use(plugin)
}

func (o *slite) Where(query interface{}, args ...interface{}) orm.Orm {
	return o.newImpl(o.g.Where(query, args...))
}

func (o *slite) WithContext(ctx context.Context) orm.Orm {
	return o.newImpl(o.g.WithContext(ctx))
}
