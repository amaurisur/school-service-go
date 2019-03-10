package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type SubjectStudents_20190309_200929 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &SubjectStudents_20190309_200929{}
	m.Created = "20190309_200929"

	migration.Register("SubjectStudents_20190309_200929", m)
}

// Run the migrations
func (m *SubjectStudents_20190309_200929) Up() {
	// -- --------------------------------------------------
	// --  Table Structure for `github.com/school/models.SubjectStudents`
	// -- --------------------------------------------------
	m.SQL(`CREATE TABLE subject_students (
        subject_id int unsigned NOT NULL,
		student_id int unsigned NOT NULL,
		PRIMARY KEY (subject_id, student_id),
		KEY subject_id_student_id (subject_id,student_id),
		CONSTRAINT subject_id_student_id_fk FOREIGN KEY (subject_id) REFERENCES subject (id),
		CONSTRAINT student_id_subject_id_fk FOREIGN KEY (student_id) REFERENCES student (id)
    );`)

}

// Reverse the migrations
func (m *SubjectStudents_20190309_200929) Down() {
	m.SQL("DROP TABLE subject_students")
}
