package decorator

import (
    "fmt"
    "testing"
)


func TestDecorator(t *testing.T) {
    component := &ConcreteComponent{}
    decoratorA := &ConcreteDecoratorA{component: component}
    decoratorB := &ConcreteDecoratorB{component: component}

    result := decoratorB.Operation()
    fmt.Println(result)

    result = decoratorA.Operation()
    fmt.Println(result)

    add := logExecution(add)  
    fmt.Println(add(1))  
}