package factory

type Person interface {
	SayHello()
}

type Student struct {
}

func (s *Student) SayHello() {
	println("I am a student")
}

type Teacher struct {
}

func (t *Teacher) SayHello() {
	println("I am a teacher")
}

func NewPerson(name string) Person {
	switch name {
	case "student":
		return new(Student)
	case "teacher":
		return new(Teacher)
	default:
		return nil
	}
}

