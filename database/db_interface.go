package database

import (
	"github.com/astaxie/beego/orm"
)

type DbInterface interface {
	All(tableName string, objects interface{}, cols ...string) (int64, error)
	Read(object interface{}, cols ...string) error
	Insert(object interface{}) (int64, error)
	Update(object interface{}) (int64, error)
	Delete(object interface{}, cols ...string) (int64, error)
	Raw(query string, args ...interface{}) orm.RawSeter
	LoadRelated(object interface{}, relation string, args ...interface{}) (int64, error)
	Begin() error
	Commit() error
	Rollback() error
	NewCondition() *orm.Condition

	// Filtering methods
	GetQueryTable(model string) orm.QuerySeter
}
