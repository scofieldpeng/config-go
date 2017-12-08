package config

import (
	"os"
	"testing"
)

func TestConfig_Init(t *testing.T) {
	fileParser := NewFileParser(os.Getenv("GOPATH") + "/src/github.com/zhuziweb/config/config/")
	fileParser.Debug = true
	if err := Init(true, fileParser); err != nil {
		t.Error(err)
	}
	version := String(Data("test").Get("app", "version"))
	if "v2.1" != version {
		t.Error("debug 模式下读取的 version 不对,version:", version)
	}
	fileParser.Debug = false
	if err := Init(false, fileParser); err != nil {
		t.Error(err)
	}
	version = String(Data("test").Get("app", "version"))
	if "v2.0" != version {
		t.Error("release 模式下读取的 version 不对,version:", version)
	}
}
