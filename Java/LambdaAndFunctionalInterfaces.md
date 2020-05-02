## Default methods in functional interfaces

### What is the need of functional interface?
They can be used to pass a block of code to another method or object. Functional Interface serves as a data type for Lambda expressions. Since a Functional interface contains only one abstract method, the implementation of that method becomes the code that gets passed as an argument to another method.

You can add default methods to the functional interface. 

```java
@FunctionalInterface
public interface Foo {
    String method();
    default void defaultMethod() {}
}
```

Functional interfaces can be extended by other functional interfaces if their abstract methods have the same signature. For example:

```java
	@FunctionalInterface
	public interface FooExtended extends Baz, Bar {}
	     
	@FunctionalInterface
	public interface Baz {  
	    String method();    
	    default void defaultBaz() {}        
	}
	     
	@FunctionalInterface
	public interface Bar {  
	    String method();    
	    default void defaultBar() {}    
	}
```

Adding too many default methods to the interface is not a very good architectural decision.

## Instantiate Functional Interfaces 
Foo foo = parameter -> parameter + " from Foo";

## Best practices for lambda

### Avoid Blocks of Code in Lambda’s Body

```java
	//good
	Foo foo = parameter -> buildString(parameter);

	private String buildString(String parameter) {
	    String result = "Something " + parameter;
	    //many lines of code
	    return result;
	}

	//bad
	Foo foo = parameter -> { String result = "Something " + parameter; 
	    //many lines of code 
	    return result; 
	};
```

### Avoid specifying parameter types
```java
	//good
	(a, b) -> a.toLowerCase() + b.toLowerCase();

	//bad
	(String a, String b) -> a.toLowerCase() + b.toLowerCase();
```

### Avoid Parentheses Around a Single Parameter
```java
	//good
	a -> a.toLowerCase();

	//bad
	(a) -> a.toLowerCase();
```

### Avoid return statement and braces
```java
	//good
	a -> a.toLowerCase();

	//bad
	a -> {return a.toLowerCase()};
```

### Use Method References
```java
	//good
	a -> a.toLowerCase();

	//but better this way. This is calling a method which is already implemented elsewhere
	String::toLowerCase;
```

## Use “Effectively Final” Variables

```java
	public void method() {
	    String localVariable = "Local";
	    Foo foo = parameter -> {
	        String localVariable = parameter;
	        return localVariable;
	    };
	}
```

The compiler will inform you that `Variable 'localVariable' is already defined in the scope.` Accessing a non-final variable inside lambda expressions will cause the compile-time error

