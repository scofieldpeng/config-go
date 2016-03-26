package config

import "strconv"

// String 返回String类型的数据
func String(val string, b bool) string {
	if b {
		return val
	}
	return ""
}

// Int 返回整数型数据
func Int(val string, b bool) int {
	if b {
		tmp, err := strconv.Atoi(val)
		if err == nil {
			return tmp
		}
	}
	return 0
}

// Bool 返回布尔值数据
func Bool(val string, b bool) bool {
	if b {
		tmp, err := strconv.ParseBool(val)
		if err == nil {
			return tmp
		}
	}
	return false
}

// Float64 返回Float64值
func Float64(val string, b bool) float64 {
	if b {
		tmp, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return tmp
		}
	}

	return float64(0)
}
