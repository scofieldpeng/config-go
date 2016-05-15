package config

import (
    "testing"
    "os"
)

func TestNew(t *testing.T) {
    // load current dir test.ini file
    dirPath := os.Getenv("GOPATH") + string(os.PathSeparator) + "src" + string(os.PathSeparator) + "github.com" + string(os.PathSeparator) + "scofieldpeng" + string(os.PathSeparator) + "config-go"
    if err := New(dirPath,false);err != nil {
        t.Error("load config mode fail(node debug mode)")
    }
    if IsDebug() {
        t.Error("is debug fail,should false bug get true")
    }
    if tmp,ok := Config("test").Get("production","key");!ok || (ok && tmp != "value"){
        t.Error("get test.ini section production:key fail,need value but get ",tmp, " ok:",ok)
    }

    if err := New(dirPath,true);err != nil {
        t.Error("load config fail(debug mode)")
    }
    if !IsDebug() {
        t.Error("get IsDebug() fail,should get true but get false")
    }
    if tmp,ok := Config("test").Get("test","testKey");!ok || (ok && tmp != "testValue" ){
        t.Error("get test_debug.ini section test:testKey fail,need testValue but get ",tmp," ok:",ok)
    }
}
