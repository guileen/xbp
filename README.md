
[![Build Status](https://travis-ci.org/flyrpc/flyrpc.svg?branch=master)](https://travis-ci.org/flyrpc/flyrpc)
[![Coverage Status](https://coveralls.io/repos/flyrpc/flyrpc/badge.svg?branch=master)](https://coveralls.io/r/flyrpc/flyrpc?branch=master)

XXBP miXed teXt and Binary Protocol.

# Protocol

## Packet Spec

|Name   | Flag   | TLength | BLength | CRC16   |  Sequence  |Text    | Binary  |
|-------|:------:|:-------:|:-------:|:-------:|:----------:|:------:|:-------:|
|Bytes  | 1      | 0,1,2,4 | 0,1,2,4 | 2       | 2          | string | *       |

### Flag Spec

| 1 - 2 | 3      | 4         | 5 - 6             | 7 - 8     |
|-------|--------|-----------|-------------------|--------------|
| Mode  |        |           | Log(Bytes of Text Length)+1 | Log(Bytes of Binary Length)+1 |

### Mode

* 00  Message   every message is standalone message.
* 01  Request   request message should be responsed. every request should have increment sequence, 0 ~ 65535.
* 10  Response  response message should match request sequence.
* 11  Extend    TODO

# FAQ
* Why provide both text and binary in one packet?

    Command and payload, is common requirment.

* Will text and binary make packet bigger?

    It will not. You can send packet only text or only binary.

# API

```
opts.setCRC16(enabled) // default false.
opts.setCRC16Salt(bytes) // default none.

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
