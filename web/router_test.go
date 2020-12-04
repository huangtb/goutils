package web

import "testing"

func TestRouter(t *testing.T) {
	router := NewRouter()
	r := router.Engine.Group("/encryption/api/v1")
	r.POST("/enc/bvn")
	router.Engine.Run(":" + "8000")
}


