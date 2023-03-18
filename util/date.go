package util

import (
	"time"
)

type FORMAT string

const (
	DATE      FORMAT = "2006-01-02"
	DATETIME  FORMAT = "2006-01-02 15:04:05"
	TIMESTAMP FORMAT = "2006-01-02 15:04:05.000"
)

// CstZone 东八区时区 使用方式：time.Now().In(CstZone)
var CstZone = time.FixedZone("CST", 8*3600)

// GetCurrent 获取 format 格式 的当前时间字符串
func GetCurrent(format FORMAT) string {
	return GetTimeFormat(time.Now(), format)
}

// GetTimeFormat 获取处理时间的format 字符串
func GetTimeFormat(handleTime time.Time, format FORMAT) string {
	return handleTime.Format(string(format))
}

// GetTime 获取字符串对应的时间
func GetTime(timeStr string, timeFormat FORMAT) (time.Time, error) {
	return time.ParseInLocation(string(timeFormat), timeStr, time.Local)
}

// TimeStrInDuration 判断时间字符串在 东八的当前时间 Duration 范围内 如 五分钟内的时间 5 * time.Minute
func TimeStrInDuration(d time.Duration, timeStr string, format FORMAT) bool {
	t := time.Now().In(CstZone).Add(-d)
	tStr := GetTimeFormat(t, format)
	return tStr < timeStr
}
