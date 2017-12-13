package config

import (
	"testing"

	"github.com/vaughan0/go-ini"
	config2 "go.zhuzi.me/config/config"
)

// Test_ManualParser 测试manual解析
func Test_ManualParser(t *testing.T) {
	p := config2.NewManualParser()
	p.SetConfig("app", "system", "version", "beta0.1")
	if err := Init(false, p); err != nil {
		t.Error(err)
	}
	if v, ok := Data("app").Get("system", "version"); !ok {
		t.Error("没有找到设置的配置")
	} else if v != "beta0.1" {
		t.Error("找到的值与预设值不一致，设置的值：beta0.1,获取到的值:", v)
	}
	t.Log("test config success")

	p = config2.NewManualParser()
	p.SetSection("app", "system", ini.Section{"appName": "config"})
	if err := Init(false, p); err != nil {
		t.Error(err)
	}
	if v, ok := Data("app").Get("system", "appName"); !ok {
		t.Error("没有找到设置的配置")
	} else if v != "config" {
		t.Error("找到的值与预设值不一致，设置的值：config,获取到的值:", v)
	}
	t.Log("test section success")

	p = config2.NewManualParser()
	p.SetFile("app", ini.File{"system": ini.Section{"author": "scofield"}})
	if err := Init(false, p); err != nil {
		t.Error(err)
	}
	if v, ok := Data("app").Get("system", "author"); !ok {
		t.Error("没有找到设置的配置")
	} else if v != "scofield" {
		t.Error("找到的值与预设值不一致，设置的值：scofield,获取到的值:", v)
	}
	t.Log("test app success")

}
