package adapter

// 适配器模式（Adapter Pattern）是一种结构型设计模式，它允许将现有的类或接口转换为另一种接口，以满足客户端的需求。
// 适配器模式通过创建一个适配器类来实现，该适配器类包装了现有的类或接口，并将其转换为另一种接口。

// 适配器模式通常在以下情况下使用：

// 当需要将现有的类或接口转换为另一种接口时。
// 当需要与现有的类或接口进行协作，但是它们的接口不兼容时。
// 当需要在不修改现有代码的情况下添加新的功能时。
// 以下是一些使用适配器模式的场景示例：

// 将现有的数据库接口适配为ORM接口：当需要将现有的数据库接口适配为ORM接口时，可以使用适配器模式。
// 例如，可以创建一个适配器类，将现有的数据库接口包装在其中，并将其转换为ORM接口。

// 将现有的日志接口适配为新的日志接口：当需要将现有的日志接口适配为新的日志接口时，可以使用适配器模式。
// 例如，可以创建一个适配器类，将现有的日志接口包装在其中，并将其转换为新的日志接口。

// 将现有的支付接口适配为新的支付接口：当需要将现有的支付接口适配为新的支付接口时，可以使用适配器模式。
// 例如，可以创建一个适配器类，将现有的支付接口包装在其中，并将其转换为新的支付接口。

// 总的来说，适配器模式是一种非常有用的设计模式，它可以帮助我们更好地组织和管理代码，并提高代码的可读性、可维护性、灵活性和可扩展性。
// 但是，开发人员应该根据具体情况选择合适的设计模式，以确保代码的可读性、可维护性、灵活性和可扩展性。

// Target is the interface that defines the desired operations.
type Target interface {
    Request() string
}

// Adaptee is the interface that defines the existing operations.
type Adaptee interface {
    SpecificRequest() string
}

// ConcreteAdaptee is a concrete implementation of the Adaptee interface.
type ConcreteAdaptee struct{}

// SpecificRequest implements the SpecificRequest method of the Adaptee interface.
func (a *ConcreteAdaptee) SpecificRequest() string {
    return "Adaptee method"
}

// Adapter is the adapter that adapts the Adaptee interface to the Target interface.
type Adapter struct {
    adaptee Adaptee
}

// Request implements the Request method of the Target interface.
func (a *Adapter) Request() string {
    return "Adapter method: " + a.adaptee.SpecificRequest()
}
