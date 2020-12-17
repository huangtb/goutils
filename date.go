package goutils

import "time"

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
