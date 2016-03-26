package config

import (
	"fmt"
	"git.name.im/bamboo/log.git"
	"github.com/howeyc/fsnotify"
	"github.com/vaughan0/go-ini"
	"path/filepath"
	"strings"
)

type conf map[string]ini.File

// 对外提供的config参数
var Config  = conf{}

// 是否为测试环境
var debug bool

// SetDebug 设置测试环境
func SetDebug(b bool) {
	debug = b
}

// Init 初始化配置文件
func (c *conf) Init() {
	// 加载所有的配置文件
	var pattern string
	if debug {
		pattern = fmt.Sprintf("%s%s", c.ConfigPath(), "*_debug.ini")
	} else {
		pattern = fmt.Sprintf("%s%s", c.ConfigPath(), "*.ini")
	}

	// 扫描目录找到配置文件
	fileList, err := filepath.Glob(pattern)
	if err != nil {
		log.Panic(err)
	}

	// 开始加载配置文件
	for _, v := range fileList {
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
	tmp, err := ini.LoadFile(filePath)
	if err != nil {
		log.Warring(err)
		return false
	}
	// 计算文件的名称
	(*c)[c.fileKey(filePath)] = tmp
	return true
}

// fileKey　通过文件地址获取文件对应的key
func (c *conf) fileKey(filePath string) string {
	return strings.Replace(filepath.Base(filePath), c.fileSuffix(), "", -1)
}

// fileSuffix 配置文件后缀
func (c *conf) fileSuffix() string {
	var fileSuffix string
	if debug {
		fileSuffix = "_debug.ini"
	} else {
		fileSuffix = ".ini"
	}
	return fileSuffix
}

// ConfigPath 获取配置文件存储目录
func (c *conf) ConfigPath() string {
	return fmt.Sprintf("%s/config/", log.RunDir())
}
