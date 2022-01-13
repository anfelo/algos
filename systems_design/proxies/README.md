## Proxies

A server that sits in the middle of a client and a server manipulating the requests and responses in some manner.

### Forward Proxy

A server that sits between a client and servers and acts on behalf of the client, typically used to mask the client's identity (IP address). Note that forward proxies are often referred to as just proxies.

In this case the source IP of where the request was originated is hidden from the final destination of the request by the forward proxy.

This is how VPNs work under the hood.

### Reverse Proxy

A server that sits between clients and servers and acts on behalf of the servers, typically used for logging, load balancing, or caching.

The client thinks it is interacting with the server but it is actually interacting with the reverse proxy.

### Nginx

Is a very popular webserver that's often used as a reverse proxy and load balancer.