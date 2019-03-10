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

type PersonMediator struct {
	Db     database.DbInterface
	logger loggo.Logger
	MediatorError
}

func NewPersonMediator(DbInterface database.DbInterface) *PersonMediator {
	loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))
	logger := loggo.GetLogger("mediators.Person")
	return &PersonMediator{
		Db:     DbInterface,
		logger: logger,
	}
}

func (sm *PersonMediator) CreatePerson(p viewmodels.Person) *MediatorError {
	sm.logger.Infof("starting process to create person")

	prs := models.Person{
		DNI:       p.DNI,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Birthdate: (time.Time)(p.Birthdate),
	}

	idPrs, err := sm.Db.Insert(&prs)
	if err != nil {
		sm.Db.Rollback()
		return sm.MediatorError.GetError("err_internal_server_error", err.Error())
	}
	fmt.Printf("\nperson created: %d\n", idPrs)

	sm.logger.Infof("end create person")
	return nil
}
