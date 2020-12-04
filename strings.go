package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"time"
)

//生成id类型生成id
func GetMd5String(s string) string {
	h := md5.New()
	var md5Str string
	io.WriteString(h, s)
	md5Str = fmt.Sprintf("%X", h.Sum(nil))
	return md5Str
}


func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		//如果失败则用md5(时间戳)
		ts := time.Now().UnixNano()
		return GetMd5String(strconv.FormatInt(ts, 10))
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
