package util

import (
	"github.com/pkg/errors"
	"strings"
	"time"
)

const (
	DateFormat     = "2006-01-02 15:04:05"
	DateTimeFormat = "20060102150405"
)

//时间转换 将1993-12-26 10:30:00转换为time
func ParseTimeByTimeStr(str, errPrefix string) (time.Time, error) {
	p := strings.TrimSpace(str)
	if p == "" {
		return time.Time{}, errors.Errorf("%s不能为空", errPrefix)
	}

	t, err := time.ParseInLocation(DateFormat, str, time.Local)
	if err != nil {
		return time.Time{}, errors.Errorf("%s格式错误", errPrefix)
	}

	return t, nil
}

//获取int64 当前时间戳/输入time时间戳
func ParseTimeToInt64(t ...time.Time) int64 {
	if len(t) == 0 {
		return time.Now().UnixNano() / 1e6
	} else {
		return t[0].UnixNano() / 1e6
	}
}

//获取int64 秒
func ParseSecondTimeToInt64() int64 {
	return time.Now().Unix()
}

//获取int64 小时
func ParseHourTimeToInt64() int64 {
	return time.Now().Unix() / 3600 * 3600
}

//获取最近的周一
func ParseCurrentMonday(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

//返回某一天的当地时区0点
func ParseMorningTime(t time.Time) time.Time {
	s := t.Format("19931226")
	result, _ := time.ParseInLocation("19931226", s, time.Local)
	return result
}

//当月第一天0点
func ParseFirstDayOfMonthMorning(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

//获取传入时间前一天的时间，不传默认是昨天
func ParseYesterdayTime(t ...time.Time) time.Time {
	if len(t) == 0 {
		return time.Now().AddDate(0, 0, -1)
	} else {
		return t[0].AddDate(0, 0, -1)
	}
}

//把int64转换成1993-12-26 10:30:00
func ParseTimeToTimeStr(intTime int64, strfmt ...string) string {
	t := time.Unix(intTime/1e3, 0)
	defaultFmt := DateFormat
	if len(strfmt) > 0 {
		defaultFmt = strfmt[0]
	}
	return t.Format(defaultFmt)
}

//int64 to time
func Int64ConvertToTime(intTime int64) time.Time {
	return time.Unix(intTime/1e3, 0)
}
