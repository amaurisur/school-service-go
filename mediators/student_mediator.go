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
