package handles

import (
	"abnet_backend/source/helper"
	"abnet_backend/source/logger"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIPInfo 获取请求者IP的信息
func GetIPInfo(c *gin.Context) {
	// 初始化数据库
	if err := helper.EnsureIPInfoDBInitialized(); err != nil {
		SendResponse(c, http.StatusInternalServerError, "ipinfo 数据库未初始化", nil)
		return
	}

	var ipAddress string
	// 尝试从查询参数获取IP
	ipParam := c.Query("ip")
	if ipParam != "" {
		ipAddress = ipParam
	} else {
		// 未提供IP参数，使用客户端IP
		ipAddress = c.ClientIP()
	}

	// 检查是否为有效IP
	parsedIP := net.ParseIP(ipAddress)
	if parsedIP == nil {
		SendResponse(c, http.StatusBadRequest, "无效的IP地址", nil)
		return
	}

	// 从 MMDB 获取 IP 信息
	ipInfo, err := helper.GetIPInfoFromMMDB(parsedIP)
	if err != nil {
		logger.Error("获取IP信息失败: %v", err)
		SendResponse(c, http.StatusInternalServerError, "获取IP信息失败", nil)
		return
	}

	// 返回结果
	SendResponse(c, http.StatusOK, "获取IP信息成功", ipInfo)
}
