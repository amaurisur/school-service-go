package viewmodels

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonBirthDate time.Time

type Person struct {
	DNI       uint          `json:"dni"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Birthdate JsonBirthDate `json:"birthdate"`
}

type Student struct {
	Person
	Tutor uint `json:"tutor_id"`
}

// Imeplement Marshaler und Unmarshalere interface
func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Maybe a Format function for printing a date
func (j JsonBirthDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
