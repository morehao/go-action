package dto

// PutObjectReq 上传对象
type PutObjectReq struct {
	Bucket      string            `json:"bucket" form:"bucket" binding:"required"`  // 必填: 存储桶
	Key         string            `json:"key" form:"key" binding:"required"`        // 必填: 对象键
	ContentType string            `json:"content_type" form:"content_type"`         // 对象 MIME 类型
	Metadata    map[string]string `json:"metadata" form:"metadata"`                 // 对象元数据
	IfNotExists bool              `json:"if_not_exists" form:"if_not_exists"`       // 仅当对象不存在时上传
}

// GetObjectReq 下载对象
type GetObjectReq struct {
	Bucket     string `json:"bucket" binding:"required"`    // 必填: 存储桶
	Key        string `json:"key" binding:"required"`       // 必填: 对象键
	RangeStart int64  `json:"range_start"`                  // 范围读取起始字节偏移
	RangeEnd   int64  `json:"range_end"`                    // 范围读取结束字节偏移
}

// DeleteObjectReq 删除单个对象
type DeleteObjectReq struct {
	Bucket string `json:"bucket" binding:"required"`  // 必填: 存储桶
	Key    string `json:"key" binding:"required"`     // 必填: 对象键
}

// DeleteObjectsReq 批量删除对象
type DeleteObjectsReq struct {
	Bucket string   `json:"bucket" binding:"required"`     // 必填: 存储桶
	Keys   []string `json:"keys" binding:"required,min=1"` // 必填: 待删除对象键列表
}

// ListObjectsReq 列举对象
type ListObjectsReq struct {
	Bucket     string `json:"bucket" binding:"required"`  // 必填: 存储桶
	Prefix     string `json:"prefix"`                     // 对象键前缀过滤
	MaxKeys    int64  `json:"max_keys"`                   // 最大返回数量
	StartAfter string `json:"start_after"`                // 从该对象键之后开始列举
	Recursive  bool   `json:"recursive"`                  // 是否递归列举子目录
}

// CreateMultipartUploadReq 创建分片上传
type CreateMultipartUploadReq struct {
	Bucket string `json:"bucket" binding:"required"`  // 必填: 存储桶
	Key    string `json:"key" binding:"required"`     // 必填: 对象键
}

// UploadPartReq 上传分片
type UploadPartReq struct {
	Bucket     string `json:"bucket" binding:"required"`        // 必填: 存储桶
	Key        string `json:"key" binding:"required"`           // 必填: 对象键
	UploadID   string `json:"upload_id" binding:"required"`     // 必填: 分片上传 ID
	PartNumber int    `json:"part_number" binding:"required,min=1"`  // 必填: 分片编号（从1开始）
}

// CompleteMultipartUploadReq 完成分片上传
type CompleteMultipartUploadReq struct {
	Bucket   string          `json:"bucket" binding:"required"`   // 必填: 存储桶
	Key      string          `json:"key" binding:"required"`      // 必填: 对象键
	UploadID string          `json:"upload_id" binding:"required"`// 必填: 分片上传 ID
	Parts    []CompletedPart `json:"parts" binding:"required"`    // 必填: 已完成的分片列表
}

type CompletedPart struct {
	PartNumber int    `json:"part_number"`
	ETag       string `json:"etag"`
}

// AbortMultipartUploadReq 取消分片上传
type AbortMultipartUploadReq struct {
	Bucket   string `json:"bucket" binding:"required"`   // 必填: 存储桶
	Key      string `json:"key" binding:"required"`      // 必填: 对象键
	UploadID string `json:"upload_id" binding:"required"`// 必填: 分片上传 ID
}

// HeadObjectReq 获取对象元数据
type HeadObjectReq struct {
	Bucket string `json:"bucket" binding:"required"`  // 必填: 存储桶
	Key    string `json:"key" binding:"required"`     // 必填: 对象键
}

// CopyObjectReq 复制对象
type CopyObjectReq struct {
	SrcBucket string `json:"src_bucket" binding:"required"`  // 必填: 源存储桶
	SrcKey    string `json:"src_key" binding:"required"`     // 必填: 源对象键
	DstBucket string `json:"dst_bucket" binding:"required"`  // 必填: 目标存储桶
	DstKey    string `json:"dst_key" binding:"required"`     // 必填: 目标对象键
}

// PresignGetReq 生成预签名下载 URL
type PresignGetReq struct {
	Bucket     string `json:"bucket" binding:"required"`  // 必填: 存储桶
	Key        string `json:"key" binding:"required"`     // 必填: 对象键
	TTLSeconds int    `json:"ttl_seconds"`                // URL 有效期（秒），默认 3600
	RangeStart int64  `json:"range_start"`                // 范围读取起始字节偏移
	RangeEnd   int64  `json:"range_end"`                  // 范围读取结束字节偏移
}

// PresignPutReq 生成预签名上传 URL
type PresignPutReq struct {
	Bucket     string `json:"bucket" binding:"required"`  // 必填: 存储桶
	Key        string `json:"key" binding:"required"`     // 必填: 对象键
	TTLSeconds int    `json:"ttl_seconds"`                // URL 有效期（秒），默认 3600
}

// ObjectDataResp 对象下载响应
type ObjectDataResp struct {
	ContentType   string `json:"content_type"`    // 对象 MIME 类型
	ContentLength int64  `json:"content_length"`  // 对象字节数
	ETag          string `json:"etag"`            // 对象 ETag，用于校验内容一致性
	LastModified  string `json:"last_modified"`   // 最后修改时间
	Path          string `json:"path"`            // 对象存储路径 URI
	Data          []byte `json:"data"`            // 对象内容二进制数据（JSON 中 Base64 编码）
}

// ObjectInfoResp 对象元数据
type ObjectInfoResp struct {
	Path         string            `json:"path"`
	Size         int64             `json:"size"`
	ETag         string            `json:"etag"`
	ContentType  string            `json:"content_type"`
	LastModified string            `json:"last_modified"`
	Metadata     map[string]string `json:"metadata"`
	PublicURL    string            `json:"public_url"`
}

// ListObjectsResp 列举对象响应
type ListObjectsResp struct {
	Contents               []ObjectInfoResp `json:"contents"`
	CommonPrefixes         []string         `json:"common_prefixes"`
	IsTruncated            bool             `json:"is_truncated"`
	NextContinuationToken  string           `json:"next_continuation_token"`
}

// UploadPartResp 上传分片响应
type UploadPartResp struct {
	PartNumber int    `json:"part_number"`
	ETag       string `json:"etag"`
}

// CreateMultipartUploadResp 创建分片上传响应
type CreateMultipartUploadResp struct {
	UploadID string `json:"upload_id"`
}

// PresignResp 预签名 URL 响应
type PresignResp struct {
	URL string `json:"url"`
}

// DeleteObjectsResp 批量删除响应
type DeleteObjectsResp struct {
	DeletedCount int             `json:"deleted_count"`
	Failures     []DeleteFailure `json:"failures,omitempty"`
}

// DeleteFailure 删除失败信息
type DeleteFailure struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}
