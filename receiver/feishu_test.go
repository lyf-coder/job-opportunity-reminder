package receiver

import (
	"encoding/json"
	"github.com/lyf-coder/job-opportunity-reminder/crawler"
	"github.com/lyf-coder/job-opportunity-reminder/receiver/tpl"
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
			t.Log(r.Body)
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
	// 由于路径问题-提前加载模版
	_, _ = tpl.Load("job_card_msg.json", "../tpl/feishu/job_card_msg.json")
	type fields struct {
		Url  string
		Data []interface{}
	}
	var data []interface{}
	data = append(data, &crawler.V2exItem{
		Item: crawler.Item{
			Title: "[远程] 招聘 golang 工程师",
			Content: `[任职要求]
- 计算机、电子等相关专业，全日制本科或以上学历。3 年以上后端开发经验
- 熟练掌握 golang 编程语言，代码功底扎实（支持 C++转）
- 熟练掌握后端架构中各种常见组件，框架，工具的特点，并能灵活组合应用，合理取舍来实现开发目标
- 熟悉敏捷开发，逻辑清晰，沟通能力强，注重团队合作；自我管理能力 /责任心强
- 英语读写熟练，愿意练习英语听说能力
`,
			Url:         "https://www.v2ex.com/",
			PublishTime: "2023-03-13 11:11:11",
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
