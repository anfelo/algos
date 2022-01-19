# Security and HTTPS

## Man-In-The-Middle Attack

An attack in which the attacker intercepts a line of communication that is thought to be private by its two communicating parties.
If a malicious actor intercepted and mutated an IP packet on its way from a client to a server, that would be a man-in-the-middle attack.
MITM attacks are the primary thread that encryption and HTTPS aim to defend against.

## Symmetric Encryption

A type of encryption that relies on only a single key to both encrypt and decrypt data. The key must be known to all parties involved in communication and must therefore typically be shared between the parties at one point or another.
Symetric-key algorithms tend to be faster than their asymetric counterparts.
The most widely used symetric-key algorithms are part of the Advanced Encryption Standard (AES)

## Asymetric Encryption

Also known as public-key encryption, asymetric encryption relies on two keys - a public key and a private key - to encrypt and decrypt data. The keys are generated using cryptographic algorithms and are matematically connected such that data encrypted with the public key can only be decrypted with the private key.
While the private key must be kept secure to maintain the fidelity of this encryption paradigm, the public key can be openly shared.
Asymetric-key algorithms tend to be slower than their symmetric counterparts.

## AES

Stands for Advanced Encryption Standard. AES is widly used encryption standard that has thress symmetric-key algorithms (AES-128, AES-192, and AES-256)

Of note, AES is considered to be the "gold-standard" in encryption and is even used by the U.S. National Security Agency to encrypt top secret information.

## HTTPS

The HyperText Transfer Protocl Secure is an extension of HTTP that's used for secure communication online. It requires servers to have trusted certificates (usually SSL certificates) and uses the TRansport Layer Security (TLS), a security protocol built on top of TCP, to encrypt data communicated between a client and a server.

## TLS

The Transport Layer Security is a security protocol over which HTTPS runs in order to achieve secure communication online. HTTP over TLS is also know as HTTPS.

## SSL Certificate

A digital certificate granted to a server by a certificate authority. Contains the server's public key, to be used as part of the TLS handshake process in an HTTPS connection.

An SSL certificate effectively confirms that a public key belongs to the server claiming it belongs to them. SSL certificates are a crucial defense against man-in-the-middle attacks.
