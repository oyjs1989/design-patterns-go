package factory

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	student := NewPerson("student")
	student.SayHello()

	teacher := NewPerson("teacher")
	teacher.SayHello()
}