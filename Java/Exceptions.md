## Checked vs uncheckded

- Checked: are the exceptions that are checked at **compile** time. If some code within a method throws a checked exception, then the method must either **handle the exception (try / catch)** or it must specify the exception using **throws** keyword.
To create a custom checked exception, we have to extend the java.lang.Exception class.

```java
public static void main(String[] args) 
{
    FileReader file = new FileReader("somefile.txt"); // assuming there's no somefile.txt
}
```

- Unchecked: are the exceptions that are **not checked at compiled** time.
In Java exceptions under Error and RuntimeException classes are unchecked exceptions, everything else under throwable is checked.
To create a custom unchecked exception we need to extend the java.lang.RuntimeException class

```java
// JVM not forcing you to check NullPointerExceptio
public static void main(String[] args) 
{
    try
    {
        FileReader file = new FileReader("pom.xml"); // assuming there's a pom.xml file available
         
        file = null;
         
        file.read();
    } 
    catch (IOException e) 
    {
        //Alternate logic
        e.printStackTrace();
    }
}
```

- checked or unchecked? *If a client can reasonably be expected to recover from an exception, make it a checked exception. If a client cannot do anything to recover from the exception, make it an unchecked exception*

- Custom exceptions are very much useful when we need to handle specific exceptions related to the business logic.

![alt text](../images/exceptions.jpg)