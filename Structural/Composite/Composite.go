package composite

// Composite Pattern是一种结构型设计模式，它可以将对象组合成树形结构，
// 从而使得客户端可以统一地处理单个对象和组合对象。

// 在Composite Pattern中，我们定义了一个Component接口，它定义了所有
// 组件的通用方法。我们还定义了两个具体的组件：Leaf和Composite。Leaf
// 表示叶子节点，它没有子节点。Composite表示复合节点，它包含多个子节点。

// Composite Pattern通常在以下情况下使用：

// 当需要将对象组合成树形结构时。
// 当需要客户端可以统一地处理单个对象和组合对象时。
// 当需要在运行时动态地添加或删除组件时。
// 以下是两个使用Composite Pattern的示例：

// 组织结构
// 在一个组织中，部门可以包含多个子部门和多个员工。我们可以使用Composite
//  Pattern来将部门和员工组合成树形结构。部门可以作为Composite节点，
// 员工可以作为Leaf节点。这样，我们就可以在运行时动态地添加或删除部门和员工，而不会影响整个组织结构。

// 图形界面
// 在图形界面中，窗口可以包含多个子控件，例如按钮、文本框、标签等。
// 我们可以使用Composite Pattern来将窗口和控件组合成树形结构。
// 窗口可以作为Composite节点，控件可以作为Leaf节点。
// 这样，我们就可以在运行时动态地添加或删除控件，而不会影响整个界面结构。

import "fmt"

// Component is the interface that defines the common methods for all components.
type Component interface {
    Add(Component)
    Remove(Component)
    Display(int)
}

// Leaf is a concrete implementation of the Component interface.
type Leaf struct {
    name string
}

// Add does nothing for a Leaf.
func (l *Leaf) Add(c Component) {}

// Remove does nothing for a Leaf.
func (l *Leaf) Remove(c Component) {}

// Display displays the name of the Leaf.
func (l *Leaf) Display(indent int) {
    fmt.Printf("%s%s\n", getIndent(indent), l.name)
}

// Composite is another concrete implementation of the Component interface.
type Composite struct {
    name       string
    components []Component
}

// Add adds a component to the Composite.
func (c *Composite) Add(component Component) {
    c.components = append(c.components, component)
}

// Remove removes a component from the Composite.
func (c *Composite) Remove(component Component) {
    for i, comp := range c.components {
        if comp == component {
            c.components = append(c.components[:i], c.components[i+1:]...)
            break
        }
    }
}

// Display displays the name of the Composite and all of its components.
func (c *Composite) Display(indent int) {
    fmt.Printf("%s%s\n", getIndent(indent), c.name)
    for _, comp := range c.components {
        comp.Display(indent + 2)
    }
}

// getIndent returns a string of spaces for indentation.
func getIndent(indent int) string {
    result := ""
    for i := 0; i < indent; i++ {
        result += " "
    }
    return result
}
