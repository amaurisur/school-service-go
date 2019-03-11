package mediators

import (
	"fmt"
	"os"
	"time"

	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
	"github.com/school/database"
	"github.com/school/models"
	"github.com/school/viewmodels"
)

type StudentMediator struct {
	Db     database.DbInterface
	logger loggo.Logger
	MediatorError
}

func NewStudentMediator(DbInterface database.DbInterface) *StudentMediator {
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))
	logger := loggo.GetLogger("mediators.Student")
	return &StudentMediator{
		Db:     DbInterface,
		logger: logger,
	}
}

func (sm *StudentMediator) CreateStudent(s viewmodels.Student) *MediatorError {
	sm.logger.Infof("starting process to create student")
	sm.Db.Begin()

	p := models.Person{
		DNI:       s.DNI,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Birthdate: (time.Time)(s.Birthdate),
	}

	idPrs, err := sm.Db.Insert(&p)
	if err != nil {
		sm.Db.Rollback()
		return sm.MediatorError.GetError("err_internal_server_error", err.Error())
	}
	fmt.Printf("\nstudent created: %d\n", idPrs)

	std := models.Student{
		Person: &p,
		Tutor: &models.Person{
			ID: s.Tutor,
		},
	}
	idStd, err := sm.Db.Insert(&std)
	if err != nil {
		sm.Db.Rollback()
		return sm.MediatorError.GetError("err_internal_server_error", err.Error())
	}
	fmt.Printf("\nstudent created: %d\n", idStd)

	sm.Db.Commit()
	sm.logger.Infof("end create student")
	return nil
}

func (sm *StudentMediator) GetAllStudentByLastName() ([]models.Student, error) {
	sm.logger.Infof("starting get all student")

	// TODO: clean this
	// [ORM]2019/03/10 23:16:07  -[Queries/default] -
	// [  OK /    db.Query /     0.5ms] - [SELECT T0.`id`, T0.`person_id`, T0.`tutor_id`, T1.`id`, T1.`dni`,
	// T1.`first_name`, T1.`last_name`, T1.`birthdate`, T1.`inserted_at` FROM `student` T0
	// INNER JOIN `person` T1 ON T1.`id` = T0.`person_id`
	// ORDER BY T1.`last_name` ASC LIMIT 1000]
	var students []models.Student
	qs := sm.Db.GetQueryTable("student")
	_, err := qs.RelatedSel("Person").OrderBy("Person__LastName").All(&students)

	sm.logger.Infof("end get all student")
	return students, err
}
