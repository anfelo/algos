# Replication and Sharding

## Replication

The act of duplicating the data from one database server to others. This is sometimes used to increase the redundancy of your system and tolerate regional failures for instance. Other times you can use replication to move data closer to your clients, thus decreasing the latency of accessing specific data.

## Sharding

Sometimes called **data partitioning**, sharding is the act of splitting a database into two or more pieces called **shards** and is typically done to increase the throughtput of your database. Popular sharding strategies include:

* Sharding based on client's region
* Sharding based on the type of data beign stored (e.g. user data gets stored in one shard, payments data gets in another shard)
* Sharding based on the hash of a column (only for structured data)