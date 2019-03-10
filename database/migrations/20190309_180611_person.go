package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Person_20190309_180611 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Person_20190309_180611{}
	m.Created = "20190309_180611"

	migration.Register("Person_20190309_180611", m)
}

// Run the migrations
func (m *Person_20190309_180611) Up() {
	// -- --------------------------------------------------
	// --  Table Structure for `github.com/school/models.Person`
	// -- --------------------------------------------------
	m.SQL(`CREATE TABLE person (
        id int unsigned AUTO_INCREMENT NOT NULL,
        dni int unsigned NOT NULL DEFAULT 0 ,
        first_name varchar(128) NOT NULL DEFAULT '' ,
        last_name varchar(128) NOT NULL DEFAULT '' ,
        birthdate datetime NOT NULL,
        inserted_at datetime NOT NULL,
		PRIMARY KEY (id),
		KEY dni_person (dni),
		KEY last_name (last_name),
		UNIQUE (dni, last_name)
    );`)

}

// Reverse the migrations
func (m *Person_20190309_180611) Down() {
	m.SQL("DROP TABLE person")
}
