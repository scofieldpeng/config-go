package config

import (
	"os"

	"errors"

	"fmt"
	"path/filepath"

	"strings"

	"os/exec"

	"io/ioutil"
	"regexp"

	"github.com/vaughan0/go-ini"
)

// FileParser 文件解析
type FileParser struct {
	Path    string
	Debug   bool
	Version string
}

var (
	// 文件没有找到的错误
	ErrFileNotExist = errors.New("file not exist")
)

const (
	// releaseSuffix 正常配置文件后缀
	releaseSuffix = ".ini"
	// debugSuffix debug模式配置文件后缀
	debugSuffix = "_debug.ini"
)

const (
	// v1版本
	V1 = "v1"
	// v2版本
	V2 = "v2"
)

// NewFileParser 新建文件 Parser
func NewFileParser(debug bool, configPath ...string) FileParser {
	parser := NewFileParserV1(debug, configPath...)

	version := os.Getenv("CONFIG_VERSION")
	if version != "" {
		switch strings.ToLower(version) {
		case V2:
			parser.Version = V2
		case V1:
			parser.Version = V1
		default:
			parser.Version = V1
		}
	}

	fmt.Println("default version:", parser.Version)

	return parser
}

func NewFileParserV1(debug bool, configPath ...string) FileParser {
	if len(configPath) == 0 {
		configPath = make([]string, 1)
		configPath[0] = FileParser{}.defaultConfigPath()
	}
	return FileParser{
		Path:    configPath[0],
		Version: V1,
		Debug:   debug,
	}
}

// 新建文件Parser（v2）版本
func NewFileParserV2(debug bool, configPath ...string) FileParser {
	p := NewFileParser(debug, configPath...)
	p.Version = V2

	return p
}

// Parse 解析
func (f FileParser) Parse() (data map[string]ini.File, err error) {
	data = make(map[string]ini.File)
	var (
		tmpData ini.File
	)
	if f.Path == "" {
		f.Path = f.defaultConfigPath()
	}
	if !strings.HasSuffix(f.Path, "/") {
		f.Path = f.Path + "/"
	}
	var (
		tmpFileList = make([]string, 0)
		fileList    = make([]string, 0)
	)
	tmpFileList, err = filepath.Glob(fmt.Sprintf("%s*%s", f.Path, releaseSuffix))
	if err != nil {
		return
	}

	// 读取配置文件的版本需要根据实现不同来实现，比如说如何根据
	// 当为v1的时候，debug环境会读取xxx_debug.ini文件
	fmt.Println("version:", f.Version)
	for _, v := range tmpFileList {
		if f.Version == V1 {
			if strings.HasSuffix(v, debugSuffix) && !f.Debug {
				continue
			}
			if f.Debug && !strings.HasSuffix(v, debugSuffix) {
				continue
			}
		}
		fileList = append(fileList, v)
	}
	for _, v := range fileList {
		if tmpData, err = f.parseFile(v); err != nil {
			return
		}
		data[f.getFileKey(v)] = tmpData
	}

	return
}

// parseFile 解析文件
func (f FileParser) parseFile(path string) (data ini.File, err error) {
	var (
		file      *os.File
		fileBytes []byte
		fileStr   string
	)
	if file, err = os.Open(path); err != nil {
		if err == os.ErrNotExist {
			err = ErrFileNotExist
		}
		return
	}
	defer file.Close()

	// 检测是否有环境变量，如果有值，替换掉
	// ${ENV}则为读取env的值进行替换
	// TODO 这里有很多优化空间，鉴于时间和各种成本因素，后期考虑
	if fileBytes, err = ioutil.ReadAll(file); err != nil {
		return
	}
	fileStr = string(fileBytes)
	reg := regexp.MustCompile(`\$\{([a-zA-Z0-9\-\_]+)\}`)
	findRes := reg.FindAllStringSubmatch(fileStr, 1)
	for _, v := range findRes {
		if len(v) > 1 {
			if env := os.Getenv(v[1]); env != "" {
				fileStr = strings.Replace(fileStr, v[0], env, -1)
			}
		}
	}

	// 解析配置文件
	data, err = ini.Load(strings.NewReader(fileStr))
	return
}

// getFileKey 获取文件的 key,key 是指 app.ini 中 key的值为 app
func (f FileParser) getFileKey(filePath string) string {
	suffix := releaseSuffix
	if f.Version == V1 && f.Debug {
		suffix = debugSuffix
	}
	return strings.Replace(filepath.Base(filePath), suffix, "", -1)
}

// runDir 当前运行的目录
func (f FileParser) runDir() string {
	rootDir, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	rootDir, err = filepath.Abs(rootDir)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(rootDir)
}

// defaultConfigPath 默认的 config 路径
func (f FileParser) defaultConfigPath() string {
	return fmt.Sprintf("%s/config/", f.runDir())
}
