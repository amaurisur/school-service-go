package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Tables_20190309_015316 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Tables_20190309_015316{}
	m.Created = "20190309_015316"

	migration.Register("Tables_20190309_015316", m)
}

// Run the migrations
func (m *Tables_20190309_015316) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE person (
		id int NOT NULL AUTO_INCREMENT,
		dni int NOT NULL,
		first_name varchar(32) NOT NULL,
		last_name varchar(128) NOT NULL,
		email varchar(128) NOT NULL,
		birth_date NOT NULL DEFAULT TIME,
		inserted_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT person_pk PRIMARY KEY (id)
	 );`)

}

// Reverse the migrations
func (m *Tables_20190309_015316) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE person")

}
