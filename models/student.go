package models

import (
	"github.com/astaxie/beego/orm"
)

type Student struct {
	ID       uint       `orm:"pk;auto;column(id)"`
	Person   *Person    `orm:"column(person_id);rel(one)"`
	Tutor    *Person    `orm:"column(tutor_id);rel(one)"`
	Subjects []*Subject `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Student))
}

func (s *Student) TableName() string {
	return "student"
}
