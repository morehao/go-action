package handler

import (
	"context"
	"errors"
	"io"
	"time"

	storage "github.com/ygpkg/storage-go"
	"github.com/morehao/go-action/bizcase/storagecase/dto"
	"github.com/morehao/go-action/bizcase/storagecase/service"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"

	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	svc *service.StorageService
}

func NewStorageHandler() *StorageHandler {
	return &StorageHandler{svc: service.NewStorageService()}
}

func (h *StorageHandler) Health(c *gin.Context) {
	gincontext.Success(c, gin.H{"status": "ok"})
}

func (h *StorageHandler) PutObject(c *gin.Context) {
	ctx := context.Background()
	var req dto.PutObjectReq
	if err := c.ShouldBind(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		gincontext.Fail(c, err)
		return
	}

	result, err := h.svc.PutObject(ctx, req.Bucket, req.Key, req.ContentType, req.Metadata, req.IfNotExists, body)
	if err != nil {
		glog.Errorf(ctx, "PutObject failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	glog.Infof(ctx, "PutObject success: %s", result.Path.URI())
	gincontext.Success(c, toObjectInfo(result.Path, 0, result.ETag, "", time.Time{}, req.Metadata))
}

func (h *StorageHandler) GetObject(c *gin.Context) {
	ctx := context.Background()
	var req dto.GetObjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	result, err := h.svc.GetObject(ctx, req.Bucket, req.Key, req.RangeStart, req.RangeEnd)
	if err != nil {
		glog.Errorf(ctx, "GetObject failed: %v", err)
		gincontext.Fail(c, err)
		return
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		gincontext.Fail(c, err)
		return
	}

	resp := dto.ObjectDataResp{
		ContentType:   result.ContentType,
		ContentLength: result.ContentLength,
		ETag:          result.ETag,
		Data:          data,
	}
	gincontext.Success(c, resp)
}

func (h *StorageHandler) DeleteObject(c *gin.Context) {
	ctx := context.Background()
	var req dto.DeleteObjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	if err := h.svc.DeleteObject(ctx, req.Bucket, req.Key); err != nil {
		glog.Errorf(ctx, "DeleteObject failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	glog.Infof(ctx, "DeleteObject success: %s/%s", req.Bucket, req.Key)
	gincontext.Success(c, nil)
}

func (h *StorageHandler) DeleteObjects(c *gin.Context) {
	ctx := context.Background()
	var req dto.DeleteObjectsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	err := h.svc.DeleteObjects(ctx, req.Bucket, req.Keys)
	if err != nil {
		glog.Errorf(ctx, "DeleteObjects partial failure: %v", err)
		resp := buildDeleteObjectsResp(req.Keys, err)
		gincontext.Success(c, resp)
		return
	}

	glog.Infof(ctx, "DeleteObjects success: %d keys from %s", len(req.Keys), req.Bucket)
	gincontext.Success(c, dto.DeleteObjectsResp{DeletedCount: len(req.Keys)})
}

func (h *StorageHandler) ListObjects(c *gin.Context) {
	ctx := context.Background()
	var req dto.ListObjectsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	result, err := h.svc.ListObjects(ctx, req.Bucket, req.Prefix, req.MaxKeys, req.StartAfter, req.Recursive)
	if err != nil {
		glog.Errorf(ctx, "ListObjects failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	resp := dto.ListObjectsResp{
		CommonPrefixes:        result.CommonPrefixes,
		IsTruncated:           result.IsTruncated,
		NextContinuationToken: result.NextContinuationToken,
	}
	for _, obj := range result.Contents {
		resp.Contents = append(resp.Contents, toObjectInfo(obj.Path, obj.Size, obj.ETag, obj.ContentType, obj.LastModified, nil))
	}
	gincontext.Success(c, resp)
}

func (h *StorageHandler) CreateMultipartUpload(c *gin.Context) {
	ctx := context.Background()
	var req dto.CreateMultipartUploadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	uploadID, err := h.svc.CreateMultipartUpload(ctx, req.Bucket, req.Key)
	if err != nil {
		glog.Errorf(ctx, "CreateMultipartUpload failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	glog.Infof(ctx, "CreateMultipartUpload success: %s/%s uploadID=%s", req.Bucket, req.Key, uploadID)
	gincontext.Success(c, dto.CreateMultipartUploadResp{UploadID: uploadID})
}

func (h *StorageHandler) UploadPart(c *gin.Context) {
	ctx := context.Background()
	var req dto.UploadPartReq
	if err := c.ShouldBind(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		gincontext.Fail(c, err)
		return
	}

	part, err := h.svc.UploadPart(ctx, req.Bucket, req.Key, req.UploadID, req.PartNumber, body)
	if err != nil {
		glog.Errorf(ctx, "UploadPart failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	glog.Infof(ctx, "UploadPart success: %s/%s part=%d", req.Bucket, req.Key, req.PartNumber)
	gincontext.Success(c, dto.UploadPartResp{PartNumber: part.PartNumber, ETag: part.ETag})
}

func (h *StorageHandler) CompleteMultipartUpload(c *gin.Context) {
	ctx := context.Background()
	var req dto.CompleteMultipartUploadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	parts := make([]storage.CompletedPart, len(req.Parts))
	for i, p := range req.Parts {
		parts[i] = storage.CompletedPart{PartNumber: p.PartNumber, ETag: p.ETag}
	}

	if err := h.svc.CompleteMultipartUpload(ctx, req.Bucket, req.Key, req.UploadID, parts); err != nil {
		glog.Errorf(ctx, "CompleteMultipartUpload failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	glog.Infof(ctx, "CompleteMultipartUpload success: %s/%s", req.Bucket, req.Key)
	gincontext.Success(c, nil)
}

func (h *StorageHandler) AbortMultipartUpload(c *gin.Context) {
	ctx := context.Background()
	var req dto.AbortMultipartUploadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	if err := h.svc.AbortMultipartUpload(ctx, req.Bucket, req.Key, req.UploadID); err != nil {
		glog.Errorf(ctx, "AbortMultipartUpload failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	glog.Infof(ctx, "AbortMultipartUpload success: %s/%s", req.Bucket, req.Key)
	gincontext.Success(c, nil)
}

func (h *StorageHandler) HeadObject(c *gin.Context) {
	ctx := context.Background()
	var req dto.HeadObjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	info, err := h.svc.HeadObject(ctx, req.Bucket, req.Key)
	if err != nil {
		glog.Errorf(ctx, "HeadObject failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	gincontext.Success(c, toObjectInfo(info.Path, info.Size, info.ETag, info.ContentType, info.LastModified, info.Metadata))
}

func (h *StorageHandler) CopyObject(c *gin.Context) {
	ctx := context.Background()
	var req dto.CopyObjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	if err := h.svc.CopyObject(ctx, req.SrcBucket, req.SrcKey, req.DstBucket, req.DstKey); err != nil {
		glog.Errorf(ctx, "CopyObject failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	glog.Infof(ctx, "CopyObject success: %s/%s -> %s/%s", req.SrcBucket, req.SrcKey, req.DstBucket, req.DstKey)
	gincontext.Success(c, nil)
}

func (h *StorageHandler) PresignGetObject(c *gin.Context) {
	ctx := context.Background()
	var req dto.PresignGetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	ttl := defaultTTL(req.TTLSeconds)
	url, err := h.svc.PresignGetObject(ctx, req.Bucket, req.Key, ttl, req.RangeStart, req.RangeEnd)
	if err != nil {
		glog.Errorf(ctx, "PresignGetObject failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	gincontext.Success(c, dto.PresignResp{URL: url})
}

func (h *StorageHandler) PresignPutObject(c *gin.Context) {
	ctx := context.Background()
	var req dto.PresignPutReq
	if err := c.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(c, err)
		return
	}

	ttl := defaultTTL(req.TTLSeconds)
	url, err := h.svc.PresignPutObject(ctx, req.Bucket, req.Key, ttl)
	if err != nil {
		glog.Errorf(ctx, "PresignPutObject failed: %v", err)
		gincontext.Fail(c, err)
		return
	}

	gincontext.Success(c, dto.PresignResp{URL: url})
}

func toObjectInfo(path storage.StoragePath, size int64, etag, contentType string, lastModified time.Time, metadata map[string]string) dto.ObjectInfoResp {
	resp := dto.ObjectInfoResp{
		Size:     size,
		ETag:     etag,
		ContentType: contentType,
		Metadata: metadata,
	}
	if path != nil {
		resp.Path = path.URI()
	}
	if !lastModified.IsZero() {
		resp.LastModified = lastModified.Format(time.RFC3339)
	}
	return resp
}

func defaultTTL(seconds int) time.Duration {
	if seconds <= 0 {
		return 15 * time.Minute
	}
	return time.Duration(seconds) * time.Second
}

func buildDeleteObjectsResp(keys []string, err error) dto.DeleteObjectsResp {
	var bulkErr *storage.BulkDeleteError
	if !errors.As(err, &bulkErr) {
		return dto.DeleteObjectsResp{}
	}

	failureMap := make(map[string]string)
	for _, f := range bulkErr.Failures {
		failureMap[f.Key] = f.Err.Error()
	}

	resp := dto.DeleteObjectsResp{}
	for _, k := range keys {
		if msg, ok := failureMap[k]; ok {
			resp.Failures = append(resp.Failures, dto.DeleteFailure{Key: k, Error: msg})
		} else {
			resp.DeletedCount++
		}
	}
	return resp
}
