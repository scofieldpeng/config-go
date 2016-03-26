# config

通用配置组件，可用于ini配置文件的管理

组件会自动监听配置文件的变化，并更新配置数据

## 配置文件存储目录

配置文件应存储在程序运行的根目录下的`config`目录中。

## 初始化

在主程序是使用以下方式进行初始化

> config.Config.Init()

## Debug模式

以下语句开启debug模式
> config.SetDebug(ture)

## 获取配置

如要获取`app.ini`中的mysql节点下的default配置项

> string, err := config.Config["app"].Get("mysql", "default")

## 直接处理类型值

字符
> string := config.String(config.Config["app"].Get("mysql", "default"))

整数
> int := config.Int(config.Config["app"].Get("mysql", "default"))

浮点数
> float64 := config.Float64(config.Config["app"].Get("mysql", "default"))

布尔值
> bool := config.bool(config.Config["app"].Get("mysql", "default"))
