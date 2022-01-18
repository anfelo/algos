# Publish/Subscribe Pattern

Often shortened as **Pub/Sub**, the Publish/Subscribe pattern is a popular messaging model that consists of **publishers** and **subscribers**. Publishers publish messages to special topics (sometimes called **channels**) without caring about or even knowing who will read those messages, and subscribers subscribe to topics an read messages coming through those topics.

Pub/Sub systems often come with very powerful guarantees like **at-least-once delivery**, **persistent storage**, **ordering of messages**, and replayability of messages.

## Idempotent Operation

An operation that has the same ultimate outcome regardless of how many times it's performed. IIf an operation can be performed multiple times without changing its overall effect, it's idempotent. Operations performed through a **Pub/Sub** messaging system typically have to be idempotent, since Pub/Sub systems tend to allow the same messages to be consumed multiple times.

For example, increasing an integer value in a database is not an idempotent operation, since repeating this operation will not have the same effect as if it had been performed only once. Conversely, setting a value to "COMPLETE" is an idempotent operation, since repeating this operation will always yield the same result: the value will be "COMPLETE".

Examples: Apache Kafka, Cloud Pub/Sub, Nats, RabbitMQ