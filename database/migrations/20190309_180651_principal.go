package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Principal_20190309_180651 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Principal_20190309_180651{}
	m.Created = "20190309_180651"

	migration.Register("Principal_20190309_180651", m)
}

// Run the migrations
func (m *Principal_20190309_180651) Up() {
	// -- --------------------------------------------------
	// --  Table Structure for `github.com/school/models.Principal`
	// -- --------------------------------------------------
	m.SQL(`CREATE TABLE IF NOT EXISTS principal (
        id int unsigned AUTO_INCREMENT NOT NULL,
		person_id int unsigned,
		PRIMARY KEY (id),
		CONSTRAINT principal_preson_fk FOREIGN KEY (person_id) REFERENCES person (id)
    );`)
}

// Reverse the migrations
func (m *Principal_20190309_180651) Down() {
	m.SQL("DROP TABLE principal")

}
