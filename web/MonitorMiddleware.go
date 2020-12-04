package web

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/huangtb/go-utils/exception"
	"gopkg.in/afmobi-QSee/statlog.v2"
	"log"
	"net/http"
	"strconv"
	"time"
)

func MonitorMiddleware(ctx *gin.Context) {
	defer HandelPanic(ctx)
	start := time.Now()

	ctx.Next()

	duration := time.Now().Sub(start).Nanoseconds() / 1e6
	url := ctx.Request.URL.Path
	var buffer bytes.Buffer
	buffer.WriteString(url)
	//性能监控
	statlog.MultCount(buffer.String(), strconv.FormatInt(duration, 10))

	log.Printf("Request URL>>%s, requestId>>%s, 耗时>>%d [ms]", url,GetSession(ctx,RequestId), duration)

}

func HandelPanic(ctx *gin.Context) {
	if err := recover(); err != nil {
		requestId := ctx.GetHeader("Request-Id")
		if requestId == "" {
			//log.Log.Error("----", string(debug.Stack()))
		} else {
			//log.Log.Error(requestId+"----", string(debug.Stack()))
		}
		ctx.JSON(http.StatusOK, exception.NewUncatchedException())
	}
}
