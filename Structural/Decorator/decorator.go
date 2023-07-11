package decorator

import (
	"fmt"
)
// 装饰器模式是一种结构型设计模式，它允许您动态地将行为添加到对象中，而不会影响从同一类中派生的其他对象的行为。
// 它是通过将对象包装在一个装饰器类的对象中来实现的，该装饰器类具有与被包装对象相同的接口，并且可以在运行时添加或删除行为。

// 装饰器模式通常在以下情况下使用：
// 当需要动态地将行为添加到对象中，而不会影响从同一类中派生的其他对象的行为时。
// 当需要透明地、无需修改代码就能够扩展对象的功能时。
// 当需要在不使用子类化的情况下添加或删除对象的功能时。
// 以下是一些使用装饰器模式的场景示例：

// 添加日志记录：当需要在不修改现有代码的情况下添加日志记录功能时，可以使用装饰器模式。
// 例如，可以创建一个日志记录装饰器类，将其包装在现有对象中，从而在运行时添加日志记录功能。

// 添加缓存：当需要在不修改现有代码的情况下添加缓存功能时，可以使用装饰器模式。
// 例如，可以创建一个缓存装饰器类，将其包装在现有对象中，从而在运行时添加缓存功能。

// 添加安全检查：当需要在不修改现有代码的情况下添加安全检查功能时，可以使用装饰器模式。
// 例如，可以创建一个安全检查装饰器类，将其包装在现有对象中，从而在运行时添加安全检查功能。

// 总的来说，装饰器模式是一种非常有用的设计模式，它可以帮助我们更好地组织和管理代码，
// 并提高代码的可读性、可维护性、灵活性和可扩展性。
// 但是，开发人员应该根据具体情况选择合适的设计模式，以确保代码的可读性、可维护性、灵活性和可扩展性。

// Component is the interface that defines the basic operations of a component.
type Component interface {
    Operation() string
}

// ConcreteComponent is a concrete implementation of the Component interface.
type ConcreteComponent struct{}

// Operation implements the Operation method of the Component interface.
func (c *ConcreteComponent) Operation() string {
    return "ConcreteComponent"
}

// Decorator is the interface that defines the basic operations of a decorator.
type Decorator interface {
    Component
}

// ConcreteDecoratorA is a concrete implementation of the Decorator interface.
type ConcreteDecoratorA struct {
    component Component
}

// Operation implements the Operation method of the Component interface.
func (d *ConcreteDecoratorA) Operation() string {
    return "ConcreteDecoratorA(" + d.component.Operation() + ")"
}

// ConcreteDecoratorB is a concrete implementation of the Decorator interface.
type ConcreteDecoratorB struct {
    component Component
}

// Operation implements the Operation method of the Component interface.
func (d *ConcreteDecoratorB) Operation() string {
    return "ConcreteDecoratorB(" + d.component.Operation() + ")"
}

// 定义一个函数类型  
type myFunc func(int) int  
  
// 定义一个函数装饰器，用于打印函数的执行信息  
func logExecution(f myFunc) myFunc {  
    return func(i int) int {  
        fmt.Printf("Function %v is executed\n", f.Addr())  
        return f(i)  
    }  
}  
  
// 定义一个普通的函数  
func add(i int) int {  
    return i + 1  
}  