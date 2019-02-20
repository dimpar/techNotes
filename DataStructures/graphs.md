## Graphs implementation:

- Adjacency Matrix

![alt text](../images/adjacencymatrix.png)

*Pros*: Representation is easier to implement and follow. Removing an edge takes O(1) time. Queries like whether there is an edge from vertex ‘u’ to vertex ‘v’ are efficient and can be done O(1).

*Cons*: Consumes more space O(V^2). Even if the graph is sparse(contains less number of edges), it consumes the same space. Adding a vertex is O(V^2) time.

- Adjacency List
![alt text](../images/listadjacency.png)