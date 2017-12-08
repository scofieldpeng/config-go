package config

import (
	"github.com/vaughan0/go-ini"
)

type (
	// Parser 接口结构体
	Parser interface {
		// Parse 解析方法
		Parse() (map[string]ini.File, error)
	}
	// config config 结构体对象
	config struct {
		// data 存储着具体的 config 参数
		data map[string]ini.File
		// Debug 是否是 debug 模式
		Debug bool
		// Parser 解析器
		parser Parser
	}
)

var (
	// Config 对外提供的config参数
	c = config{}
	// 兼容旧的调用方法,不建议使用
	Config = c.data
	// 版本号
	Version = "v2.0"
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

// Init 初始化
func Init(debug bool, parser ...Parser) (err error) {
	return c.init(debug, parser...)
}

// Reload 重新载入
func Reload() error {
	return c.Load()
}

// Init 初始化配置文件
func (c *config) init(debug bool, parser ...Parser) (err error) {
	c.Debug = debug
	if len(parser) == 0 {
		parser[0] = FileParser{Debug: c.Debug}
	}
	c.parser = parser[0]
	return c.Load()
}

func (c *config) Load() (err error) {
	c.data, err = c.parser.Parse()
	Config = c.data
	return
}
