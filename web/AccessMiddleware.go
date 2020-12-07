package web

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)


const RequestId = "Request-Id"


func AccessMiddleware(ctx *gin.Context) {
	requestId := ctx.Request.Header.Get(RequestId)
	ctx.Set(RequestId,requestId)
	//log.Log.Info("RequestId>>%s, Body>>%s", requestId,GetRequestBody(ctx))
}


func GetSession(ctx *gin.Context, sessionKey string) string {
	if v, ok := ctx.Get(sessionKey); ok {
		return v.(string)
	}
	return ""
}


func GetRequestBody(ctx *gin.Context) string {
	// read the response body to a variable
	bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
	bodyString := string(bodyBytes)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) //必须要有
	return bodyString
}
