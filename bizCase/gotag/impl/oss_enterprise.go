//go:build enterprise

package impl

import (
	"fmt"
)

type AliyunOSS struct {
	endpoint string
	bucket   string
}

func NewOSS() *AliyunOSS {
	return &AliyunOSS{
		endpoint: "https://oss-cn-hangzhou.aliyuncs.com",
		bucket:   "enterprise-bucket",
	}
}

func (a *AliyunOSS) Upload(bucket, key string, data []byte) error {
	fmt.Printf("[AliyunOSS] Uploading %s/%s to %s (size: %d bytes)\n", bucket, key, a.endpoint, len(data))
	return nil
}

func (a *AliyunOSS) Download(bucket, key string) ([]byte, error) {
	fmt.Printf("[AliyunOSS] Downloading %s/%s from %s\n", bucket, key, a.endpoint)
	return []byte("enterprise-data"), nil
}
