package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Employee_20190309_180638 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Employee_20190309_180638{}
	m.Created = "20190309_180638"

	migration.Register("Employee_20190309_180638", m)
}

// Run the migrations
func (m *Employee_20190309_180638) Up() {
	// -- --------------------------------------------------
	// --  Table Structure for `github.com/school/models.Employee`
	// -- --------------------------------------------------
	m.SQL(`CREATE TABLE employee (
		id int unsigned AUTO_INCREMENT NOT NULL,
		rol varchar(255),
        person_id int unsigned NOT NULL,
		director_id int unsigned,
		PRIMARY KEY (id),
		CONSTRAINT employee_preson_fk FOREIGN KEY (person_id) REFERENCES person (id),
		CONSTRAINT employee_ppal_fk FOREIGN KEY (director_id) REFERENCES person (id)
	);`)
}

// Reverse the migrations
func (m *Employee_20190309_180638) Down() {
	m.SQL("DROP TABLE employee")

}
