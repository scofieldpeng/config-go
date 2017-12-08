package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/vaughan0/go-ini"
	"github.com/zhuziweb/log"
)

type (
	config struct {
		data       map[string]ini.File
		Debug      bool
		ConfigPath string
	}
)

var (
	// Config 对外提供的config参数
	c = config{}
	// 兼容旧的调用方法,不建议使用
	Config = c.data
)

// SetDebug 设置测试环境
func SetDebug(b bool) {
	c.Debug = b
}

// Debug 返回debug状态
func Debug() bool {
	return c.Debug
}

// Data 获取数据
func Data(fileName string) ini.File {
	return c.data[fileName]
}

const (
	// NoramlSuffix 正常配置文件后缀
	NoramlSuffix = ".ini"
	// DebugSuffix debug模式配置文件后缀
	DebugSuffix = "_debug.ini"
)

// Init 初始化配置文件
func (c *config) Init() {
	// 加载所有的配置文件
	pattern := fmt.Sprintf("%s*%s", c.ConfigPath(), NoramlSuffix)

	// 扫描目录找到配置文件
	fileList, err := filepath.Glob(pattern)
	if err != nil {
		log.Panic(err)
	}

	// 开始加载配置文件
	for _, v := range fileList {
		if !c.Debug {
			if strings.Index(v, DebugSuffix) > -1 {
				continue
			}
		}
		c.loadFile(v)
	}
}

// loadFile 加载配置文件
func (c *config) loadFile(filePath string) bool {
	var (
		tmp ini.File
		err error
	)
	for {
		tmp, err = ini.LoadFile(filePath)
		if err != nil {
			if err == os.ErrNotExist && c.Debug {
				filePath = strings.Replace(filePath, DebugSuffix, NoramlSuffix, 1)
			} else {
				log.Warring(err)
				return false
			}
		} else {
			break
		}
	}

	// 计算文件的名称
	c.data[c.fileKey(filePath)] = tmp
	return true
}

// fileKey　通过文件地址获取文件对应的key
// 通过 filePath 参数来获取对应的文件名做为key
func (c *config) fileKey(filePath string) string {
	return strings.Replace(filepath.Base(filePath), c.fileSuffix(filePath), "", -1)
}

// fileSuffix 配置文件对应的后缀
// filePath 当设置了文件路径的时候，通过判断是否有debug文件后缀来确定其使用的是哪种后缀
func (c *config) fileSuffix(filePath ...string) string {
	var fileSuffix string
	if len(filePath) > 0 {
		if strings.Index(filePath[0], DebugSuffix) > 1 {
			fileSuffix = DebugSuffix
		} else {
			fileSuffix = NoramlSuffix
		}
	} else {
		if c.Debug {
			fileSuffix = DebugSuffix
		} else {
			fileSuffix = NoramlSuffix
		}
	}

	return fileSuffix
}

// ConfigPath 获取配置文件存储目录
//func (c *config) ConfigPath() string {
//	return fmt.Sprintf("%s/config/", log.RunDir())
//}
