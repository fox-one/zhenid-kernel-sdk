# Service Provider API

业务方的职责主要包含：a. 数据仓库功能; b. 辅助完成多人合同签署。数据仓库功能主要为写入文件，增加/删除文件读写权限，读取文件；辅助完成多人合同签署主要为申请签发合同，收集并聚合签名。

业务方的鉴权基于用户的Private Encrypt Key进行。请求时，使用Private Encrypt Key签发的SigningMethodES256的jwt token，并附带在http header "Authorization"中。

## 数据仓库

### 上传文件

```http
POST /files

Content-Type: multipart/form-data
Authorization: Bearer **token**
```

**Params:**

```form
file: file
ciphers: [{"a":"HXxxxx","k":"xxx","p":"READ"}]
```

**Response:**

```javascript
{
  "code": 0,
  "data": {
    "file_hash": "61f954807ce17b252ceb170110bca4a575cb0e0035b0551072ef163ead8c78a0",
    "file_size": 91
  }
}
```

### 读取文件

```http
GET /files/:file-hash/:file-size

Content-Type: application/json
Authorization: Bearer **token**
```

**Response:**

```javascript
{
  "code": 0,
  "data": {
    "url": "https://xxx",
    "file_hash": "61f954807ce17b252ceb170110bca4a575cb0e0035b0551072ef163ead8c78a0",
    "file_size": 91,
    "ciphers": []
  }
}
```

### 修改文件权限

```http
PATCH /files/:file_hash/:file-size/ciphers

Content-Type: application/json
Authorization: Bearer **token**
```

**Params:**

```javascript
[{
    "action": "ADD",
    "a": "HXxxx",       // Hengxin Address
    "k": "xxx",         // encrypted keys
    "p":"READ"          // permission: READ or ADMIN
}, {
    "action": "DEL",
    "a": "HXxxx"
}]
```

**Response:**

```javascript
{
    "code": 0
}
```

### 删除文件

```http
DELETE /files/:file_hash/:file-size

Content-Type: application/json
Authorization: Bearer **token**
```

**Response:**

```javascript
{
    "code": 0
}
```

### 申请签发文件

## 多人联合签署合同 (TODO)

### 收集签名
