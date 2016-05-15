package config

import (
    "strconv"
    "github.com/pquerna/ffjson/ffjson"
)

// Int 将config值转化为int类型,如果转化失败,第二个参数为false
func Int(value string,ok bool)(int,bool) {
    if !ok {
        return 0,false
    }
    res,err := strconv.Atoi(value)
    if err != nil {
        return 0,false
    }
    return res,true
}

// Int64 将config值转化为int64类型,如果转化失败,第二个参数为false
func Int64(value string,ok bool)(int64,bool) {
    if !ok {
        return 0,false
    }
    res,err := strconv.ParseInt(value,10,64)
    if err != nil {
        return 0,false
    }

    return res,true
}

// Bool 将config值转化为bool类型,如果转化失败,第二个参数为false
func Bool(value string,ok bool)(bool,bool) {
    if !ok {
        return false,false
    }
    res,err := strconv.ParseBool(value)
    if err != nil {
        return false,false
    }

    return res,true
}

// IntSlice 将config值转化为int的slice类型,如果转化失败,第二个参数为false
func IntSlice(value string,ok bool)([]int,bool) {
    if !ok {
        return []int{},false
    }
    var res []int
    if err := ffjson.Unmarshal([]byte(value),&res);err != nil {
        return []int{},false
    }

    return res,true
}

// Int64Slice 将config值转化为int的slice类型,如果转化失败,第二个参数为false
func Int64Slice(value string,ok bool) ([]int64,bool) {
    if !ok {
        return []int64{},false
    }
    var res []int64
    if err := ffjson.Unmarshal([]byte(value),&res);err != nil {
        return []int64{},false
    }

    return res,true
}
