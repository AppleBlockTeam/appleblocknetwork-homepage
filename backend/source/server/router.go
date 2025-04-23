package server

import (
	"abnet_backend/source/server/handles"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://ys.mihoyo.com")
	})

	// 添加简化后的API路由
	api := r.Group("/api/v0")
	{
		// 获取 ASN
		api.GET("/ipinfo", handles.GetIPInfo)
		// 获取服务器列表
		bird := api.Group("/bird")
		{
			bird.GET("/allservers", handles.GetAllServers)
		}
	}
	return r
}
