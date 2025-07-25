package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// PauseRequest 暂停请求结构体
type PauseRequest struct {
	AccountToken string `json:"account_token"`
	Lang         string `json:"lang"`
}

// PauseResponse 暂停响应结构体
type PauseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// HTTPClient HTTP客户端接口
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client 雷神加速器客户端
type Client struct {
	BaseURL    string
	HTTPClient HTTPClient
	Timeout    time.Duration
}

// NewClient 创建新的雷神客户端
func NewClient() *Client {
	return &Client{
		BaseURL: "https://webapi.leigod.com",
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		Timeout: 5 * time.Second,
	}
}

// Pause 暂停加速器
func (c *Client) Pause(accountToken, lang string) (*PauseResponse, error) {
	pauseReq := PauseRequest{
		AccountToken: accountToken,
		Lang:         lang,
	}

	jsonData, err := json.Marshal(pauseReq)
	if err != nil {
		return nil, fmt.Errorf("序列化请求数据失败: %v", err)
	}

	url := c.BaseURL + "/api/user/pause"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var pauseResp PauseResponse
	if err := json.Unmarshal(body, &pauseResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	return &pauseResp, nil
}
