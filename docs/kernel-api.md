# Kernel API

Kernel只与接入点进行交互，不对外提供服务。

Kernel的主要职责为:

- 读取链上snapshot list;
- 写入transaction;
- 读取UTXO信息。

## Transaction APIs

### Extra

```javascript
{
    "h": "xxx",     // file hash
    "exp": 1,       // expired at, unix timestamp, in s
    "sp": 1,        // service provider id
    "csp": 1,       // optional, certificate service provider id
    "c": 1,         // card type
}
```

### Transaction Model

```javascript
{
    "asset": "a99c2e0e2b1da4d648755ef19bd95139acbbe6564cfb06dec7cd34931ca72cdc",
    "extra": "{\"h\":\"xxxx\"}",
    "hash": "xxx",
    "inputs": [{
        "hash": "xxx",
        "index": 1
    }],
    "outputs": [{
        "amount": "0.01",
        "keys": [
            "xxxx"
        ],
        "mask": "xxx",
        "script": "fffe01",
        "type": 0
    }],
    "version": 1
}
```

### Write Transaction

```http
POST /transaction

Content-Type: application/json
Authorization: Bearer **token**
```

**Params:**

```javascript
{
    "transaction": {
        "asset": "xxx",
        "extra": "{\"h\":\"xxxx\"}",
        "hash": "xxx",
        "inputs": [{
            "hash": "xxx",
            "index": 1
        }],
        "outputs": [{
            "amount": "0.01",
            "keys": [
                "xxxx"
            ],
            "mask": "xxx",
            "script": "fffe01",
            "type": 0
        }],
        "version": 1
    },
    "signatures": [["xxx"]]
}
```

**Response:**

```javascript
{
    "code": 0
}
```

### Read Snapshots

```http
GET /snapshots

Content-Type: application/json
Authorization: Bearer **token**
```

**Params:**

```form
since: 1234
```

**Response:**

```javascript
{
    "data": [{
        "transaction": {
            "asset": "a99c2e0e2b1da4d648755ef19bd95139acbbe6564cfb06dec7cd34931ca72cdc",
            "extra": "{\"h\":\"xxxx\"}",
            "hash": "xxx",
            "inputs": [{
                "hash": "xxx",
                "index": 1
            }],
            "outputs": [{
                "amount": "0.01",
                "keys": [
                    "xxxx"
                ],
                "mask": "xxx",
                "script": "fffe01",
                "type": 0
            }],
            "version": 1
        },
        "timestamp": 1575387731172983000,
        "topology": 12792651,
        "transaction_signatures": [["xxx"]],
        "signature": "xxx",
        "version": 1
    }],
    "code": 0
}
```

### Read UTXO

```http
GET /transaction/utxo

Content-Type: application/json
Authorization: Bearer **token**
```

**Param:**

```form
transaction_hash = xxx
index = 0
```

**Response:**

```javascript
{
    "code": 0,
    "data": {
        "output": {
            "amount": "0.01",
            "keys": [
                "xxxx"
            ],
            "mask": "xxx",
            "script": "fffe01",
            "type": 0
        },
        "transaction_hash": "xxx",
        "index": 0,
        "lock": "xxx" // 如果output被消费，则指向消费该output的transaction；否则为空
    }
}
```
