# Hashing

A function that takes in a specific data type (such as a string or an identifier) and outputs a number. Different inputs may have the same output, but a good hashing function attempts to minimize those hashing collisions (which is equivalent to maximizing uniformity)

## Consistent Hashing

A type of hashing that minimizes the number of keys that need to be remapped when a hash table gets resized. It's often used by load balancers to distribute traffic to servers; it minimizes the number of requests that get forwarded to different servers when new servers are added or when existing servers are brought down.

## Rendezvous Hashing

A type of hashing also coined **highest random weight** hashing. Allows for minimal redistribution of mappings when a server goes down.

## SHA

Short for "Secure Hash Algorithms", the SHA is a collection of cryptographic hash function used in the industry. These days, SHA-3 is a popular choice to use in a system.