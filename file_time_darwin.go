// +build darwin

package fcache

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

type FileTime struct {
	CreateTime int64
	AccessTime int64
	ModifyTime int64
}

func GetFileTime(file string) (*FileTime, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if stats, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
		return &FileTime{
			CreateTime: time.Unix(stats.Ctimespec.Unix()).UnixNano(),
			AccessTime: time.Unix(stats.Atimespec.Unix()).UnixNano(),
			ModifyTime: time.Unix(stats.Mtimespec.Unix()).UnixNano(),
		}, nil
	}
	return nil, fmt.Errorf("not support file info in current platform")
}
