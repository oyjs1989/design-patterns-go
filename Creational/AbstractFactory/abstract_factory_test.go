package abstractfactory

import (
	"testing"
)

func TestLunchFaotry(t *testing.T) {
	NewLunchFactory().CreateStudent().SayHello()
	NewLunchFactory().CreateTeacher().SayHello()
}