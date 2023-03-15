package receiver

import (
	"encoding/json"
	"github.com/lyf-coder/job-opportunity-reminder/crawler"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFeiShuReceiver_Receive(t *testing.T) {
	// 模拟一个返回
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {

		switch strings.TrimSpace(r.URL.Path) {
		case "/":
			respBody := &respBody{
				Code: 0,
				Msg:  "成功",
				Data: struct{}{},
			}
			data, _ := json.Marshal(respBody)
			_, _ = io.WriteString(w, string(data))
		case "/test":
			respBody := &respBody{
				Code: 19021,
				Msg:  "出现错误",
				Data: struct{}{},
			}
			data, _ := json.Marshal(respBody)
			_, _ = io.WriteString(w, string(data))

		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))
	type fields struct {
		Url  string
		Data []interface{}
	}
	var data []interface{}
	data = append(data, &crawler.V2exItem{
		Item: crawler.Item{
			Title:   "[远程] 招聘 golang 工程师",
			Content: `具体详细说明xxxxxxxxxx`,
			Url:     "https://www.v2ex.com/",
		},
		LastReplyTime: "2023-03-13 11:11:11",
		ReplyCount:    0,
	})
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "正常执行测试",
			fields: fields{
				Url:  server.URL,
				Data: data,
			},
		},
		{
			name: "服务端返回错误编码测试",
			fields: fields{
				Url:  server.URL + "/test",
				Data: data,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FeiShuReceiver{
				Url:  tt.fields.Url,
				Data: tt.fields.Data,
			}
			r.Receive()
		})
	}
}
