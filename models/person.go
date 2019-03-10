package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Person registers information about a persons.
type Person struct {
	ID         uint      `orm:"pk;auto;column(id)" json:"-"`
	DNI        uint      `orm:"column(dni)"`
	FirstName  string    `orm:"size(128)"`
	LastName   string    `orm:"size(128)"`
	Birthdate  time.Time `orm:"type(datetime)"`
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)" json:"inserted_at"`
	// email and some others fields

	Principal *Principal `orm:"reverse(one)"`
	Student   *Student   `orm:"reverse(one)"`
	Employee  *Employee  `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Person))
}

func (s *Person) TableName() string {
	return "person"
}

// multiple fields index
func (t *Person) TableIndex() [][]string {
	return [][]string{
		[]string{"LastName", "DNI"},
	}
}

// TableUnique _
func (t *Person) TableUnique() [][]string {
	return [][]string{
		[]string{"DNI", "LastName"},
	}
}

// AddPerson insert a new Person into database and returns
// last inserted ID on success.
func AddPerson(m *Person) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPersonById retrieves Person by ID. Returns error if
// ID doesn't exist
func GetPersonById(id uint) (v *Person, err error) {
	o := orm.NewOrm()
	v = &Person{ID: id}
	if err = o.QueryTable(new(Person)).Filter("ID", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPerson retrieves all Person matches certain condition. Returns empty list if
// no records exist
func GetAllPerson(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Person))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Person
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePerson updates Person by ID and returns error if
// the record to be updated doesn't exist
func UpdatePersonById(m *Person) (err error) {
	o := orm.NewOrm()
	v := Person{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePerson deletes Person by ID and returns error if
// the record to be deleted doesn't exist
func DeletePerson(id uint) (err error) {
	o := orm.NewOrm()
	v := Person{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Person{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
