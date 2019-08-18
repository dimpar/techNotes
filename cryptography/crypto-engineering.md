## Cryptocurrency Engineering and Design

### Signatures
3 function needed:
- GenerateKeys() - makes secret/public key.
Return private(secret) and public key. Returns only randomness

- Sign(secretKey, message) - returns a signature in bytes.

- Verify(publicKey, message, signature) - returns a boolean. Verify using a public key.

Signatures from hashes are called "Lamport Signatures". sha256 is a set of cryptographic hash functions.
Signing more than once with the same priv/pub key reveals more pieces of the private key.

Signature schemes:
> hash based (ex. Lamport)
- 1 time use (you can use a key once)
- kind of big

> RSA
- not used in Bitcoin (or any currency)
- basic setup:
-- make 2 primes: p, q
-- n = p*q (given p,q computing n is easy)
-- given n, finding p,q is hard.
-- it's a one way function, but not a hash function
-- more explanation on RSA encryption: https://www.khanacademy.org/computing/computer-science/cryptography#modern-crypt
-- RSA key sizes are smaller than hash based signatures; ofter 2048 bits (256 bytes)
-- lots of ways to lose a private key

> ECDSA (elliptic curves)
- Bitcoin's curve is: y^2 = x^3 + 7
- ECDSA explained in: https://youtu.be/0Q5IimX-AAc?list=PLUl4u3cNGP61KHzhg3JIJdK08JLSlcLId&t=2100
- EC private key a = 256 bit scalar
- EC public key A = a*G (G is a generator)
- Elliptic Curve Signature:
-- Have message m, privkey a
-- make 'k', a new random private key
-- R = k*G
-- s = k - hash(m, R)a
-- signature = R, s

> EC schnorr
- pattented version of ECDSA ? Not sure... 


### Potencial problems with digital token transfer

- Intercept transfer and steal funds
- Spend money without authorization
- Double spend
- Equivocate - meaning undoing a transfer

### Sybil attack
The Sybil attack in computer security is an attack where a reputation system is sabotaged by forging identities in peer-to-peer networks. Ex. fake accounts on FB / Twitter etc. In blockchain it means creating many nodes with a great computational power that can overtake good nodes and it can lead to 51% attack (taking over the network)

How to prevent a sybil attack?

### Proof of work
A Proof-of-Work (PoW) system (or protocol, or function) is a measure to deter denial of service attacks and other service abuses such as spam on a network by requiring some work from the service requester, usually meaning processing time by a computer.

