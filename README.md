# 设计模式练习
本仓库是为了练习Go语言中的设计模式而创建的。

"Introduction to Design Patterns in 23" refers to the Gang of Four's (GoF) book "Design Patterns: Elements of Reusable Object-Oriented Software", which is a seminal work in the field of software engineering. The book presents 23 design patterns that are widely used in object-oriented programming.

The 23 design patterns are divided into three categories: Creational Patterns, Structural Patterns, and Behavioral Patterns. Creational Patterns deal with object creation mechanisms, trying to create objects in a manner suitable to the situation. Structural Patterns deal with object composition, trying to form large structures from individual objects. Behavioral Patterns deal with communication between objects, trying to define the manner in which objects interact and distribute responsibility.

The 23 design patterns are:

Creational Patterns:
1. Abstract Factory
2. Builder
3. Factory Method
4. Prototype
5. Singleton

Structural Patterns: 
6. Adapter 
7. Bridge 
8. Composite 
9. Decorator 
10. Facade 
11. Flyweight 
12. Proxy

Behavioral Patterns: 
13. Chain of Responsibility 
14. Command 
15. Interpreter 
16. Iterator 
17. Mediator 
18. Memento 
19. Observer 
20. State 
21. Strategy 
22. Template Method 
23. Visitor

Each design pattern has its own unique advantages and disadvantages, and is suited to different situations. By using these design patterns, developers can create more modular, flexible, and maintainable software systems.

Design Patterns是软件开发中常用的设计模式概念。有如下几类:
1. Creational Patterns(创建型模式):这些模式提供了创建对象的不同方式。
- Simple Factory:使用一个工厂类来创建对象。
- Factory Method:把类的实例化推迟到子类。
- Abstract Factory:提供一个接口来创建相关或相互依赖的对象族的实例。 
- Builder:将对象的构建与表示分离开。
- Prototype:用原型实例指定创建对象的种类。 
2. Structural Patterns(结构化型模式):这些模式主要关注类和对象的组成和结构。
- Adapter:将一个类的接口转换成客户希望的另外一个接口。
- Bridge:将抽象与实现分离,使它们可以独立地变化。
- Composite:将对象组合成树状结构以表示"部分-整体"的层次结构。 
- Decorator:动态地给一个对象添加额外的功能。
- Facade:为子系统提供一个一致的接口。
- Flyweight:共享已有的实例而不是创建新的实例。
3. Behavioral Patterns(行为型模式):关注对象之间通信的有效方式。
- Chain of Responsibility:使多个对象都有机会处理请求。 
- Command:将请求封装成一个对象,从而使你可用不同的请求对客户进行参数化。
- Iterator:提供一种方法顺序访问一个聚合对象中的各个元素。
- Mediator:用一个中介对象封装一组对象的交互。
- Memento:在不破坏封装性的前提下,捕获一个对象的内部状态,并在该对象之外保存这个状态。   
- Observer:定义对象之间的一种一对多的依赖关系。当一个对象改变状态,则其所有依赖者都会收到通知。
- State:允许对象在内部状态改变时改变它的行为。
- Strategy:定义一系列的算法,把它们一个个封装起来,并且使它们可相互替换。
- Template Method:定义一个操作中的算法的骨架,而将一些步骤延迟到子类中。  
- Visitor:表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。
总的来说,Design Patterns提供了一套开发经验和原则来解决软件设计上的问题。

设计模式是指在现实应用程序开发中重复出现的常见软件设计问题的可重用解决方案。
它们是通用的解决方案，可以根据各种情况进行调整。

设计模式并不特定于任何特定的编程语言或技术，而是适用于一般的软件开发。
它们旨在帮助开发人员创建更加模块化、灵活和可维护的软件。

有许多不同类型的设计模式，包括创作模式、结构模式和行为模式。
每种类型的模式都涉及软件设计的不同方面，例如对象创建、类组成或对象之间的通信。

通过使用设计模式，开发人员可以通过重用已验证的常见问题解决方案来节省时间和精力，而不是每次出现新问题时都重新发明轮子。
这可以带来更高效和有效的软件开发，以及更可靠和可维护的软件系统。

在Go语言中，最常见的设计模式包括：
单例模式（Singleton Pattern）：确保一个类只有一个实例，并提供全局访问点。
工厂模式（Factory Pattern）：定义一个用于创建对象的接口，让子类决定实例化哪个类。
建造者模式（Builder Pattern）：将一个复杂对象的构建过程与其表示分离，使得同样的构建过程可以创建不同的表示。
原型模式（Prototype Pattern）：通过克隆现有对象来创建新对象，而无需从头开始构建。
适配器模式（Adapter Pattern）：将一个类的接口转换成客户希望的另一个接口，使得原本由于接口不兼容而不能一起工作的类可以一起工作。
装饰器模式（Decorator Pattern）：动态地给一个对象添加一些额外的职责，而不需要修改其原始类。
桥接模式（Bridge Pattern）：将抽象部分与它的实现部分分离，使它们都可以独立地变化。
外观模式（Facade Pattern）：为子系统中的一组接口提供一个一致的界面，使得子系统更容易使用。
组合模式（Composite Pattern）：将对象组合成树形结构以表示“部分-整体”的层次结构，使得客户端对单个对象和组合对象的使用具有一致性。
享元模式（Flyweight Pattern）：运用共享技术来有效地支持大量细粒度对象的复用。
这些设计模式都是非常有用的，可以帮助我们更好地组织和管理代码，并提高代码的可读性、可维护性、灵活性和可扩展性。但是，开发人员应该根据具体情况选择合适的设计模式，以确保代码的可读性、可维护性、灵活性和可扩展性。

设计模式二十三，听我来一一念：

单例模式保证唯一，工厂方法创对象。

抽象工厂一起用，建造者造复杂。

原型模式复制它，适配器改接口。

桥接模式连接它，装饰者来修饰。

外观模式来简化，享元模式缓存它。

组合模式树形结构，代理模式控制它。

模板方法定义步骤，策略模式随便换。

命令模式发出命令，职责链传递它。

状态模式状态转移，观察者来观察。

备忘录模式备份它，访问者增加功能。

解释器模式解释它，中介者调和它。

这就是23种，你都学会了吗？

[toc]

## 简介
本仓库包含了Go语言中常见的设计模式的示例代码。每个示例都包含了一个简单的说明和示例代码。

## 安装
克隆本仓库
安装Go语言环境
## 使用
进入到相应的示例目录
运行示例代码
## 贡献
欢迎贡献代码！请遵循以下步骤：

```
Fork 本仓库
创建一个新的分支 (git checkout -b feature/your-feature)
提交你的修改 (git commit -am 'Add some feature')
推送到分支 (git push origin feature/your-feature)
创建一个新的 Pull Request
```
## 许可证
本仓库使用 MIT 许可证。