package builder

// Product represents the object being built.
type Product struct {
    PartA string
    PartB string
    PartC string
}

// Builder is an interface for building the product.
type Builder interface {
    BuildPartA()
    BuildPartB()
    BuildPartC()
    GetProduct() *Product
}

// ConcreteBuilder is a concrete implementation of the Builder interface.
type ConcreteBuilder struct {
    product *Product
}

func (b *ConcreteBuilder) BuildPartA() {
    b.product.PartA = "Part A"
}

func (b *ConcreteBuilder) BuildPartB() {
    b.product.PartB = "Part B"
}

func (b *ConcreteBuilder) BuildPartC() {
    b.product.PartC = "Part C"
}

func (b *ConcreteBuilder) GetProduct() *Product {
    return b.product
}

// Director is responsible for managing the building process.
type Director struct {
    builder Builder
}

func (d *Director) Construct() {
    d.builder.BuildPartA()
    d.builder.BuildPartB()
    d.builder.BuildPartC()
}

func (d *Director) SetBuilder(builder Builder) {
    d.builder = builder
}
