package config

import (
	"os"
	"testing"
)

func TestConfig_Init(t *testing.T) {
	os.Setenv("TEST_NAME", "helloworld")
	appDir := os.Getenv("GOPATH") + "/src/go.zhuzi.me/config/example/"
	fileParser := NewFileParser(appDir)
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

	// test v2
	fileParser = NewFileParserV2(appDir)
	Init(true, fileParser)
	version = String(Data("test_debug").Get("app", "version"))
	if "v2.1" != version {
		t.Error("v2模式下读取test_debug.ini的version值不正确,version:", version)
	}
	version = String(Data("test").Get("app", "version"))
	if "v2.0" != version {
		t.Error("v2模式下读取test.ini的version值不正确,version:", version)
	}

	// 测试环境变量替换
	fileParser = NewFileParserV1(appDir)
	Init(true, fileParser)
	if envVal := String(Data("test").Get("app", "env")); envVal != "" {
		if envVal != os.Getenv("TEST_NAME") {
			t.Error("环境变量读取失败,env:", envVal,",right:",os.Getenv("TEST_NAME"))
		}
	}
}
