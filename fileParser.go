package config

import (
	"os"

	"errors"

	"fmt"
	"path/filepath"

	"strings"

	"os/exec"

	"github.com/vaughan0/go-ini"
)

// FileParser 文件解析
type FileParser struct {
	Path  string
	Debug bool
}

var (
	// 文件没有找到的错误
	FileNotExist = errors.New("file not exist")
)

const (
	// releaseSuffix 正常配置文件后缀
	releaseSuffix = ".ini"
	// debugSuffix debug模式配置文件后缀
	debugSuffix = "_debug.ini"
)

// NewFileParser 新建文件 Parser
func NewFileParser(configPath ...string) FileParser {
	if len(configPath) == 0 {
		configPath = make([]string, 1)
		configPath[0] = FileParser{}.defaultConfigPath()
	}
	return FileParser{
		Path: configPath[0],
	}
}

// Parse 解析
func (f FileParser) Parse() (data map[string]ini.File, err error) {
	data = make(map[string]ini.File)
	var (
		tmpData ini.File
	)
	if f.Path == "" {
		f.Path = f.defaultConfigPath()
	}
	if !strings.HasSuffix(f.Path, "/") {
		f.Path = f.Path + "/"
	}
	var (
		tmpFileList = make([]string, 0)
		fileList    = make([]string, 0)
	)
	tmpFileList, err = filepath.Glob(fmt.Sprintf("%s*%s", f.Path, releaseSuffix))
	if err != nil {
		return
	}

	// 如果是线上环境，只读取线上的 ini,debug 环境只读取 debug 环境
	for _, v := range tmpFileList {
		if strings.HasSuffix(v, debugSuffix) && !f.Debug {
			continue
		}
		if f.Debug && !strings.HasSuffix(v, debugSuffix) {
			continue
		}
		fileList = append(fileList, v)
	}
	for _, v := range fileList {
		if tmpData, err = f.parseFile(v); err != nil {
			return
		}
		data[f.getFileKey(v)] = tmpData
	}

	return
}

// parseFile 解析文件
func (f FileParser) parseFile(path string) (data ini.File, err error) {
	data, err = ini.LoadFile(path)
	if err != nil {
		if err == os.ErrNotExist {
			err = FileNotExist
		}
	}
	return
}

// getFileKey 获取文件的 key,key 是指 app.ini 中 key的值为 app
func (f FileParser) getFileKey(filePath string) string {
	suffix := releaseSuffix
	if f.Debug {
		suffix = debugSuffix
	}
	return strings.Replace(filepath.Base(filePath), suffix, "", -1)
}

// runDir 当前运行的目录
func (f FileParser) runDir() string {
	rootDir, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	rootDir, err = filepath.Abs(rootDir)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(rootDir)
}

// defaultConfigPath 默认的 config 路径
func (f FileParser) defaultConfigPath() string {
	return fmt.Sprintf("%s/config/", f.runDir())
}
