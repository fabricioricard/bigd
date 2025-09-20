# Release version pktd-v1.7.0
Jan 23, 2024

## Major changes

### New voting system!
This version adds voting capability to pktwallet. It adds a new RPC endpoint `sendvote`
which is like `sendfrom` but rather than sending any significant amount of coins, it
sends a *vote* transaction. It selects enough UTXOs from the "from" address to pay the
fees, and cycles all the rest back as change.

Example:

```bash
./bin/pktctl --wallet sendvote pkt1qj6nldg0klm5pqyam8cg6lvzcuq432kpc5d9scn pkt1qpl48jdwyn3t0uj7lu303f5szejgl5dmc0p2vpu
```

In this example, `pkt1qj6nldg0klm5pqyam8cg6lvzcuq432kpc5d9scn` is my address, the one
I am voting with, and `pkt1qpl48jdwyn3t0uj7lu303f5szejgl5dmc0p2vpu` is the address I
am voting for.

There are also three more optional arguments to `sendvote`:
* `is_candidate` (boolean): If true, this wallet will not only vote but also represent
itself as a candidate to potentially become the Network Steward. Default: false
* `minconf` (uint32): Do not source funds from any payment with less than this number of
confirmations. Default 1 = anything which is in the blockchain at all.
* `maxinputs` (uint32): Do not source funds from any more inputs than this. Default: no
specific limit, use the wallet's internal limits.
* `minheight` (uint32): Do not source funds from any payments OLDER (lower block height)
than this number. Default 0 = no limit

After you have cast a vote, you will be able to see your current vote for each address
in your `getaddressbalances` output.

`user@armee pktd % ./bin/pktctl --wallet getaddressbalances`
```json
[
  {
    "address": "pkt1qj6nldg0klm5pqyam8cg6lvzcuq432kpc5d9scn",
    "total": 99.99999986309558,
    "stotal": "107374182253",
    "spendable": 99.99999986309558,
    "sspendable": "107374182253",
    "immaturereward": 0,
    "simmaturereward": "0",
    "unconfirmed": 0,
    "sunconfirmed": "0",
    "outputcount": 1,
    "vote": {
      "vote_for": "pkt1qpl48jdwyn3t0uj7lu303f5szejgl5dmc0p2vpu",
      "vote_txid": "4d590a6f45d1f992bea8e97995539c789e942425c40c32f928cc5308f16ccdd2",
      "vote_block": 2336429,
      "expiration_block": 2860589,
      "estimated_expiration_sec": 1737462218
    }
  }
]
```

The `vote` field is an optional field, it only appears if a vote has been registered.
If a new vote is cast which replaces the old vote, the vote field will be updated. Old votes
are not remembered.

The `vote_for` field is relatively self-explanitory. The `vote_txid` field is the transaction
ID of the vote transaction, in this example it is transaction
[4d590a6f45d1f992...](https://packetscan.io/tx/4d590a6f45d1f992bea8e97995539c789e942425c40c32f928cc5308f16ccdd2).
The `vote_block` field is the block where the vote transaction was stored. The vote will not
appear in the `getaddressbalances` until it has landed in a block. The `expiration_block` is
the block number at which the vote will expire (if it is not renewed). The
`estimated_expiration_sec` field is an estimation of the time when the expiration block will
come. It is not exact but is based on the blockchain's target time window. In this case, the
vote is expected to expire on January 21 or 2025, Almost exactly one year after it was cast.

You can re-vote or change your vote at any time and the new vote will take the place of the
old vote. To renew a vote before expiration, simply re-vote for the same candidate.

### Faster UTXO table
This version merges some code that was written in mid-2022 to improve the performance of UTXO
scans. Previously each UTXO data access required also accessing a transaction record in order
to get the full information about the UTXO. This change significantly improves performance
when scanning over a large utxo set.

There are no user-facing changes to this, but it does irreversibly update the database. This
update is already present in the PKT-Lightning-Wallet so this update makes it possible to
switch back and forth between pktwallet and PKT-Lightning-Wallet.

# Release version pktd-v1.6.3
Jan 4, 2024

This is a patch release which introduces the `--ignoremined` flag so that exchanges
and merchants can protect their wallets from micro-transaction spam.

# Release version pktd-v1.6.2
Jan 2, 2024

This is a patch release to introduce the capability to create a transaction without
unlocking the wallet, but only if the transaction is unsigned.

# Release version pktd-v1.6.1
Oct 26, 2022

## Major changes

### Fix PacketCryptProof mutation Denial of Service
Until this version, it was possible to mutate (add extranious data to) a PacketCryptProof
and while this would not be spread across the network, any node who downloads blocks *from
you* would store that data on their disk.

The maximum size of a message is 32MB, so an attacker can add a maximum of 32MB of
extranious data to a victim's harddrive per-block.

#### Impacts

1. An attacker may be able to expend up to 46 GB of storage space on a victim's pktd server
per day.
2. If the victim is syncing a full node for the first time and using the attacker's server,
the attacker can expend far more than this.
3. If the attacker includes malware binary samples and the victim uses antivirus, it may be
possible that they download a malware sample to disk while the antivirus is inactive, and
then upon activating it, the antivirus quarentines the block file, disrupting the function
of pktd.

#### Change
As of 1.6.1, PacketCryptProofs which are "non-standard" (contain any extra data not needed
to validate them) will fail verification.

This change is NOT a soft fork because any block which would have been accepted before
still will, only with the PacketCryptProof represented in its standard form.

#### Special thanks
Thank you to Rob Daniell for identifying this issue and creating a proof of concept.

## Minor changes

* Fix of a null dereference crash in pld.
* Always check PacketCrypt proofs, even if the block is behind a checkpoint, in order to
prevent mutation attack.

# Release version pktd-v1.6.0
July 5, 2022

## Major changes

### New RPC system in pld
New RPC system. The RPC system of LND is GRPC, the problem with GRPC is that it's
difficult to use and the software is not maintained as well as the developers at Google
would have you believe.

In one example, the Java GRPC client can trigger a crash (nil deref segfault) in the Golang
GRPC server by making a request. It's worth noting that a mature RPC library should be
impossible to crash, even intentionally with malicious requests. The fact that the Golang
GRPC server is crashable not just by a malformed request but in fact by a normal request
made by another GRPC library tells us that there's a total lack of quality assurance.

Secondly, the GRPC client has abysmal performance on Android, and the iOS AnodeVPN team
was not able to get it working in their React Native stack, at all.

But GRPC is not just buggy and difficult to integrate, the API is also "opinionated" in
that it requires the use of it's own internal HTTP/2 server rather than offering an API
by which request data can be passed in. This makes it impossible to integrate into an
existing server which also does other things, and makes it married to the HTTP/2 protocol
which is not always available, for example in a web browser based client without HTTPS
encryption, or in our case, in a React Native based app.

Finally, GRPC is not at all conducive to testing and debugging. It relies exclusively on
Protocol Buffers binary representation of data, which is nice for sending data around in
production, but you can't just make requests using a tool like curl.

So we replaced GRPC with a simple http REST server, it uses the same protobuf structures
but accepts both binary protobuf and also JSON using jsonpb to encode and decode it.



### Major reduction of disk usage in neutrino db
The first version of neutrino db, the version from LND, used a combination of a database
and a couple of flat files. Unfortunately, synchronizing a db transaction with a flat file
write operation is incredibly difficult to do correctly, and the developers of LND
unfortunately did NOT do it correctly.

The result was data corruption when the wallet was hard stopped (by way of killing the
process or shutting off the device) and a lot of complex code for trying to manage
atomicity, without actually getting it right!

We deleted almost all of this code, switching to use of the neutrino.db database
completely which is based on bbolt database, a transactional key/value store inspired by
LMDB.

This solved the corruption problem, but the resulting database became quite large. The
neutrino db needs to be able to do a few things.

1. Get block header by block height OR block hash
2. Get filter header by block height OR block hash

The way we implemented all of these things was by creating a table of block header by hash,
filter header by hash, and hash by height. This worked but it used a lot of storage.
The hash was included three times as well as the block header and filter header.

In version 1.6.0 we replaced this schema with a new schema which stores block header and
filter header by height, and which stores height by 32 bit hash *prefix*. Because we only
store 32 bits of the hash, we accept that there may be a hash collision - i.e. multiple
block hashes which have the same first 32 bits. To resolve this, we store a *list* of
block heights for a given hash, and when doing a lookup, we check each one of the headers
and compute it's hash until we find the hash which matches. On average we will only have
a single entry in the list, but by being prepared for multiple entries we make the storage
size much smaller.

With these changes we were able to cut the neutrino.db file size roughly in half. But be
aware that when you update, the file will not get smaller because bboltdb is unable to
shrink the file, it can only create free space inside of the file. If you want the file
to be smaller, you will need to delete the neutrino.db file and allow pktwallet / pld to
re-create it.


## Minor changes

1. pld: Feature: Add pktdir param to pld - so you can specify the folder where the wallet is - f39d8d0dd8cf27f5ff0858abc0470b24e6903bdf
2. pktd: Bug: Fixed occasional crash caused by misuse of locks - 15c277ef8aa0c50d271fa68fb1650ce713bced84
3. pld: Bug: If some necessary directories don't exist, create them rather than crashing. 2a3949932b18a38ae5ca0deb35cc6a5567d3df5c 2a3949932b18a38ae5ca0deb35cc6a5567d3df5c
4. all: Improvement: Add github build process - c475471eff27c5a7963b5eccd13214c9fdf58013
5. pld: Bug: Fix how password is specified - 3af1c059a5716f8f0ffab86ed7cffa6c0fde14fb
6. pld: Bug: Private password must be verified when unlocking wallet or else there will be a crash later - 67e09e45328f59fca487e16443137d23d0446b99
7. all: Improvement: Remove non-working DNS seed nodes - 7d77cf273c25c43182dfff9ff657f2b629c93dc3
8. all: Improvement: Add seed nodes seed.pkt.ai and seed.pkt.pink - 281535f3b7b06261602e8edc97ee4325d84d0cc9 35ac19d2fb99d2f66c03b92e6082ab67b3c56677
9. pld: Improvement: Remove graceful shutdown so that the process will not try ot 
10. pld: Feature: Added ability to generate a wallet seed without creating the wallet.

## Stats
The pktd codebase unfortunately contains a fairly significant amount of *generated* code.
This exists because the tooling to generate it is too awkward to include in the build
scripts. I have separated the stats for generated code from the stats for hand-written
code.

Over this development period, a total of 215 files were changed, there were 82803 lines
of code altered or added. From these, 42 files and 46499 lines of generated code were
changed or added.

The size of the project increased by 5645 lines of written code, and 20373 lines of
generated code.

A total of 36304 lines, or 6.3% of the codebase, was manually altered.

pktd-v1.5.1 - 575425 lines of written code, 40319 lines of generated code
pktd-v1.6.0 - 581070 lines of written code, 60692 lines of generated code

## Future goals
In coming releases, we will be seeking to:
* Debug and stabilize Lightning Network
* Improve scalability
* Improve test coverage
* Move code generation to build time
* Reduce total hand-written code


## Thanks!
This release has received contributions from:
* Ivo Smits
* Aldebaran Perseke
* Dimitris Koukoulakis
* James P. Thomas
* Kyle Mason
