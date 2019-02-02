**DNS** - A Domain Name System (DNS) translates a domain name such as www.example.com to an IP address.

**CDN** - A content delivery network (CDN) is a globally distributed network of proxy servers, serving content from locations closer to the user.

**Load balancers** - Load balancers distribute incoming client requests to computing resources such as application servers and databases.
- Preventing requests from going to unhealthy servers
- Preventing overloading resources
- Helping eliminate single points of failure

**Sharding** distributes data across different databases such that each database can only manage a subset of the data. Taking a users database as an example, as the number of users increases, more shards are added to the cluster.

**Denormalization** attempts to improve read performance at the expense of some write performance. Redundant copies of the data are written in multiple tables to avoid expensive joins.

**REST** is an architectural style enforcing a client/server model where the client acts on a set of resources managed by the server. The server provides a representation of resources and actions that can either manipulate or get a new representation of resources. All communication must be stateless and cacheable.

There are four qualities of a RESTful interface:
- Identify resources (URI in HTTP) - use the same URI regardless of any operation.
- Change with representations (Verbs in HTTP) - use verbs, headers, and body.
- Self-descriptive error message (status response in HTTP) - Use status codes, don't reinvent the wheel.

**Authentication using stateless approach**:
A browser or mobile client makes a request to the authentication server containing user login information. The authentication server generates a new JWT (Json Web Token) access token and returns it to the client. And client needs to store that token and send on every request to a restricted resource in the Authorization header. Server then validates the token and if it’s valid, returns the secure resource to the client.

**Reverse proxy**:
- caching web acceleration
- security against external traffic
- Isolation
- can act as a load balancer