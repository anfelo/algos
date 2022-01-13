# Key Terms

## IP

Stands for **Internet Protocol**. This network protocol outlines how almost all
machine-to-machine communications should happen in the world. Other protocols
like **TCP**, **UDP** and **HTTP** are built on top of IP.

### IP Packet

Sometimes more broadly referred to as just a (network) packet, an IP packet is
effectively the smallest unit used to describe data being sent over IP, asude
from butes. An IP packet consists of:

- an IP header, which contains the source and destination IP addresses as well
  as other information related to the network
- a payload, which is just the data being sent over the network

## TCP

Network protocol built on top of the Internet Protocol (IP). Allows for ordered,
reliable data delivery between machines over the public internet by creating a
**connection (handshake)**.

TCP is usually implemented in the kernel, which exposes **sockets** to
applications that they can use to stream data through an open connection.

## HTTP

The HyperText Transfer Protocol is a very common network implmented on top of
TCP. Clients make HTTP requests, and severs respond with a response.
