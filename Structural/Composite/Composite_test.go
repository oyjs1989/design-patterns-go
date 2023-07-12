package composite

import "testing"

func TestComposite(t *testing.T) {
    root := &Composite{name: "root"}
    root.Add(&Leaf{name: "leaf1"})
    root.Add(&Leaf{name: "leaf2"})

    comp := &Composite{name: "composite1"}
    comp.Add(&Leaf{name: "leaf3"})
    comp.Add(&Leaf{name: "leaf4"})

    root.Add(comp)

    root.Display(0)
}