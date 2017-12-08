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
// 路径必须为绝对路径，并且以/结尾
absolutePath := `/home/namer/app/config/`
config.Init( debug, NewFileParser(absolutePath))
```

## 具体 API

详见[godoc.org/github.com/zhuziweb/config](godoc.org/github.com/zhuziweb/config)