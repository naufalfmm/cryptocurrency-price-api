package orm

import "gorm.io/gorm"

type Association interface {
	Append(values ...interface{}) error
	Clear() error
	Count() (count int64)
	Delete(values ...interface{}) error
	Find(out interface{}, conds ...interface{}) error
	Replace(values ...interface{}) error
	Unscoped() *gorm.Association
}
