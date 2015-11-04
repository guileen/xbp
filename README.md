
[![Build Status](https://travis-ci.org/flyrpc/flyrpc.svg?branch=master)](https://travis-ci.org/flyrpc/flyrpc)
[![Coverage Status](https://coveralls.io/repos/flyrpc/flyrpc/badge.svg?branch=master)](https://coveralls.io/r/flyrpc/flyrpc?branch=master)

XBP teXt Binary Protocol.

# What the protocol defined.

* Asynchronous Request/Response
* Request with unlimited string code and unlimited binary payload.
* Response with unlimited string code and unlimited binary payload.
* Compress code and payload
* Request can have an implicit `ack response` or `no response`.

# Protocol

## Packet Spec

|Name   | Flag   | Sequence |Code    | Length | Payload |
|-------|:------:|:--------:|:------:|:------:|:-------:|
|Bytes  | 1      | 2        |string\0| 1,2,4,8| *       |

### Flag Spec

| 1      | 2           | 3 | 4 | 5      | 6         | 7 - 8        |
|--------|-------------|---|---|--------|-----------|--------------|
|Response|Request      |   |   |Zip Code|Zip Payload| length bytes |

# API

```
conn.onPacket(flag, seq, code, payload)
conn.sendPacket(flag, seq, code, payload)

conn.sendMessage(code, payload)
conn.sendRequest(code, payload)
conn.sendResponse(seq, code, payload)

conn.onMessage(function handler(seq, code, payload))
conn.onRequest(function handler(seq, code, payload))
conn.onResponse(function handler(seq, code, payload))

conn.request(code, payload, function callback(err, code, payload))
conn.handle(code, function handler(payload, function reply(code, payload)))
```
