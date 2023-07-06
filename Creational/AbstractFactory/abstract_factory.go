package abstractfactory

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


type LunchFactory interface {
	CreateTeacher() Person
	CreateStudent() Person
}


type SimpleLuchFactory struct {
}

func NewLunchFactory() LunchFactory {
	return new(SimpleLuchFactory)
}

func (s *SimpleLuchFactory) CreateTeacher() Person {
	return new(Teacher)
}

func (s *SimpleLuchFactory) CreateStudent() Person {
	return new(Student)
}