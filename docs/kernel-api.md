# Kernel API

Kernel只与API接入点进行交互，不对外提供服务。Kernel将为所有的接入点提供一个session id/secret来完成鉴权。所有的请求，都需要使用session id/secret签发SigningMethodES256的jwt token。

Kernel的主要职责为: a. 读取链上snapshot list; b. 写入transaction; c. 读取公开信息（已注册机构列表, 已注册接入点列表, 业务方列表, 证书类型列表，等）

## Transaction APIs

### Transaction Model

```javascript
{
    "asset": "a99c2e0e2b1da4d648755ef19bd95139acbbe6564cfb06dec7cd34931ca72cdc",
    "extra": "{\"file_hash\":\"xxxx\"}",
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
        "extra": "{\"file_hash\":\"xxxx\"}",
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
    "transaction": {
        "asset": "a99c2e0e2b1da4d648755ef19bd95139acbbe6564cfb06dec7cd34931ca72cdc",
        "extra": "{\"file_hash\":\"xxxx\"}",
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
            "extra": "{\"file_hash\":\"xxxx\"}",
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
        "signature": "xxx"
    }],
    "code": 0
}
```

## 公共数据

### 机构列表

```http
GET /organizations

Content-Type: application/json
Authorization: Bearer **token**
```

**Response:**

```javascript
{
    "code": 0,
    "data": [{
        "name": "派出所",
        "address": "HXxxxx"
    }]
}
```

### 接入点列表

```http
GET /endpoints

Content-Type: application/json
Authorization: Bearer **token**
```

**Response:**

```javascript
{
    "code": 0,
    "data": [{
        "name": "Hengxin 1",
        "endpoint": "https://api1.heng.xin"
    }]
}
```

### 业务方列表

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
        "name": "Hengxin 1",
        "endpoint": "https://service-api1.heng.xin"
    }]
}
```

### 证件列表

```http
GET /cards

Content-Type: application/json
Authorization: Bearer **token**
```

**Response:**

```javascript
{
    "code": 0,
    "data": [{
        "card_type": 1,
        "name": "Hengxin 1",
        "logo": "https://imgs.heng.xin/haha.png"
    }]
}
```
