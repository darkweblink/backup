package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitee.com/xiaoyutab/gatway/model"
)

// 实现URL接口协议，用于处理URL回调协议
type URL struct {
	model.AppRouter
}

// 超时时间
func (u *URL) Timeout() time.Duration {
	return time.Second * 60
}

// 长连接超时时间
func (u *URL) KeepAlive() time.Duration {
	return time.Second * 60
}

// TLS握手超时时间
func (u *URL) TLSHandshakeTimeout() time.Duration {
	return time.Second * 10
}

// 转发网址列表
func (u *URL) Urls() []string {
	urls := []string{}
	fmt.Println("4444444444444", u.Targets)
	if err := json.Unmarshal([]byte(u.Targets), &urls); err == nil {
		return urls
	}
	return []string{}
}

// 处理请求
func (u *URL) QuestUrl(request *http.Request) {
	fmt.Println("33333333333333")
}

// 响应处理
func (u *URL) ResponseUrl(response *http.Response) ([]byte, error) {
	fmt.Println("11111111111111")
	return nil, nil
}

// 错误处理
func (u *URL) Error(w http.ResponseWriter, r *http.Request, err error) {
	w.Write([]byte(err.Error()))
}
