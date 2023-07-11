package singleton
import (
    "sync"
)

// Singleton is a singleton object.
type Singleton struct {
    name string
}

var instance *Singleton
var once sync.Once

// GetInstance returns the singleton instance.
func GetInstance(name string) *Singleton {
    once.Do(func() {
        instance = &Singleton{name: name}
    })
    return instance
}