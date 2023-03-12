package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
)

// V2exCrawler v2ex 爬虫
var V2exCrawler = &v2exCrawler{
	// 默认配置-查询 2 个页面
	PagesNum: 2,
}

// 爬取的url地址
const jobUrl = "https://www.v2ex.com/go/jobs?p="

// V2exItem v2ex 网站的返回的条目数据结构
type V2exItem struct {
	Item
	// 最后回复时间
	LastReplyTime string `json:"lastReplyTime,omitempty"`
	// 回复数量
	ReplyCount int `json:"replyCount,omitempty"`
}

// v2exCrawler v2ex 爬虫
type v2exCrawler struct {
	// 查询页数
	PagesNum int
	// 代理地址 可选 示例："socks5://127.0.0.1:3128"
	ProxyUrl string
}

// crawlPage 爬取具体的页面
func crawlPage(pageNum int, proxyUrl string) ([]interface{}, error) {
	var list []interface{}
	c := colly.NewCollector()

	// id 为 TopicsNode 的 div
	c.OnHTML(`div[id=TopicsNode]`, func(e *colly.HTMLElement) {
		// 找到 table 元素列表遍历
		e.ForEach(`table`, func(i int, eTable *colly.HTMLElement) {
			count := 0
			countStr := eTable.ChildText(`a.count_livid`)
			if len(countStr) > 0 {
				var err error
				count, err = strconv.Atoi(countStr)
				if err != nil {
					fmt.Println("转换回复数量时出错", err)
				}
			}
			list = append(list, &V2exItem{
				Item: Item{
					Title: eTable.ChildText(`a.topic-link`),
					Url:   eTable.ChildAttr(`a.topic-link`, "href"),
				},
				LastReplyTime: eTable.ChildAttr(`span`, "title"),
				ReplyCount:    count,
			})
		})
	})
	// 设置代理
	if len(proxyUrl) > 0 {
		err := c.SetProxy(proxyUrl)
		if err != nil {
			return nil, err
		}
	}

	err := c.Visit(jobUrl + strconv.Itoa(pageNum))
	return list, err
}

func (crawler *v2exCrawler) Crawl() []interface{} {
	var list []interface{}
	for i := 1; i <= crawler.PagesNum; i++ {
		pageDataList, err := crawlPage(i, crawler.ProxyUrl)
		if err != nil {
			fmt.Println("爬取页面失败", i, err)
			continue
		}
		list = append(list, pageDataList...)
	}
	return list
}
