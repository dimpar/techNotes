# Performance vs scalability

A service is **scalable** if it results in increased performance in a manner proportional to resources added. Generally, increasing performance means serving more units of work, but it can also be to handle larger units of work, such as when datasets grow.1

Another way to look at performance vs scalability:

- If you have a performance problem, your system is slow for a single user.

- If you have a scalability problem, your system is fast for a single user but slow under heavy load.

# Latency vs throughput

**Latency** is the time to perform some action or to produce some result.

**Throughput** is the number of such actions or results per unit of time.

Generally, you should aim for maximal throughput with acceptable latency.

# Replication

**Master-slave replication** - The master serves reads and writes, replicating writes to one or more slaves, which serve only reads.

**Master-master replication** - Both masters serve reads and writes and coordinate with each other on writes.