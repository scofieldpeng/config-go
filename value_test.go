package config

import "testing"

var(
    testValue1 string = "hello world"
    testValue2 string = "true"
    testValue3 string = "false"
    testValue4 string = "0"
    testValue5 string = "1"
    testValue6 string = "[1,2,3,4]"
    testValue7 string = "['1','2',3,\"hello world\"]"
)

func TestInt(t *testing.T) {
    if res,ok := Int(testValue1,true);ok || res != 0 {
        t.Error("testValue1 convert to int fail,get value:",res)
    }
    if res,ok := Int(testValue2,true);ok || res != 0 {
        t.Error("testValue2 convert to int false,get value:",res)
    }
    if res,ok := Int(testValue3,true);ok || res != 0 {
        t.Error("testValue3 convert to int fail,get vlaue:",res)
    }
    if res,ok := Int(testValue4,true);!ok || (ok && res != 0 ){
        t.Error("testValue4 convert to int fail,get value:",res)
    }
    if res,ok := Int(testValue5,true);!ok || (ok && res != 1 ){
        t.Error("testValue5 convert to int fail,get value:",res)
    }
}

func TestInt64(t *testing.T) {
    if res,ok := Int(testValue1,true);ok || res != 0 {
        t.Error("testValue1 convert to int64 fail,get value:",res)
    }
    if res,ok := Int(testValue2,true);ok || res != 0 {
        t.Error("testValue2 convert to int64 false,get value:",res)
    }
    if res,ok := Int(testValue3,true);ok || res != 0 {
        t.Error("testValue3 convert to int64 fail,get vlaue:",res)
    }
    if res,ok := Int(testValue4,true);!ok || (ok && res != 0 ){
        t.Error("testValue4 convert to int64 fail,get value:",res)
    }
    if res,ok := Int(testValue5,true);!ok || (ok && res != 1 ){
        t.Error("testValue5 convert to int64 fail,get value:",res)
    }
}

func TestBool(t *testing.T) {
    if res,ok := Bool(testValue1,true);ok || res {
        t.Error(testValue1," convert to bool fail,get value:",res)
    }
    if res,ok := Bool(testValue2,true);!ok || !res {
        t.Error(testValue2," convert to bool fail,get value:",res)
    }
    if res,ok := Bool(testValue3,true); !ok || res {
        t.Error(testValue3," convert to bool fail,get value:",res)
    }
    if res,ok := Bool(testValue4,true); !ok || res {
        t.Error(testValue3," convert to bool fail,get value:",res)
    }
    if res,ok := Bool(testValue5,true); !ok || !res {
        t.Error(testValue4," convert to bool fail,get value:",res)
    }
}

func TestIntSlice(t *testing.T) {
    if res,ok := IntSlice(testValue6,true);!ok || (ok && len(res) != 4 ) || (ok && (res[0] != 1 || res[1] != 2 || res[2] != 3 || res[3] != 4)) {
        t.Error(testValue6," convert to int slice fail,get value:",res)
    }
    if res,ok := IntSlice(testValue7,true);ok || (ok && len(res) == 4) {
        t.Error(testValue7," convert to int slice fail,get value",res)
    }
}

func TestInt64Slice(t *testing.T) {
    if res,ok := Int64Slice(testValue6,true);!ok || (ok && len(res) != 4 ) || (ok && (res[0] != 1 || res[1] != 2 || res[2] != 3 || res[3] != 4)) {
        t.Error(testValue6," convert to int64 slice fail,get value:",res)
    }
    if res,ok := Int64Slice(testValue7,true);ok ||(ok && len(res) == 4) {
        t.Error(testValue7," convert to int64 slice fail,get value",res)
    }
}
