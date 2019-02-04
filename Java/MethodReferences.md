### When to use

Instead of using
*AN ANONYMOUS CLASS*
you can use
*A LAMBDA EXPRESSION*
And if this just calls one method, you can use
*A METHOD REFERENCE*

## Static method reference
```java
	class Numbers {
	  public static boolean isMoreThanFifty(int n1, int n2) {
	    return (n1 + n2) > 50;
	  }
	  public static List<Integer> findNumbers(
	    List<Integer> l, BiPredicate<Integer, Integer> p) {
	      List<Integer> newList = new ArrayList<>();
	      for(Integer i : l) {
	        if(p.test(i, i + 10)) {
	          newList.add(i);
	        }
	      }
	      return newList;
	  }
	}


	List<Integer> list = Arrays.asList(12,5,45,18,33,24,40);

	// Using an anonymous class
	findNumbers(list, new BiPredicate<Integer, Integer>() {
	  public boolean test(Integer i1, Integer i2) {
	    return Numbers.isMoreThanFifty(i1, i2);
	  }
	});

	// Using a lambda expression
	findNumbers(list, (i1, i2) -> Numbers.isMoreThanFifty(i1, i2));

	// Using a method reference
	findNumbers(list, Numbers::isMoreThanFifty);
```

## Instance method reference of an object of a particular type
```java
	(obj, args) -> obj.instanceMethod(args)
```
can be turned into
```java
	ObjectType::instanceMethod
```

## Instance method reference of an existing object
```java
	(args) -> obj.instanceMethod(args)
```
can be turned into
```java
	obj::instanceMethod
```


## Constructor method reference
```java
	(args) -> new ClassName(args)
```
can be turned into
```java
	ClassName::new
```
