# config-go

a simple ini configurator for golang projects

## Install

`go get github.com/scofield/config-go`

## Usage

```go
// when we want to use debug mode,set the second parameter as true,the package will only read those xxx_debug.ini files
if err := config.New(configDir,true);err != nil {
    log.Fataln("init configurator fail!,error information:",err)
}

// assume we have the test_debug.ini file and the content is below:
//
// [test]
// key=value
val,ok := config.Config("test").Get("test","key") // val == value and ok == true

// if we want to assign the value to Int,Int64,Bool,IntSlice or Int64Slice,you can use as these way
// 
// assume we have the test_debug.ini file and the content is below:
//
// [test]
// intKey=1
// boolKey=true
// intSlice=[1,2,3]

intVal,ok := config.Int(config.Config("test").Get("test","intKey")) // intValue==1 and ok == true
boolVal,ok := config.Bool(config.Config("test").Get("test","boolKey") // boolValue == true and ok == true
intSlice,ok := config.IntSlice(config.Config("test").Get("test","intSlice") // intSlice == [1,2,3] and ok == true
```

## Licence

MIT Licence

## Thanks

Thank to thses golang package

[github.com/vaughan0/go-ini](github.com/vaughan0/go-ini)
[github.com/pquerna/ffjson/ffjson](github.com/pquerna/ffjson/ffjson)
