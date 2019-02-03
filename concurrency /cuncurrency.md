## What is the difference between a process and a thread?
- Both processes and threads are units of concurrency. Processes do not share a common memory, while threads do.
- From the operating system’s point of view, a process is an independent piece of software that runs in its own virtual memory space.
- On the contrary, a thread is a part of an application that shares a common memory with other threads of the same application. Using common memory allows to shave off lots of overhead, design the threads to cooperate and exchange data between them much faster.

## **Thread pool**
- In computer programming, a thread pool is a software design pattern for achieving concurrency of execution in a computer program.
- A common method of scheduling tasks for thread execution is a synchronized queue, known as a task queue. The threads in the pool remove waiting tasks from the queue, and place them into a completed task queue after completion of execution.

## How can you create a Thread instance and run it?
- pass a Runnable instance to its constructor and call start(). Runnable is a functional interface, so it can be passed as a lambda expression.

## Different states of a Thread:
- New - a new Thread instance that was not yet started
- Runnable - a running thread.
- Blocked - a running thread becomes blocked if it needs to enter a synchronized section
- Waiting - waits for another thread to perform actions
- Terminated - a thread has completed the xecution of its Runnable.run() method

## Runnabale vs Callable interface
- The Runnable interface has a single run method. No value return and no throwing exceptions
- The Callable interface has a single call method and represents a task that has a value. That’s why the call method returns a value. It can also throw exceptions.

## What is the meaning of a synchronized keyword?
- The synchronized keyword before a block means that any thread entering this block has to acquire the monitor (the object in brackets).
```java
synchronized(obejct) {
	// ...
}
```
- A synchronized instance method has the same semantics, but the instance itself acts as a monitor.
```java
synchronized void instanceMethod() {
	// ...
}
```
- For a static synchronized method, the monitor is the *Class* object representing the declaring class.
```java
static synchronized void staticMethod() {
	// ...
}
```

## If two threads call a synchronized method on different object instances simultaneously, could one of these threads block? What if the method is static?
- If the method is an instance method, then the instance acts as a monitor for the method. Two threads calling the method on different instances acquire different monitors, so none of them gets blocked.

If the method is static, then the monitor is the Class object. For both threads, the monitor is the same, so one of them will probably block and wait for another to exit the synchronized method.