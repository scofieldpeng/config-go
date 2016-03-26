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

当配置文件存在着在其文件名上加_debug的另一个文件时，debug 模式下将会优先加载_debug文件，当不存在_debug文件时加载原文件。
如：有app.ini的文件，同时还有app_debug.ini文件，当为debug模式时会优先加载app_debug.ini。如果app_debug.ini不存在时，将加载app.ini文件

在正常模式下都只加载app.ini文件


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
