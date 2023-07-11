package adapter

import (
    "testing"
    "fmt"
)

func TestAdapter(t *testing.T) {
    adaptee := &ConcreteAdaptee{}
    adapter := &Adapter{adaptee: adaptee}
    result := adapter.Request()
    fmt.Println(result)
}