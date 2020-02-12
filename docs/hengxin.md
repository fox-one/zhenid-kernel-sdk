# Endpoint APIs

## Public Data

### Service Provider列表

```http
GET /service-providers

Content-Type: application/json
Authorization: Bearer **token**
```

**Response:**

```javascript
{
    "code": 0,
    "data": [{
        "id": 1,
        "name": "Hengxin 1",
        "endpoint": "https://service-api1.heng.xin"
    }]
}
```

## Transaction

### Snapshot List

```curl
GET /snapshots
```

**Params:**

```form
from: 1234
```

**Response:**

```javascript
{
    "code": 0,
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
    }]
}
```

### Read UTXO

```curl
GET /transactions/utxo
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

### Request UTXO

```curl
POST /transactions/utxo
```

**Param:**

```javascript
{
    "addresses": ["HXxxx"]
}
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
        "index": 0
    }
}
```

### Write Transaction

```http
POST /transaction
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

## File

读写文件时，需要申请该文件使用的Service Provider。sp为SP id。

### Upload File

```http
POST /file

Content-Type: multipart/form-data
```

**Params:**

```form
file: file
permissions: [{"address":"HXxxxx","encrypted_key":"xxx","permission":"READ"}]
sp: 1
```

**Response:**

```javascript
{
    "data": {
        "file_hash": "xxx"
    },
    "code": 0,
}
```

### Read File

```http
GET /files/:file-hash?sp=1

Content-Type: application/json
Authorization: Bearer **token**
```

**Response:**

```javascript
{
    "data": {
        "permissions": [{"address": "HXxxx", "permission": "READ","encrypted_key": "xxx"}],
        "file_type": "URL", // RAW OR URL
        "file": "https://www.file.f/f"
    },
    "code": 0,
}
```
