package config

import (
	"fmt"
	"git.name.im/bamboo/log.git"
	"github.com/howeyc/fsnotify"
	"github.com/vaughan0/go-ini"
	"os"
	"path/filepath"
	"strings"
)

type conf map[string]ini.File

// 对外提供的config参数
var Config = conf{}

// 是否为测试环境
var debug bool

// SetDebug 设置测试环境
func SetDebug(b bool) {
	debug = b
	log.SetDebug(b)
}

const (
	NORAML_SUFFIX = ".ini"       // 正常配置文件后缀
	DEBUG_SUFFIX  = "_debug.ini" // debug模式配置文件后缀
)

// Init 初始化配置文件
func (c *conf) Init() {
	// 加载所有的配置文件
	pattern := fmt.Sprintf("%s*%s", c.ConfigPath(), NORAML_SUFFIX)

	// 扫描目录找到配置文件
	fileList, err := filepath.Glob(pattern)
	if err != nil {
		log.Panic(err)
	}

	// 开始加载配置文件
	for _, v := range fileList {
		if !debug {
			if strings.Index(v, DEBUG_SUFFIX) > -1 {
				continue
			}
		}
		c.loadFile(v)
	}

	// 文件加载完成的时候，开始监听文件的变化
	go c.Watch()

}

// Watcher 监听文件的变化
func (c *conf) Watch() {
	configPath := c.ConfigPath()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Panic(err)
	}
	defer watcher.Close()

	err = watcher.Watch(configPath)
	if err != nil {
		log.Panic(err)
	}

	fileSuffix := c.fileSuffix()

	for {
		select {
		case v := <-watcher.Event:
			// 判断是否为配置文件变化
			if !strings.HasSuffix(v.Name, fileSuffix) {
				break
			}

			// 当为删除、重命名操作时，删除对应的配置项
			if v.IsDelete() || v.IsRename() {
				delete((*c), c.fileKey(v.Name))
			}

			// 当为修改、创建操作时
			if v.IsModify() || v.IsCreate() {
				c.loadFile(v.Name)
			}
		}
	}

}

// loadFile 加载配置文件
func (c *conf) loadFile(filePath string) bool {
	var tmp ini.File
	var err error
	for {
		tmp, err = ini.LoadFile(filePath)
		if err != nil {
			if err == os.ErrNotExist && debug {
				filePath = strings.Replace(filePath, DEBUG_SUFFIX, NORAML_SUFFIX, 1)
			} else {
				log.Warring(err)
				return false
			}
		} else {
			break
		}
	}

	// 计算文件的名称
	(*c)[c.fileKey(filePath)] = tmp
	return true
}

// fileKey　通过文件地址获取文件对应的key
// 通过 filePath 参数来获取对应的文件名做为key
func (c *conf) fileKey(filePath string) string {
	return strings.Replace(filepath.Base(filePath), c.fileSuffix(filePath), "", -1)
}

// fileSuffix 配置文件对应的后缀
// filePath 当设置了文件路径的时候，通过判断是否有debug文件后缀来确定其使用的是哪种后缀
func (c *conf) fileSuffix(filePath ...string) string {
	var fileSuffix string
	if len(filePath) > 0 {
		if strings.Index(filePath[0], DEBUG_SUFFIX) > 1 {
			fileSuffix = DEBUG_SUFFIX
		} else {
			fileSuffix = NORAML_SUFFIX
		}
	} else {
		if debug {
			fileSuffix = DEBUG_SUFFIX
		} else {
			fileSuffix = NORAML_SUFFIX
		}
	}

	return fileSuffix
}

// ConfigPath 获取配置文件存储目录
func (c *conf) ConfigPath() string {
	return fmt.Sprintf("%s/config/", log.RunDir())
}
