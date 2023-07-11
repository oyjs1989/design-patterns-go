package builder

// Builder Pattern（建造者模式）通常在以下情况下使用：
// 创建复杂对象：当需要创建复杂对象时，特别是当对象具有多个可选组件时，可以使用Builder Pattern。
// 例如，创建一个汽车对象，它包含多个可选组件，如发动机、座椅、轮胎等。

// 将对象的构建过程与其表示分离：当需要将对象的构建过程与其表示分离时，可以使用Builder Pattern。
// 例如，创建一个HTML文档对象，它包含多个元素，如标题、段落、列表等。
// 使用Builder Pattern可以将HTML文档对象的构建过程与其表示分离，从而提高代码的可读性和可维护性。

// 更好地控制对象的构建过程：当需要更好地控制对象的构建过程时，可以使用Builder Pattern。
// 例如，创建一个订单对象，它包含多个订单项，每个订单项包含商品、数量、价格等信息。
// 使用Builder Pattern可以更好地控制订单对象的构建过程，从而提高代码的灵活性和可扩展性。

// 创建不同表示的对象：当需要创建不同表示的对象时，而不需要更改其构建过程时，可以使用Builder Pattern。
// 例如，创建一个图形对象，它可以表示不同类型的图形，如圆形、矩形、三角形等。
// 使用Builder Pattern可以更容易地创建不同表示的图形对象，而不需要更改其构建过程。

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
