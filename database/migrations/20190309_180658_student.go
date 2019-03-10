package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Student_20190309_180658 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Student_20190309_180658{}
	m.Created = "20190309_180658"

	migration.Register("Student_20190309_180658", m)
}

// Run the migrations
func (m *Student_20190309_180658) Up() {
	// -- --------------------------------------------------
	// --  Table Structure for `github.com/school/models.Student`
	// -- --------------------------------------------------
	m.SQL(`CREATE TABLE student (
        id int unsigned AUTO_INCREMENT NOT NULL,
		tutor_id int unsigned NOT NULL,
		person_id int unsigned NOT NULL,
		PRIMARY KEY (id),
		KEY student_preson_id_fkey (id),
		CONSTRAINT student_tutor_fk FOREIGN KEY (tutor_id) REFERENCES person (id),
		CONSTRAINT student_person_fk FOREIGN KEY (person_id) REFERENCES person (id)
		);`)

}

// Reverse the migrations
func (m *Student_20190309_180658) Down() {
	m.SQL("DROP TABLE student")
}
