//go:build !enterprise

package impl

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	rootDir string
}

func NewOSS() *LocalStorage {
	return &LocalStorage{rootDir: "/tmp/gotag-storage"}
}

func (l *LocalStorage) Upload(bucket, key string, data []byte) error {
	dir := filepath.Join(l.rootDir, bucket)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}
	return os.WriteFile(filepath.Join(dir, key), data, 0644)
}

func (l *LocalStorage) Download(bucket, key string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(l.rootDir, bucket, key))
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}
	return data, nil
}
