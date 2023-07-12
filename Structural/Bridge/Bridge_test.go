package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	implementorA := &ConcreteImplementorA{}
	implementorB := &ConcreteImplementorB{}

	abstractionA := &RefinedAbstraction{implementor: implementorA}
	resultA := abstractionA.Operation()
	fmt.Println(resultA)

	abstractionB := &RefinedAbstraction{implementor: implementorB}
	resultB := abstractionB.Operation()
	fmt.Println(resultB)
}
