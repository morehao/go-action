# storagecase 设计文档

## 概述

基于 `github.com/insmtx/storage-go` 统一对象存储抽象层，构建一个覆盖其全部能力的 HTTP 服务，支撑 MinIO、腾讯云 COS、SeaweedFS、Local 四种存储后端，通过配置即可切换。

## 目录结构

```
bizcase/storagecase/
├── main.go              # 入口：创建 gin router，注册路由，启动服务
├── go.mod               # 独立 module
├── config/
│   └── config.go        # Config 结构体 + YAML 加载 + NewStorage() 工厂
├── dto/
│   └── dto.go           # 请求/响应结构体
├── service/
│   └── storage.go       # 业务逻辑层，封装 storage.Storage 调用
└── handler/
    └── storage.go       # HTTP handler 层
```

## 分层职责

| 层 | 职责 |
|---|---|
| main.go | 加载配置，创建 Storage 实例，组装依赖，注册路由，启动 `:8080` |
| config | 读取 YAML 配置，提供 `NewStorage(driver, cfg) (Storage, error)` 工厂 |
| dto | 纯数据结构体，请求 binding tag，响应结构 |
| service | 持有 `storage.Storage` 接口，逐个方法封装，注入到 handler |
| handler | 解析请求参数，调用 service，用 `gincontext.Success/Fail` 统一响应 |

## 路由设计

### Base 路由组 `/api/storage/base`

| 路径 | 说明 |
|------|------|
| `POST /put` | 单次上传（multipart/form-data 传文件） |
| `POST /get` | 下载对象（支持 Range） |
| `POST /delete` | 删除单个对象 |
| `POST /delete-batch` | 批量删除 |
| `POST /list` | 列举对象（支持分页） |

### Multipart 路由组 `/api/storage/multipart`

| 路径 | 说明 |
|------|------|
| `POST /create` | 初始化分片上传 |
| `POST /upload-part` | 上传单个分片（multipart/form-data） |
| `POST /complete` | 合并完成分片上传 |
| `POST /abort` | 取消分片上传 |

### Ext 路由组 `/api/storage/ext`

| 路径 | 说明 |
|------|------|
| `POST /head` | 获取对象元数据 |
| `POST /copy` | 服务端复制 |
| `POST /presign-get` | 预签名下载 URL |
| `POST /presign-put` | 预签名上传 URL |

### 公共

| 路径 | 说明 |
|------|------|
| `GET /api/storage/health` | 健康检查 |

## DTO 设计

### 请求

- `PutObjectReq`: bucket (required), key (required), content_type, metadata, if_not_exists; 文件通过 multipart form-data
- `GetObjectReq`: bucket (required), key (required), range_start, range_end
- `DeleteObjectReq`: bucket (required), key (required)
- `DeleteObjectsReq`: bucket (required), keys (required)
- `ListObjectsReq`: bucket (required), prefix, max_keys, start_after, recursive
- `CreateMultipartUploadReq`: bucket (required), key (required)
- `UploadPartReq`: bucket (required), key (required), upload_id (required), part_number (required); 文件通过 multipart
- `CompleteMultipartUploadReq`: bucket (required), key (required), upload_id (required), parts (required)
- `AbortMultipartUploadReq`: bucket (required), key (required), upload_id (required)
- `HeadObjectReq`: bucket (required), key (required)
- `CopyObjectReq`: src_bucket (required), src_key (required), dst_bucket (required), dst_key (required)
- `PresignGetReq`: bucket (required), key (required), ttl_seconds, range_start, range_end
- `PresignPutReq`: bucket (required), key (required), ttl_seconds

### 响应

- `ObjectDataResp`: content_type, content_length, etag, data (base64)
- `ObjectInfoResp`: path, size, etag, content_type, last_modified, metadata
- `ListObjectsResp`: contents, common_prefixes, is_truncated, next_continuation_token
- `UploadPartResp`: part_number, etag
- `PresignResp`: url

## 配置

YAML 格式，支持四种 driver 切换：

```yaml
driver: minio  # minio | cos | seaweedfs | local
endpoint: play.min.io:9000
access_key: minioadmin
secret_key: minioadmin
bucket: test-bucket
use_ssl: false
region: ""
root_dir: ""       # local driver
http_base_url: ""  # local driver public URL
timeout: 30
max_retries: 3
```

## 依赖

- `github.com/insmtx/storage-go` - 统一存储抽象层
- `github.com/insmtx/storage-go/driver/minio` - MinIO driver
- `github.com/insmtx/storage-go/driver/cos` - COS driver
- `github.com/insmtx/storage-go/driver/seaweedfs` - SeaweedFS driver
- `github.com/insmtx/storage-go/driver/local` - Local driver
- `github.com/gin-gonic/gin` - HTTP 框架
- `github.com/morehao/golib/glog` - 日志
- `github.com/morehao/golib/gcontext/gincontext` - 统一响应
- `github.com/morehao/golib/conf` - 配置加载
- `gopkg.in/yaml.v3` - YAML 解析

## 错误处理

- 初始化失败：panic 退出
- 操作错误：将 storage-go 的 sentinel error（ErrNotFound、ErrAlreadyExists 等）映射到 HTTP 响应中
- 统一使用 glog 记录请求日志和错误日志
