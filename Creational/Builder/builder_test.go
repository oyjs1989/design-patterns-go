package builder
import (
	"testing"
	"fmt"
)


func TestBuilder(t *testing.T){
    builder := &ConcreteBuilder{product: &Product{}}
    director := &Director{builder: builder}

    director.Construct()
    product := builder.GetProduct()

    fmt.Printf("Product Parts: %s, %s, %s\n", product.PartA, product.PartB, product.PartC)
}