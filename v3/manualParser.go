package config

import "github.com/vaughan0/go-ini"

// ManualParser 手动解析器
type ManualParser struct {
	config map[string]ini.File
}

// NewManualParser 新建手动解析器
func NewManualParser() ManualParser {
	return ManualParser{
		config: make(map[string]ini.File),
	}
}

// SetFile 设置file的值
func (mp *ManualParser) SetFile(fileName string, data ini.File) {
	mp.config[fileName] = data
}

// SetSection 设置section的值
func (mp *ManualParser) SetSection(fileName, section string, data ini.Section) {
	if _, ok := mp.config[fileName]; !ok {
		mp.config[fileName] = make(ini.File)
	}
	mp.config[fileName][section] = data
}

// SetConfig 设置value
func (mp *ManualParser) SetConfig(appName, section, key, value string) {
	if _, ok := mp.config[appName]; !ok {
		mp.config[appName] = make(ini.File)
	}
	if _, ok := mp.config[appName][section]; !ok {
		mp.config[appName][section] = make(ini.Section)
	}
	mp.config[appName][section][key] = value
}

// Parse 解析
func (mp ManualParser) Parse() (data map[string]ini.File, err error) {
	data = mp.config
	return
}
