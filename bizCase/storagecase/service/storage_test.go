package service

import (
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	storage "github.com/ygpkg/storage-go"
	"github.com/morehao/go-action/bizcase/storagecase/pkg/s3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBucket = "test"

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "storagecase-test-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	if err := s3.Init("local", storage.Config{
		Bucket:   testBucket,
		BaseDir: dir,
	}); err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestPutGetObject(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()
	body := []byte("hello world")

	result, err := svc.PutObject(ctx, testBucket, "putget.txt", "text/plain", map[string]string{"author": "test"}, false, body)
	require.NoError(t, err)
	assert.NotEmpty(t, result.ETag)
	assert.NotEmpty(t, result.Path.URI())

	get, err := svc.GetObject(ctx, testBucket, "putget.txt", 0, 0)
	require.NoError(t, err)
	defer get.Body.Close()

	data, err := io.ReadAll(get.Body)
	require.NoError(t, err)
	assert.Equal(t, body, data)
	assert.Equal(t, "text/plain", get.ContentType)
}

func TestGetObjectWithRange(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()
	body := []byte("0123456789")

	_, err := svc.PutObject(ctx, testBucket, "range.txt", "", nil, false, body)
	require.NoError(t, err)

	get, err := svc.GetObject(ctx, testBucket, "range.txt", 0, 100)
	require.NoError(t, err)
	defer get.Body.Close()

	data, _ := io.ReadAll(get.Body)
	assert.Equal(t, body, data)
}

func TestGetObjectNotFound(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.GetObject(ctx, testBucket, "notexists.txt", 0, 0)
	assert.True(t, errors.Is(err, storage.ErrNotFound))
}

func TestDeleteObject(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.PutObject(ctx, testBucket, "del.txt", "", nil, false, []byte("data"))
	require.NoError(t, err)

	err = svc.DeleteObject(ctx, testBucket, "del.txt")
	assert.NoError(t, err)

	err = svc.DeleteObject(ctx, testBucket, "del.txt")
	assert.NoError(t, err)
}

func TestDeleteObjects(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	keys := []string{"b1.txt", "b2.txt", "b3.txt"}
	for _, k := range keys {
		_, err := svc.PutObject(ctx, testBucket, k, "", nil, false, []byte("data"))
		require.NoError(t, err)
	}

	err := svc.DeleteObjects(ctx, testBucket, keys)
	assert.NoError(t, err)

	for _, k := range keys {
		err := svc.DeleteObject(ctx, testBucket, k)
		assert.NoError(t, err)
	}
}

func TestListObjects(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.PutObject(ctx, testBucket, "list/a.txt", "", nil, false, []byte("a"))
	require.NoError(t, err)
	_, err = svc.PutObject(ctx, testBucket, "list/b.txt", "", nil, false, []byte("b"))
	require.NoError(t, err)

	result, err := svc.ListObjects(ctx, testBucket, "list/", 10, "", false)
	require.NoError(t, err)
	assert.Len(t, result.Contents, 2)
	assert.Contains(t, result.Contents[0].Path.URI(), "list/")
}

func TestListObjectsPagination(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	for _, k := range []string{"page/a.txt", "page/b.txt", "page/c.txt"} {
		_, err := svc.PutObject(ctx, testBucket, k, "", nil, false, []byte("data"))
		require.NoError(t, err)
	}

	result, err := svc.ListObjects(ctx, testBucket, "page/", 2, "", false)
	require.NoError(t, err)
	assert.Len(t, result.Contents, 2)
	assert.True(t, result.IsTruncated)
	assert.NotEmpty(t, result.NextContinuationToken)
}

func TestHeadObject(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.PutObject(ctx, testBucket, "head.txt", "text/plain", nil, false, []byte("head"))
	require.NoError(t, err)

	info, err := svc.HeadObject(ctx, testBucket, "head.txt")
	require.NoError(t, err)
	assert.Equal(t, "text/plain", info.ContentType)
	assert.Equal(t, int64(4), info.Size)
	assert.NotEmpty(t, info.ETag)
}

func TestHeadObjectNotFound(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.HeadObject(ctx, testBucket, "nohead.txt")
	assert.True(t, errors.Is(err, storage.ErrNotFound))
}

func TestCopyObject(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()
	body := []byte("copy me")

	_, err := svc.PutObject(ctx, testBucket, "src.txt", "", nil, false, body)
	require.NoError(t, err)

	err = svc.CopyObject(ctx, testBucket, "src.txt", testBucket, "dst.txt")
	require.NoError(t, err)

	get, err := svc.GetObject(ctx, testBucket, "dst.txt", 0, 0)
	require.NoError(t, err)
	defer get.Body.Close()

	data, _ := io.ReadAll(get.Body)
	assert.Equal(t, body, data)
}

func TestCopyObjectNotFound(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	err := svc.CopyObject(ctx, testBucket, "nosrc.txt", testBucket, "dst.txt")
	assert.True(t, errors.Is(err, storage.ErrNotFound))
}

func TestMultipartUpload(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	uploadID, err := svc.CreateMultipartUpload(ctx, testBucket, "multi.txt")
	require.NoError(t, err)
	assert.NotEmpty(t, uploadID)

	part1, err := svc.UploadPart(ctx, testBucket, "multi.txt", uploadID, 1, []byte("part1-"))
	require.NoError(t, err)
	assert.Equal(t, 1, part1.PartNumber)
	assert.NotEmpty(t, part1.ETag)

	part2, err := svc.UploadPart(ctx, testBucket, "multi.txt", uploadID, 2, []byte("part2"))
	require.NoError(t, err)
	assert.Equal(t, 2, part2.PartNumber)

	err = svc.CompleteMultipartUpload(ctx, testBucket, "multi.txt", uploadID, []storage.CompletedPart{
		*part1, *part2,
	})
	require.NoError(t, err)

	get, err := svc.GetObject(ctx, testBucket, "multi.txt", 0, 0)
	require.NoError(t, err)
	defer get.Body.Close()

	data, _ := io.ReadAll(get.Body)
	assert.Equal(t, []byte("part1-part2"), data)
}

func TestAbortMultipartUpload(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	uploadID, err := svc.CreateMultipartUpload(ctx, testBucket, "abort.txt")
	require.NoError(t, err)

	_, err = svc.UploadPart(ctx, testBucket, "abort.txt", uploadID, 1, []byte("data"))
	require.NoError(t, err)

	err = svc.AbortMultipartUpload(ctx, testBucket, "abort.txt", uploadID)
	assert.NoError(t, err)
}

func TestPutObjectWithMetadata(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	meta := map[string]string{"x-custom": "myvalue"}
	_, err := svc.PutObject(ctx, testBucket, "meta.txt", "", meta, false, []byte("meta"))
	require.NoError(t, err)

	info, err := svc.HeadObject(ctx, testBucket, "meta.txt")
	require.NoError(t, err)
	assert.Equal(t, "myvalue", info.Metadata["x-custom"])
}

func TestPutObjectIfNotExists(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.PutObject(ctx, testBucket, "ifne.txt", "", nil, true, []byte("first"))
	require.NoError(t, err)

	_, err = svc.PutObject(ctx, testBucket, "ifne.txt", "", nil, true, []byte("second"))
	assert.True(t, errors.Is(err, storage.ErrAlreadyExists))
}

func TestPresignGetObjectNotSupported(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.PresignGetObject(ctx, testBucket, "key", time.Minute, 0, 0)
	assert.True(t, errors.Is(err, storage.ErrNotSupported))
}

func TestPresignPutObjectNotSupported(t *testing.T) {
	svc := NewStorageService()
	ctx := context.Background()

	_, err := svc.PresignPutObject(ctx, testBucket, "key", time.Minute)
	assert.True(t, errors.Is(err, storage.ErrNotSupported))
}
