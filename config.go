package config

import (
	"errors"
	"github.com/vaughan0/go-ini"
	"os"
	"path/filepath"
	"strings"
)

type configT struct {
	nodes map[string]ini.File // 支持的节点
	Debug bool                   // 是否是debug模式
}

const(
	ConfigSuffix string = ".ini" // 配置文件后缀
	DebugSuffix string = "_debug" // debug模式下配置文件后缀
)

var (
	config configT
)

// New 新建一个config结构体对象,传入参数dirPath和
func New(dirPath string, debug bool) error {
	config = configT{
		Debug:debug,
		nodes:make(map[string]ini.File),
	}

	if dirPath == "" {
		dirPath = filepath.Dir(os.Args[0])
	}
	fileObj, err := os.Open(dirPath)
	if err != nil {
		return err
	}

	fileInfo, err := fileObj.Stat()
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return errors.New("dirPath param is not a valid dir!")
	}
	configFiles, err := fileObj.Readdir(0)
	if err != nil {
		return errors.New("read dirPath config file fail!error:" + err.Error())
	}
	// 循环遍历该文件夹下的所有ini文件,将其写入到config.nodes变量里面
	for _,configFile := range configFiles {
		if !configFile.IsDir() && strings.HasSuffix(configFile.Name(),ConfigSuffix) {
			// debug模式下跳过非debug模式的文件,非debug模式则相反
			if IsDebug() {
				if !strings.HasSuffix(configFile.Name(),DebugSuffix + ConfigSuffix) {
					continue
				}
			} else {
				if strings.HasSuffix(configFile.Name(),DebugSuffix + ConfigSuffix) {
					continue
				}
			}
			tmp,err := ini.LoadFile(dirPath + string(os.PathSeparator) + configFile.Name())
			if err != nil {
				return errors.New("read config file,file name:(" + configFile.Name() + ") fail!error:" + err.Error())
			}
			config.nodes[strings.Split(configFile.Name(),".")[0]] = tmp
		}
	}

	return nil
}

// IsDebug 是否是debug模式
func IsDebug() bool {
	return config.Debug
}

// Config 读取配置文件Read config
func Config(fileName string) ini.File {
	if IsDebug() {
		fileName = fileName + "_debug"
	}

	return config.nodes[fileName]
}