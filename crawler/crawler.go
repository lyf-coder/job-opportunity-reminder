package crawler

import (
	"github.com/lyf-coder/job-opportunity-reminder/util"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Item 数据条目结构体
type Item struct {
	// 标题
	Title string `json:"title,omitempty"`
	// 内容
	Content string `json:"content"`
	// 详细链接
	Url string `json:"url,omitempty"`
	// 发布时间
	PublishTime string `json:"publishTime"`
	// 序号
	Num int
}

// Crawler 爬虫
type Crawler interface {
	// Crawl 进行爬虫
	Crawl() []interface{}
}

type crawler struct {
}

// 根据时间范围进行过滤
func (c *crawler) filterByDurationSec(list []interface{}) []interface{} {
	// 假如 duration_sec 为 0 则不进行过滤
	durationSec := viper.GetInt("duration_sec")
	if durationSec != 0 {
		// 过滤符合时间范围内的数据
		t := time.Now().In(util.CstZone).Add(-time.Duration(durationSec))
		tStr := util.GetTimeFormat(t, util.DATETIME)
		var filteredList []interface{}
		for _, listItem := range list {
			if item, ok := listItem.(Item); ok && tStr < item.PublishTime[0:19] {
				filteredList = append(filteredList, item)
			}
		}
		return filteredList
	}
	return list
}

// 在目标时间范围内
func (c *crawler) isInDurationSec(item *Item) bool {
	durationSec := viper.GetInt("duration_sec")
	if durationSec != 0 {
		return util.TimeStrInDuration(time.Duration(durationSec), item.PublishTime, util.DATETIME)
	}
	return true
}

// Get 发起 http Get 请求
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
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
