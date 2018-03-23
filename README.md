# config

一款简单的配置服务,使用.ini 文件来作为配置源

## 快速使用

```go
go get github.com/zhuziweb/config
```

### 1. 初始化

```go
if err := config.Init(true);err != nil {
    log.Info(err)
    os.Exit(1)
}
```

在应用同级目录下创建 config 目录，专门用于管理配置文件，例如：

```bash
/example-app
 |- /config
 |    |- app.ini # release 环境会读取到该文件
 |    |- app_debug.ini # debug 环境会读取到该文件 
 |- app.go    
```

### 2. 使用

假设现在是 debug 模式，app_debug.ini 有如下配置

```ini
# app_debug.ini
[info]
version=1.0
```

```go
// 第一个参数返回 version 值，如果不存在，第二个参数返回 false
// 其中的Data("app")中的 app 值为 app.ini（debug 模式下会读取 app_debug.ini 下的值）
version,ok := config.Data("app").Get("info","version")

// 另一种快捷方法
version := config.String(config.Data("app").Get("info","version"))
```

## 重载配置

```go
err := config.Reload()
```

## 快速获取对应类型的值

**字符:**

```go
string := config.String(config.Config["app"].Get("mysql", "default"))
```

**整数:**

```go
int := config.Int(config.Config["app"].Get("mysql", "default"))
```


**浮点数:**

```go
float64 := config.Float64(config.Config["app"].Get("mysql", "default"))
```

**布尔值:**

```gp
bool := config.bool(config.Config["app"].Get("mysql", "default"))
```

## 自定义配置文件目录

```go
// 设置debug环境
debug := false
// 路径必须为绝对路径，并且以/结尾
absolutePath := `/home/namer/app/config/`
config.Init( debug, NewFileParser(debug,absolutePath))
```

## 文件解析器的debug模式和非debug模式

在使用文件作为config时，默认情况下debug模式下会读取配置目录下的xxx_debug.ini文件，非debug读取xxx.ini文件

有时候我们不想进行_debug.ini和非_debug.ini的区分，如果有此需求，采用v2的默认解析器(文件解析)即可,用法

方法1： 添加环境变量CONFIG_VERSION=v2的值即可

方法2: 初始化时用v2进行初始化:

```go
config.Init(debug,config.NewFileParserV2(debug))
```

## 环境变量配置文件

某些时候，我们的某些配置是是需要写成环境变量实现动态生成配置，默认的文件解析器支持环境变量，语法为${ENV_NAME},例如有下列配置文件

```ini
# app.ini
[redis]
default=${REDIS_DEFAULT_ADDR}
```

上面的${REDIS_DEFAULT_ADDR}代表该值需要用环境变量来替换，如果在运行时指代了环境变量，那么config在初始化后会默认将该值替换成对应的值

有时候我们需要当环境变量不存在时设置默认值，那么可以设置为这样

```ini
# app.ini
[redis]
default=${REDIS_DEFAULT_ADDR:=127.0.0.1:6379}
```

当找不到REDIS_DEFAULT_ADDR这个环境变量时，default值会用127.0.0.1:6379去填充

> 快捷设置环境变量
> 在配置文件的同目录下建立`.env`文件，一行一个，KEY=VALUE格式即可，程序启动时会自动将该文件的值设置到环境变量

## 具体 API

详见[godoc.org/github.com/zhuziweb/config](godoc.org/github.com/zhuziweb/config)

## 添加手动解析

只要实现了Parser结构体即可:

```go
package config

// Parser 接口结构体
type Parser interface {
    // Parse 解析方法
    Parse() (map[string]ini.File, error)
}
```

系统已经内置了FileParser和ManualParser结构体，前者为默认的解析器，用于解析文件格式的配置，后者用户配置成手动注入的解析器

### ManualParser用法

```go
// 如果有如下配置想手动注入
// app.ini
// [system]
// version = beta0.1

p := config.NewManualParser()

// 方法1
p.SetConfig("app","system","version","beta0.1")
// 也可以直接快捷注入ini.Section
p.SetConfig("app","system",ini.Section{"version":"beta0.1"})

// 设置直接注入ini.File
p.SetConfig("app",ini.File{"system":ini.Section{"version":"beta0.1"}})

// 初始化config,Init的第一个Debug参数随便设置
if err := config.Init(false,p);err != nil {
	log.Error(err)
}

// 正常使用
config.String(config.Data("app").Get("sytem","version"))
```