package service

import (
	"bytes"
	"context"
	"time"

	storage "github.com/ygpkg/storage-go"
	"github.com/morehao/go-action/bizcase/storagecase/pkg/s3"
)

type StorageService struct {
}

func NewStorageService() *StorageService {
	return &StorageService{}
}

func (svc *StorageService) client() storage.Storage {
	return s3.Client()
}

func (svc *StorageService) PutObject(ctx context.Context, bucket, key, contentType string, metadata map[string]string, ifNotExists bool, body []byte) (*storage.PutObjectResult, error) {
	var opts []storage.PutOption
	if contentType != "" {
		opts = append(opts, storage.WithContentType(contentType))
	}
	if len(metadata) > 0 {
		opts = append(opts, storage.WithMetadata(metadata))
	}
	if ifNotExists {
		opts = append(opts, storage.WithIfNotExists())
	}
	return svc.client().PutObject(ctx, bucket, key, bytes.NewReader(body), opts...)
}

func (svc *StorageService) GetObject(ctx context.Context, bucket, key string, rangeStart, rangeEnd int64) (*storage.GetObjectResult, error) {
	var opts []storage.GetOption
	if rangeStart != 0 || rangeEnd != 0 {
		opts = append(opts, storage.WithByteRange(rangeStart, rangeEnd))
	}
	return svc.client().GetObject(ctx, bucket, key, opts...)
}

func (svc *StorageService) DeleteObject(ctx context.Context, bucket, key string) error {
	return svc.client().DeleteObject(ctx, bucket, key)
}

func (svc *StorageService) DeleteObjects(ctx context.Context, bucket string, keys []string) error {
	return svc.client().DeleteObjects(ctx, bucket, keys)
}

func (svc *StorageService) CreateMultipartUpload(ctx context.Context, bucket, key string) (string, error) {
	return svc.client().CreateMultipartUpload(ctx, bucket, key)
}

func (svc *StorageService) UploadPart(ctx context.Context, bucket, key, uploadID string, partNumber int, body []byte) (*storage.CompletedPart, error) {
	return svc.client().UploadPart(ctx, bucket, key, uploadID, partNumber, bytes.NewReader(body))
}

func (svc *StorageService) CompleteMultipartUpload(ctx context.Context, bucket, key, uploadID string, parts []storage.CompletedPart) error {
	return svc.client().CompleteMultipartUpload(ctx, bucket, key, uploadID, parts)
}

func (svc *StorageService) AbortMultipartUpload(ctx context.Context, bucket, key, uploadID string) error {
	return svc.client().AbortMultipartUpload(ctx, bucket, key, uploadID)
}

func (svc *StorageService) HeadObject(ctx context.Context, bucket, key string) (*storage.ObjectInfo, error) {
	return svc.client().HeadObject(ctx, bucket, key)
}

func (svc *StorageService) CopyObject(ctx context.Context, srcBucket, srcKey, dstBucket, dstKey string) error {
	return svc.client().CopyObject(ctx, srcBucket, srcKey, dstBucket, dstKey)
}

func (svc *StorageService) PresignGetObject(ctx context.Context, bucket, key string, ttl time.Duration, rangeStart, rangeEnd int64) (string, error) {
	var opts []storage.GetOption
	if rangeStart != 0 || rangeEnd != 0 {
		opts = append(opts, storage.WithByteRange(rangeStart, rangeEnd))
	}
	return svc.client().PresignGetObject(ctx, bucket, key, ttl, opts...)
}

func (svc *StorageService) PresignPutObject(ctx context.Context, bucket, key string, ttl time.Duration) (string, error) {
	return svc.client().PresignPutObject(ctx, bucket, key, ttl)
}

// ListObjects 列举对象
func (svc *StorageService) ListObjects(ctx context.Context, bucket, prefix string, maxKeys int64, startAfter string, recursive bool) (*storage.ListObjectsOutput, error) {
	var opts []storage.ListOption
	if maxKeys > 0 {
		opts = append(opts, storage.WithMaxKeys(maxKeys))
	}
	if startAfter != "" {
		opts = append(opts, storage.WithStartAfter(startAfter))
	}
	if recursive {
		opts = append(opts, storage.WithRecursive(true))
	}
	return svc.client().ListObjects(ctx, bucket, prefix, opts...)
}

