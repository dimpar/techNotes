The Factory Method pattern is a design pattern used to define a runtime interface for creating an object. It’s called a factory because it creates various types of objects without necessarily knowing what kind of object it creates or how to create it.

![alt text](../../../images/factory-method.gif)

- *Product* defines the interface of objects the factory method creates
- *ConcreteProduct* implements the Product interface
- *Creator* declares the factory method, which returns an object of type Product
- *ConcreteCreator* overrides the factory method to return an instance of a Concrete Product

Purpose
- Allows the sub-classes to choose the type of objects to create at runtime
- It provides a simple way of extending the family of objects with minor changes in application code.
- Promotes the loose-coupling by eliminating the need to bind application-specific structs into the code

