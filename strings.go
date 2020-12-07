package goutils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

var num2char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NumToBHex(num, n int64) string {
	var numStr string
	for num != 0 {
		yu := num % n
		numStr = string(num2char[yu]) + numStr
		num = num / n
	}
	return numStr
}

//只取前8位的UUID
func uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x", b[0:4]) //只取前8位
	return uuid
}

func GenerateEncode(prefix string) string  {
	return prefix+NumToBHex(time.Now().UnixNano(), 62) + uuid()
}


//生成id类型生成id
func GetMd5String(s string) string {
	h := md5.New()
	var md5Str string
	io.WriteString(h, s)
	md5Str = fmt.Sprintf("%X", h.Sum(nil))
	return md5Str
}

func Md5UUID() string {
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		//如果失败则用md5(时间戳)
		ts := time.Now().UnixNano()
		return GetMd5String(strconv.FormatInt(ts, 10))
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func UUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", b)
}



func main()  {
	log.Printf("UUID>%s",Md5UUID())
	log.Printf("UUID2>%s",UUID())

	log.Printf("uuid>%s",uuid())

}
