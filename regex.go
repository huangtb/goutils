package goutils

import (
	"regexp"
)

func IsMatchRegex(str, regex string) bool {
	r, _ := regexp.Compile(regex)
	return r.MatchString(str)
}

func FindString(str, regex string) string {
	r := regexp.MustCompile("(?i)" + regex)
	return r.FindString(str)
}

func FindAllString(str, regex string) []string {
	r := regexp.MustCompile("(?i)" + regex)
	return r.FindAllString(str, -1)
}

func FindStringSubmatch(str, regex string) []string {
	r := regexp.MustCompile("(?i)" + regex)
	return r.FindStringSubmatch(str)
}

//判断字符串是否含有字母
func IsContainNonNumeric(str string) bool {
	match, err := regexp.MatchString("\\D", str)
	if err != nil {
		return false
	}
	return match

}
