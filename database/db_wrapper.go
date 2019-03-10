package database

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type DbWrapper struct {
	db orm.Ormer

	DisableCommit bool
}

func (d *DbWrapper) All(tableName string, objects interface{}, cols ...string) (int64, error) {
	return d.db.QueryTable(tableName).All(objects, cols...)
}

func (d *DbWrapper) Read(object interface{}, cols ...string) error {
	return d.db.Read(object, cols...)
}

func (d *DbWrapper) Insert(object interface{}) (int64, error) {
	return d.db.Insert(object)
}

func (d *DbWrapper) Update(object interface{}) (int64, error) {
	return d.db.Update(object)
}

func (d *DbWrapper) Delete(object interface{}, cols ...string) (int64, error) {
	return d.db.Delete(object, cols...)
}

func (d *DbWrapper) Raw(query string, args ...interface{}) orm.RawSeter {
	return d.db.Raw(query, args...)
}

func (d *DbWrapper) LoadRelated(object interface{}, relation string, args ...interface{}) (int64, error) {
	return d.db.LoadRelated(object, relation, args)
}

func (d *DbWrapper) Begin() error {
	return d.db.Begin()
}

func (d *DbWrapper) Commit() error {
	if d.DisableCommit {
		return nil
	}

	return d.db.Commit()
}

func (d *DbWrapper) Rollback() error {
	return d.db.Rollback()
}

func (d *DbWrapper) GetQueryTable(model string) orm.QuerySeter {
	return d.db.QueryTable(model)
}

func (d *DbWrapper) NewCondition() *orm.Condition {
	return orm.NewCondition()
}

func NewDbWrapper() DbInterface {
	return &DbWrapper{db: orm.NewOrm()}
}

func NewDbWrapperWithSettings(settings map[string]string) DbInterface {
	if ok, _ := strconv.ParseBool(settings["disableCommit"]); ok {
		return &DbWrapper{db: orm.NewOrm(), DisableCommit: ok}
	}
	return &DbWrapper{db: orm.NewOrm()}
}
