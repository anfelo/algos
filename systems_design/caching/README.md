## Cache

A piece of hardware or software that stores data, typically meant to retrieve that data faster than otherwise. Caches are often used to store responses to network requests of computationally-long operations.

Note that data in cache can become stale if the main source of truth for that data gets updated and the cache doesn't.

### Cache Hit

When requested data is found in a cache.

### Cache Miss

When requested data could have been found in a cache but isn't. This is typically used to refer to a negative consequence of a system failure or of a poor design choice. For example:

*If a server goes down, our load balancer will have to forward requests to a new server, which will result in cache misses*

### Cache Eviction Policy

The policy by which values get evicted or removed from a cache. Popular cache eviction policies include LRU (least-recently used), FIFO (first in first out), and LFU (least-frecuently used)