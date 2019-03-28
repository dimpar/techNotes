## Chapter 5

### Decoupling and the Law of Demeter
- minimize coupling -  careful about how many other modules you
interact with and, more importantly, how you came to interact with
them.
- think of a general contractor. Ex. You deal with a main builder contractor and
he deals with other subcontractors (plumbing, roofing etc.)
- In order to keep the dependencies to a minimum, we’ll use the Law of Demeter to design
our methods and functions.
- In practice, this means that you will be writing
a large number of wrapper methods that simply forward the request on
to a delegate. These wrapper methods will impose both a runtime cost
and a space overhead, which may be significant—even prohibitive—in
some applications.

The Law of Demeter for functions states that any method of an object should call only methods
belonging to:
![alt text](../images/lawOfDemeterForFunctions.png)

### Metaprogramming

- Put Abstractions in Code, Details in Metadata
- not a lot of modern languages support metaprogramming
- metaprogramming is close to writing actual rules. You're writing the same thing that you're saying.
- trying to be close to user's language