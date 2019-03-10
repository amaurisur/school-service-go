package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/school/database"
	"github.com/school/mediators"
	"github.com/school/viewmodels"
)

// SchoolController _
type SchoolController struct {
	Controller
}

// PostStudent allow post student to School
// @Title Post Student
// @Description Post student to School
// @Param body body viewmodels.Student true "Struct that contains student information"
// @Success 200 {object} viewmodels.Message
// @Failure 400 Bad Request
// @router /student [post]
func (sc *SchoolController) PostStudent() {
	fmt.Printf("body: %s", sc.Ctx.Request.Body)
	if sc.Ctx == nil || sc.Ctx.Request.Body == nil {
		sc.ServeResponse(sc.ComposeResponseError(http.StatusBadRequest, "missing_param", "Missing param Body"))
		return
	}

	decoder := json.NewDecoder(sc.Ctx.Request.Body)
	decoder.DisallowUnknownFields()

	var s viewmodels.Student
	if err := decoder.Decode(&s); err != nil {
		sc.ServeResponse(sc.ComposeResponseError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()))
		return
	}

	m := mediators.NewStudentMediator(database.NewDbWrapper())
	err := m.CreateStudent(s)
	if err != nil {
		sc.ServeResponse(sc.ComposeResponseError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()))
		return
	}

	sc.ServeResponse(sc.ComposeResponse(http.StatusOK, map[string]string{"message": "student registred"}))
}

// PostPerson allow post person to School
// @Title Post Person
// @Description Post person to School
// @Param body body viewmodels.Person true "Struct that contains person information"
// @Success 200 {object} viewmodels.Message
// @Failure 400 Bad Request
// @router /person [post]
func (sc *SchoolController) PostPerson() {
	fmt.Printf("body: %s", sc.Ctx.Request.Body)
	if sc.Ctx == nil || sc.Ctx.Request.Body == nil {
		sc.ServeResponse(sc.ComposeResponseError(http.StatusBadRequest, "missing_param", "Missing param Body"))
		return
	}

	decoder := json.NewDecoder(sc.Ctx.Request.Body)
	decoder.DisallowUnknownFields()

	var p viewmodels.Person
	if err := decoder.Decode(&p); err != nil {
		sc.ServeResponse(sc.ComposeResponseError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()))
		return
	}

	m := mediators.NewPersonMediator(database.NewDbWrapper())
	err := m.CreatePerson(p)
	if err != nil {
		sc.ServeResponse(sc.ComposeResponseError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()))
		return
	}

	sc.ServeResponse(sc.ComposeResponse(http.StatusOK, map[string]string{"message": "person registred"}))
}
