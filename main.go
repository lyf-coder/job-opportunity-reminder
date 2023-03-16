package main

import (
	"github.com/lyf-coder/job-opportunity-reminder/crawler"
	"github.com/lyf-coder/job-opportunity-reminder/receiver"
	"log"
	"os"
)

func main() {
	var list []interface{}
	// 爬虫数组
	crawlers := []crawler.Crawler{
		// v2ex 爬虫
		&crawler.V2exCrawler{
			PagesNum: 1,
			ProxyUrl: os.Getenv("proxy_url"),
		},
	}
	for _, c := range crawlers {
		list = append(list, c.Crawl()...)
	}

	receivers := []receiver.Receiver{
		&receiver.FeiShuReceiver{
			Url:  os.Getenv("fei_shu_bot_webhook_url"),
			Data: list,
		},
	}
	for _, r := range receivers {
		r.Receive()
	}

	log.Println("执行完成")
}
