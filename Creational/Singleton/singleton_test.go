package singleton
import (
    "fmt"
    "testing"
)

func TestSingletion(t *testing.T) {
    singleton1 := GetInstance("singleton1")
    singleton2 := GetInstance("singleton2")

    fmt.Printf("Singleton 1: %s\n", singleton1.name)
    fmt.Printf("Singleton 2: %s\n", singleton2.name)
}