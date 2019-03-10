package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Subject_20190309_180706 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Subject_20190309_180706{}
	m.Created = "20190309_180706"

	migration.Register("Subject_20190309_180706", m)
}

// Run the migrations
func (m *Subject_20190309_180706) Up() {
	// -- --------------------------------------------------
	// --  Table Structure for `github.com/school/models.Subject`
	// -- --------------------------------------------------
	m.SQL(`CREATE TABLE subject (
        id int unsigned AUTO_INCREMENT NOT NULL,
        title varchar(128) NOT NULL DEFAULT '' ,
		description varchar(128),
		PRIMARY KEY (id)
    );`)

}

// Reverse the migrations
func (m *Subject_20190309_180706) Down() {
	m.SQL("DROP TABLE subject")
}
