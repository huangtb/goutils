package web


import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	engine := gin.Default()
	engine.Use(AccessMiddleware, MonitorMiddleware)
	pprof.Register(engine)
	return &Router{Engine: engine}
}

