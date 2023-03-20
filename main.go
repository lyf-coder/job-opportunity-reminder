package main

import (
	"github.com/lyf-coder/job-opportunity-reminder/config"
	"github.com/lyf-coder/job-opportunity-reminder/crawler"
	"github.com/lyf-coder/job-opportunity-reminder/receiver"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// 加载配置
	config.LoadConfig("./")
	log.Println(viper.Get("proxy_url"))
	log.Println(viper.Get("duration_sec"))
	var list []interface{}
	// 爬虫数组
	crawlers := []crawler.Crawler{
		// v2ex 爬虫
		&crawler.V2exCrawler{
			PagesNum: 1,
			ProxyUrl: viper.GetString("proxy_url"),
		},
		// eleduck 电鸭社区 爬虫
		&crawler.EleDuckCrawler{},
	}
	for _, c := range crawlers {
		list = append(list, c.Crawl()...)
	}

	receivers := []receiver.Receiver{
		&receiver.FeiShuReceiver{
			Url:  viper.GetString("fei_shu_bot_webhook_url"),
			Data: list,
		},
	}
	for _, r := range receivers {
		r.Receive()
	}

	log.Println("执行完成")
}
