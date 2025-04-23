package helper

import (
	"abnet_backend/source"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WTF ？？？？？？？？？这tm什么狗屎api逻辑，一个服务器掉线整个api没法访问？？？？？？？？？

type ServerListRequest struct {
	Servers []string `json:"servers"`
	Type    string   `json:"type"`
	Args    string   `json:"args"`
}

type ServerData struct {
	Servers []string `json:"servers"`
}

// GetServerList 从 Looking Glass 获取服务器列表
func GetServerList() (*ServerData, error) {
	baseURL := source.AppConfig.Server.LG_BaseURL

	// 准备请求数据
	reqData := ServerListRequest{
		Servers: []string{},
		Type:    "server_list",
		Args:    "",
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", baseURL+"/api/", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// 解析响应
	var tempResponse struct {
		Result []struct {
			Server string `json:"server"`
		} `json:"result"`
	}
	if err := json.Unmarshal(body, &tempResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// 将服务器名称添加到数组中
	serverNames := make([]string, 0, len(tempResponse.Result))
	for _, server := range tempResponse.Result {
		serverNames = append(serverNames, server.Server)
	}

	response := &ServerData{
		Servers: serverNames,
	}
	return response, nil
}

// GetServerNames 获取服务器名称列表
func GetServerNames() ([]string, error) {
	response, err := GetServerList()
	if err != nil {
		return nil, err
	}
	return response.Servers, nil
}
