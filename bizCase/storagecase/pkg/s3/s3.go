package s3

import (
	"sync"

	storage "github.com/ygpkg/storage-go"

	_ "github.com/ygpkg/storage-go/driver/cos"
	_ "github.com/ygpkg/storage-go/driver/local"
	_ "github.com/ygpkg/storage-go/driver/minio"
	_ "github.com/ygpkg/storage-go/driver/seaweedfs"
)

var (
	inst storage.Storage
	mu   sync.RWMutex
)

func Init(driver string, cfg storage.Config) error {
	mu.Lock()
	defer mu.Unlock()
	s, err := storage.New(driver, cfg)
	if err != nil {
		return err
	}
	inst = s
	return nil
}

func Client() storage.Storage {
	mu.RLock()
	defer mu.RUnlock()
	return inst
}
