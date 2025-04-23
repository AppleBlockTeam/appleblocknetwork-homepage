package handles

import (
	"abnet_backend/source/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetServers 获取所有Looking Glass服务器列表
func GetAllServers(c *gin.Context) {
	response, err := helper.GetServerList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	SendResponse(c, http.StatusOK, "获取服务器列表成功", gin.H{
		"data": response.Servers,
	})
}
