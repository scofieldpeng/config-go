package config

import (
	"os"
	"testing"
)

// TestFileParser_Parse 测试文件解析器
func TestFileParser_Parse(t *testing.T) {
	configPath := os.Getenv("GOPATH") + "/src/github.com/zhuziweb/config/config/"
	parser := NewFileParser(true, configPath)
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

func TestEnv_Parse(t *testing.T) {
	f := FileParser{}
	os.Setenv("NAME", "scofield")
	os.Setenv("AGE", "26")
	if res := f.replaceEnvVar("${NAME}"); res != "scofield" {
		t.Error("replace name fail,wrong value:", res)
	}
	if res := f.replaceEnvVar("${NAME}${AGE}"); res != "scofield26" {
		t.Error("replace name,age fail!,wrong value:", res)
	}
	testValue3 := `
[user]
default=name:${NAME},age:${AGE},city:${LOCATION:=china,sichuan,成都},sex:男
`
	res := f.replaceEnvVar(testValue3)
	t.Log(res)
}
