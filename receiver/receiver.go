package receiver

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Receiver 接收者
type Receiver interface {
	Receive() error
}

const JsonContentType = "application/json;charset=utf-8"

// Post 发起Post请求
func Post(url string, body interface{}) ([]byte, error) {
	jsonByte, err := json.Marshal(body)
	if err != nil {
		log.Println("转换json错误！")
		return nil, err
	}
	resp, err := http.Post(url, JsonContentType, bytes.NewReader(jsonByte))
	if err != nil {
		log.Println("http请求失败！")
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err, "关闭请求响应body时出现错误！")
		}
	}(resp.Body)
	return ioutil.ReadAll(resp.Body)
}
