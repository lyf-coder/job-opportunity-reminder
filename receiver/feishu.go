package receiver

import (
	"encoding/json"
	"errors"
	"github.com/lyf-coder/job-opportunity-reminder/crawler"
	"log"
	"strings"
)

// FeiShuReceiver  飞书webhook作为接受者
type FeiShuReceiver struct {
	Url  string
	Data []interface{}
}

func (r *FeiShuReceiver) Receive() error {
	var contentArr []string

	for _, itemData := range r.Data {
		item, ok := itemData.(*crawler.V2exItem)
		if ok {
			contentArr = append(
				contentArr, item.Title, "\n\n",
				item.Content, "\n\n",
				"原始链接：", item.Url, "\n")
		}
	}

	respData, err := Post(r.Url, &textMsgBody{
		MsgType: text,
		Content: struct {
			Text string `json:"text"`
		}{
			Text: strings.Join(contentArr, ""),
		},
	})
	if err != nil {
		log.Println(`发送飞书消息失败！`)
		return err
	}

	var resp respBody
	err = json.Unmarshal(respData, &resp)
	if err != nil {
		log.Println(`转换飞书响应值失败！`, string(respData))
	}
	if resp.Code != 0 {
		log.Println(`飞书响应错误编码！`, string(respData))
		return errors.New(resp.Msg)
	}
	return nil
}

// MsgType 飞书消息类型
// 参考 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content#c9e08671
// https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json#45e0953e
type MsgType string

const (
	// text 文本消息类型
	text MsgType = "text"
	// post 富文本
	post MsgType = "post"
	// 消息卡片
	interactive MsgType = "interactive"
)

// 文本消息结构
type textMsgBody struct {
	MsgType MsgType `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

// 发送给飞书的请求响应体
type respBody struct {
	// 0 为成功
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
	} `json:"data"`
}
