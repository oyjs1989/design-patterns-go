package bridge

// 桥模式（Bridge Pattern）是一种结构型设计模式，它可以将一个大类或一系列
// 紧密相关的类拆分为抽象和实现两个独立的层次结构，从而使它们可以独立地变化。

// 在桥模式中，抽象部分包含一个指向实现部分的引用，而实现部分则包含实现抽象
// 部分所需的具体实现。这种分离使得抽象部分和实现部分可以独立地扩展和变化，而不会相互影响。

// 桥模式通常在以下情况下使用：

// 当需要将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构时。
// 当需要在运行时动态地将抽象部分和实现部分组合在一起时。
// 当需要避免在抽象部分中使用继承时。
// 以下是两个使用桥模式的示例：

// 操作系统和文件系统
// 在操作系统中，文件系统是一个重要的组成部分。文件系统可以有多种实现方式，
// 例如FAT32、NTFS、EXT4等。操作系统需要使用文件系统来管理文件和目录，但是
// 它不应该依赖于任何特定的文件系统实现。这时，我们可以使用桥模式来将操作系
// 统和文件系统分离开来。操作系统作为抽象部分，文件系统作为实现部分，它们之
// 间通过桥接器进行连接。这样，我们就可以在运行时动态地将不同的文件系统实现
// 与操作系统组合在一起，而不会影响操作系统的其他部分。

// 形状和颜色
// 在图形设计中，形状和颜色是两个重要的属性。我们可以定义一个Shape抽象类来
// 表示形状，它包含一个指向Color实现类的引用。Color实现类包含实现Shape所需
// 的具体颜色。这样，我们就可以在运行时动态地将不同的颜色实现与形状组合在一
// 起，而不会影响形状的其他部分。例如，我们可以定义一个Circle类来表示圆形，
// 它包含一个Color实现类的引用。我们可以创建不同的Color实现类来表示不同的颜色，
// 例如Red、Green、Blue等。这样，我们就可以在运行时动态地将不同的颜色实现与
// Circle组合在一起，从而创建不同颜色的圆形。

// Implementor is the interface that defines the implementation methods.
type Implementor interface {
    OperationImpl() string
}

// ConcreteImplementorA is a concrete implementation of the Implementor interface.
type ConcreteImplementorA struct{}

// OperationImpl implements the OperationImpl method of the Implementor interface.
func (a *ConcreteImplementorA) OperationImpl() string {
    return "ConcreteImplementorA"
}

// ConcreteImplementorB is another concrete implementation of the Implementor interface.
type ConcreteImplementorB struct{}

// OperationImpl implements the OperationImpl method of the Implementor interface.
func (b *ConcreteImplementorB) OperationImpl() string {
    return "ConcreteImplementorB"
}

// Abstraction is the interface that defines the abstraction methods.
type Abstraction interface {
    Operation() string
}

// RefinedAbstraction is a concrete implementation of the Abstraction interface.
type RefinedAbstraction struct {
    implementor Implementor
}

// Operation implements the Operation method of the Abstraction interface.
func (a *RefinedAbstraction) Operation() string {
    return "RefinedAbstraction: " + a.implementor.OperationImpl()
}
