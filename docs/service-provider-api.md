# Service Provider API

业务方的职责主要包含：a. 数据仓库功能; b. 辅助完成多人合同签署。数据仓库功能主要为写入文件，增加/删除文件读写权限，读取文件；辅助完成多人合同签署主要为申请签发合同，收集并聚合签名。

业务方的鉴权基于用户的Private Encrypt Key进行。请求时，使用Private Encrypt Key签发的SigningMethodES256的jwt token，并附带在http header "Authorization"中。

## 数据仓库

### 上传文件

```http
POST /file

Content-Type: multipart/form-data
Authorization: Bearer **token**
```

**Params:**

```form
file: file
permissions: [{"address":"HXxxxx","encrypted_key":"xxx","permission":"READ"}]
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

### 更新文件权限

```http
PUT /file/:file-hash

Content-Type: application/json
Authorization: Bearer **token**
```

**Params:**

```javascript
[{
    "action": "ADD",
    "address": "HXxxx",
    "encrypted_key": "xxx",
    "permission":"READ" // READ or ADMIN
}, {
    "action": "DEL",
    "address": "HXxxx"
}]
```

**Response:**

```javascript
{
    "data": {
        "permissions": [{"address": "HXxxx", "permission": "READ"}]
    },
    "code": 0,
}
```

### 读取文件

```http
GET /file/:file-hash

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

## 多人联合签署合同 (TODO)

### 申请签发文件

### 收集签名
