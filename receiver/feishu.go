package receiver

import (
	"encoding/json"
	"errors"
	"github.com/lyf-coder/job-opportunity-reminder/crawler"
	"github.com/lyf-coder/job-opportunity-reminder/receiver/tpl"
	"github.com/lyf-coder/job-opportunity-reminder/util"
	"log"
	"time"
)

// FeiShuReceiver  飞书webhook作为接受者
type FeiShuReceiver struct {
	Url  string
	Data []interface{}
}

// Receive 多并发会导致飞书机器人接收失败报错：{"code":9499,"msg":"too many request","data":{}} 所以不用协程
func (r *FeiShuReceiver) Receive() {
	count := 0
	// 东八的当前时间  -5分钟 小于发布时间即在五分钟之内
	c := time.Now().In(util.CstZone).Add(-5 * time.Minute)
	for i, itemData := range r.Data {
		if item, ok := itemData.(*crawler.V2exItem); ok {
			// 处理日期 2023-03-14 17:52:13 +08:00
			if len(item.PublishTime) >= 19 {
				t, err := util.GetTime(item.PublishTime[0:19], util.DATETIME)
				if err != nil {
					log.Println("发布日期格式处理错误！", err)
				}
				// 五分钟内
				log.Println(i, util.GetTimeFormat(c, util.DATETIME), util.GetTimeFormat(t, util.DATETIME))
				if c.Before(t) {
					count++
					item.Num = i + 1
					// 内容需要处理一下-主要是抓取的数据内容格式不转换会导致消息发送失败
					b, _ := json.Marshal(item.Content)
					item.Content = string(b)
					msg := tpl.GetTemplateResultStr("job_card_msg.json", tpl.GetTplPath("feishu/job_card_msg.json"), item)
					err := r.eachPost(msg)
					if err != nil {
						log.Println(msg, err)
					}
				}
			}

		}
	}
	log.Println("发送条数：", count)
}

// 单条飞书消息发送
func (r *FeiShuReceiver) eachPost(msg interface{}) error {
	respData, err := Post(r.Url, msg)
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

// 发送给飞书的请求响应体
type respBody struct {
	// 0 为成功
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
	} `json:"data"`
}
