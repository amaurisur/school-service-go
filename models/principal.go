package models

import "github.com/astaxie/beego/orm"

type Principal struct {
	ID         int64       `orm:"pk;auto;column(id)"`
	Person     *Person     `orm:"column(person_id);rel(one)"`
	Supervised []*Employee `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Principal))
}

func (s *Principal) TableName() string {
	return "principal"
}
