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

### Blockchain
you can't change anything in a past without having it reflected in the future. Each block has a hash of a previous block and if you change smth in that block, hash will change too. Hashes are collision resistant.

#### What do we need in a transaction
Account based model: (this is what Ethereum uses)
- store list of accounts and balances
- a transaction is valid if there is enough balance in the account
- sender debited, receiver credited

Replay attack - rebroadcasting the same trasaction again and again until a sendter has no money left.

Unspent transaction outputs:
- all coins are not the same
- refer to specific coins when spending
- coins are consumed, create new ones
- a coin can only be spent once

Bitcoin transaction format:
- Input:
-- prev txn ID } identifies an output
-- index       }
-- scriptSig (signature)
they way use spend it produces a new output (coin)
-- value in satoshis 10^8satoshis = bitcoin
-- scriptPubKey

### Pay-to-PubKey-Hash 
(Pay-to-Public-Key-Hash, P2PKH) is the basic form of making a transaction and is the most common form of transaction on the Bitcoin network. Transactions that pay to a Bitcoin address contain P2PKH scripts that are resolved by sending the public key and a digital signature created by the corresponding private key.

What are the transaction models? 
(Unspent Transaction Output) Model and the second one is the Account/Balance Model. The UTXO model is employed by Bitcoin, and Ethereum uses the Account/Balance Model. ... All of the unspent transactions are kept in each fully-synchronized node, and therefore this model is named “UTXO”

### Wallets and SPV
Addresses are 20bytes long. 4 last bytes are checksums

keep private keys offline

---
generate pub key without a private key?
BIP32 Bitcoin improvement protocol 32 simplified:
pubkey P, randomized r, privatekey p

A = P + hash(r, 1) * G -> public key on a curve, G is a generator
a = p + hash(r, 1) -> private key on a curve

- can put pubkey and random data on server
- server can make addresses as needed
- observers can't link the addresses
- revealing P and r would allow linking addresses but not stealing funds

Have I gotten paid?
add your pubkey hashes to a list
for every transaction, look at every output script
- it the scirpt matches your PKH script, you got money

Wallet utxo list:
- keep track of received payment
- save all the utxos to disk
- txid:index, amount, which key, height