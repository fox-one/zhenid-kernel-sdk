# Public Data API

Authority Center，作为Hengxin Network的顶层机构，负责公共数据的审核，发布。所有的接入点，CSP，SP上线之前都需要向Authority Center发出申请，待Authority Center审核，发布之后，才可以上线。Authority Center亦将其审核过的信息上链。

## 接入点列表

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

## Service Provider列表

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

## CSP 机构列表

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

## 证件类型信息(CSP注册) 列表

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
        "card_id": 1,
        "name": "Hengxin 1",
        "logo": "https://imgs.heng.xin/haha.png"
    }]
}
```
