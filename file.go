package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/vaughan0/go-ini"
)

// FileParser 文件解析
type FileParser struct {
	Path string
}

var (
	// 文件没有找到的错误
	FileNotExist = errors.New("file not exist")
)

// Parse 解析
func (f FileParser) Parse() (data ini.File, err error) {
	data, err = ini.LoadFile(f.Path)
	if err != nil {
		if err == os.ErrNotExist {
			err = FileNotExist
		}
	}
	return
}
