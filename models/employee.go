package models

import "github.com/astaxie/beego/orm"

type Employee struct {
	ID       int64 `orm:"pk;auto;column(id)"`
	Rol      string
	Person   *Person    `orm:"column(person_id);rel(one)"`
	Director *Principal `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Employee))
}

func (e *Employee) TableName() string {
	return "employee"
}
