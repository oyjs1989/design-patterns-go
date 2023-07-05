package facade

type CarModel struct {
}

func (c *CarModel) SetModel() {
	println("CarModel - SetModel")
}

func NewCarMOdel() *CarModel {
	return new(CarModel)
}

type CarEngine struct {
}

func (c *CarEngine) SetEngine() {
	println("CarEngine - SetEngine")
}

func NewCarEngine() *CarEngine {
	return new(CarEngine)
}

type CarBody struct {
}

func (c *CarBody) SetBody() {

	println("CarBody - SetBody")
}

func NewCarBody() *CarBody {
	return new(CarBody)
}


type CarFacade struct {

	model *CarModel
	engine *CarEngine
	body *CarBody

}

func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetBody()
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model: NewCarMOdel(),
		engine: NewCarEngine(),
		body: NewCarBody(),
	}
}
