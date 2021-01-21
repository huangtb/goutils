package goutils

import (
	"strconv"
	"time"
)

const (
	YYYY_mm_DDHH_mm_SS = "2006-01-02 15:04:05"
	YYYY_mm_DD         = "2006-01-02"
	YYYYmmDD           = "20060102"
)

//获取当前时间字符串（自定义格式）
func GetCurrFormatDate(layout string) string {
	return time.Now().Format(layout)
}

//获取n天前时间字符串（自定义格式）
func GetBeforeDate(n int, layout string) string {

	return time.Now().AddDate(0, 0, -n+1).Format(layout)
}

//两个日期相减，如果  d2-d1
func TimeSub(d1, d2 string) int {
	var t1, t2 time.Time
	layout := "2006-01-02"
	if d2 == "" {
		d2 = time.Now().Format(layout)
	}

	t1, _ = time.Parse(layout, d1)
	t2, _ = time.Parse(layout, d2)
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	return int(t2.Sub(t1).Hours() / 24)
}


//格式化时间戳
func FormatTimestamp(ts int64, layout string) string {
	if layout == "" {
		return ""
	}
	str := strconv.FormatInt(ts, 10)
	if len(str) == 10 {
		return time.Unix(ts, 0).Format(layout)
	}

	if len(str) == 13 {
		return time.Unix(ts/1000, 0).Format(layout)
	}
	return ""
}


//获取传入的时间所在月份的第一天
func GetFirstDateOfMonth(date string) string {

	local, _ := time.LoadLocation("Local")

	d, _ := time.ParseInLocation("2006-01-02", date, local)

	d = d.AddDate(0, 0, -d.Day() + 1)

	firstDate := d.Format("2006-01-02")

	return firstDate
}
