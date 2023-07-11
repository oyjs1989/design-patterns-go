package prototype

import (
	"testing"
    "fmt"
)

func TestConcreate(t *testing.T) {
    prototype := &ConcretePrototype{name: "prototype"}
    clone := prototype.Clone().(*ConcretePrototype)

    fmt.Printf("Original: %s\n", prototype.name)
    fmt.Printf("Clone: %s\n", clone.name)
}