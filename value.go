package config

import (
	"github.com/pquerna/ffjson/ffjson"
	"strconv"
)

// Int 将config值转化为int类型
func Int(value string, ok bool) int {
	if !ok {
		return 0
	}
	res, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return res
}

// Int64 将config值转化为int64类型
func Int64(value string, ok bool) int64 {
	if !ok {
		return 0
	}
	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}

	return res
}

// Bool 将config值转化为bool类型
func Bool(value string, ok bool) bool {
	if !ok {
		return false
	}
	res, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return res
}

// IntSlice 将config值转化为int的slice类型
func IntSlice(value string, ok bool) []int {
	if !ok {
		return []int{}
	}
	var res []int
	if err := ffjson.Unmarshal([]byte(value), &res); err != nil {
		return []int{}
	}

	return res
}

// Int64Slice 将config值转化为int的slice类型
func Int64Slice(value string, ok bool) []int64 {
	if !ok {
		return []int64{}
	}
	var res []int64
	if err := ffjson.Unmarshal([]byte(value), &res); err != nil {
		return []int64{}
	}

	return res
}
