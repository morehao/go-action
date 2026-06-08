package dto

type PutObjectReq struct {
	Bucket      string            `json:"bucket" binding:"required"`
	Key         string            `json:"key" binding:"required"`
	ContentType string            `json:"content_type"`
	Metadata    map[string]string `json:"metadata"`
	IfNotExists bool              `json:"if_not_exists"`
}

type GetObjectReq struct {
	Bucket     string `json:"bucket" binding:"required"`
	Key        string `json:"key" binding:"required"`
	RangeStart int64  `json:"range_start"`
	RangeEnd   int64  `json:"range_end"`
}

type DeleteObjectReq struct {
	Bucket string `json:"bucket" binding:"required"`
	Key    string `json:"key" binding:"required"`
}

type DeleteObjectsReq struct {
	Bucket string   `json:"bucket" binding:"required"`
	Keys   []string `json:"keys" binding:"required,min=1"`
}

type ListObjectsReq struct {
	Bucket     string `json:"bucket" binding:"required"`
	Prefix     string `json:"prefix"`
	MaxKeys    int64  `json:"max_keys"`
	StartAfter string `json:"start_after"`
	Recursive  bool   `json:"recursive"`
}

type CreateMultipartUploadReq struct {
	Bucket string `json:"bucket" binding:"required"`
	Key    string `json:"key" binding:"required"`
}

type UploadPartReq struct {
	Bucket     string `json:"bucket" binding:"required"`
	Key        string `json:"key" binding:"required"`
	UploadID   string `json:"upload_id" binding:"required"`
	PartNumber int    `json:"part_number" binding:"required,min=1"`
}

type CompleteMultipartUploadReq struct {
	Bucket   string          `json:"bucket" binding:"required"`
	Key      string          `json:"key" binding:"required"`
	UploadID string          `json:"upload_id" binding:"required"`
	Parts    []CompletedPart `json:"parts" binding:"required"`
}

type CompletedPart struct {
	PartNumber int    `json:"part_number"`
	ETag       string `json:"etag"`
}

type AbortMultipartUploadReq struct {
	Bucket   string `json:"bucket" binding:"required"`
	Key      string `json:"key" binding:"required"`
	UploadID string `json:"upload_id" binding:"required"`
}

type HeadObjectReq struct {
	Bucket string `json:"bucket" binding:"required"`
	Key    string `json:"key" binding:"required"`
}

type CopyObjectReq struct {
	SrcBucket string `json:"src_bucket" binding:"required"`
	SrcKey    string `json:"src_key" binding:"required"`
	DstBucket string `json:"dst_bucket" binding:"required"`
	DstKey    string `json:"dst_key" binding:"required"`
}

type PresignGetReq struct {
	Bucket     string `json:"bucket" binding:"required"`
	Key        string `json:"key" binding:"required"`
	TTLSeconds int    `json:"ttl_seconds"`
	RangeStart int64  `json:"range_start"`
	RangeEnd   int64  `json:"range_end"`
}

type PresignPutReq struct {
	Bucket     string `json:"bucket" binding:"required"`
	Key        string `json:"key" binding:"required"`
	TTLSeconds int    `json:"ttl_seconds"`
}

type ObjectDataResp struct {
	ContentType   string `json:"content_type"`
	ContentLength int64  `json:"content_length"`
	ETag          string `json:"etag"`
	Data          []byte `json:"data"`
}

type ObjectInfoResp struct {
	Path         string            `json:"path"`
	Size         int64             `json:"size"`
	ETag         string            `json:"etag"`
	ContentType  string            `json:"content_type"`
	LastModified string            `json:"last_modified"`
	Metadata     map[string]string `json:"metadata"`
}

type ListObjectsResp struct {
	Contents               []ObjectInfoResp `json:"contents"`
	CommonPrefixes         []string         `json:"common_prefixes"`
	IsTruncated            bool             `json:"is_truncated"`
	NextContinuationToken  string           `json:"next_continuation_token"`
}

type UploadPartResp struct {
	PartNumber int    `json:"part_number"`
	ETag       string `json:"etag"`
}

type CreateMultipartUploadResp struct {
	UploadID string `json:"upload_id"`
}

type PresignResp struct {
	URL string `json:"url"`
}

type DeleteObjectsResp struct {
	DeletedCount int              `json:"deleted_count"`
	Failures     []DeleteFailure  `json:"failures,omitempty"`
}

type DeleteFailure struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}
