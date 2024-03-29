package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
	"sync"
	"time"
)

// 爬取的url地址
const v2exUrl = "https://www.v2ex.com"
const jobUrl = v2exUrl + "/go/jobs?p="

// V2exItem v2ex 网站的返回的条目数据结构
type V2exItem struct {
	Item
	// 最后回复时间
	LastReplyTime string `json:"lastReplyTime,omitempty"`
	// 回复数量
	ReplyCount int `json:"replyCount,omitempty"`
}

// V2exCrawler v2ex 爬虫
type V2exCrawler struct {
	crawler
	// 查询页数
	PagesNum int
	// 代理地址 可选 示例："socks5://127.0.0.1:3128"
	ProxyUrl string
}

// 爬取详情页面，相关处理数据存储到 v2exItem 中
func (v2exItem *V2exItem) crawlDetailPage(proxyUrl string) {
	c := colly.NewCollector()
	// 处理页面-class 为 topic_content
	c.OnHTML(`div.topic_content`, func(e *colly.HTMLElement) {
		// 只取文本
		v2exItem.Content = e.Text
	})
	c.OnHTML(`div.header small.gray`, func(e *colly.HTMLElement) {
		v2exItem.PublishTime = e.ChildAttr(`span`, `title`)
	})
	// 设置代理
	if len(proxyUrl) > 0 {
		err := c.SetProxy(proxyUrl)
		if err != nil {
			fmt.Println("设置代理地址失败！", proxyUrl, err)
		}
	}

	err := c.Visit(v2exItem.Url)
	if err != nil {
		fmt.Println("访问详细页面失败！", v2exItem.Url, err)
	}
}

// crawlPage 爬取具体的页面
func (crawler *V2exCrawler) crawlPage(pageNum int) ([]interface{}, error) {
	// 保证多个协程完成执行
	var wg = sync.WaitGroup{}
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
			v2exItem := &V2exItem{
				Item: Item{
					Title: eTable.ChildText(`a.topic-link`),
					Url:   v2exUrl + eTable.ChildAttr(`a.topic-link`, "href"),
					Flag:  v2ex,
				},
				LastReplyTime: eTable.ChildAttr(`span`, "title"),
				ReplyCount:    count,
			}
			list = append(list, v2exItem)
			wg.Add(1)
			go func() {
				v2exItem.crawlDetailPage(crawler.ProxyUrl)
				wg.Done()
			}()
			// 防止同一时间的高并发请求导致被禁止访问
			time.Sleep(200 * time.Millisecond)
		})
	})
	// 设置代理
	if len(crawler.ProxyUrl) > 0 {
		err := c.SetProxy(crawler.ProxyUrl)
		if err != nil {
			return nil, err
		}
	}

	err := c.Visit(jobUrl + strconv.Itoa(pageNum))
	if err != nil {
		log.Println("error", "访问页面失败！", err)
	}
	// 等待详情信息抓取处理完成
	wg.Wait()
	log.Println("v2ex未过滤的数量", len(list))
	return crawler.filterByDurationSec(list), err
}

func (crawler *V2exCrawler) Crawl() []interface{} {
	var list []interface{}
	for i := 1; i <= crawler.PagesNum; i++ {
		pageDataList, err := crawler.crawlPage(i)
		if err != nil {
			fmt.Println("爬取页面失败", i, err)
			continue
		}
		list = append(list, pageDataList...)
	}
	log.Println("v2ex", len(list))
	return list
}
