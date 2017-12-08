package config

import (
	"os"
	"testing"
)

// TestFileParser_Parse 测试文件解析器
func TestFileParser_Parse(t *testing.T) {
	configPath := os.Getenv("GOPATH") + "/src/github.com/zhuziweb/config/config/"
	parser := NewFileParser(configPath)
	parser.Debug = true
	data, err := parser.Parse()
	if err != nil {
		t.Error("debug时解析出错,err:", err.Error())
	}
	version := String(data["test"].Get("app", "version"))
	if version != "v2.1" {
		t.Error("debug 环境下读取的 app:version 值不正确,version:", version)
	}

	parser.Debug = false
	data, err = parser.Parse()
	if err != nil {
		t.Error("release时解析出错,err:", err.Error())
	}
	version = String(data["test"].Get("app", "version"))
	if version != "v2.0" {
		t.Error("release 时读取 app:version 值不正确,version:", version)
	}
}
