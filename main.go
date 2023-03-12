package main

import (
	"encoding/json"
	"fmt"
	"github.com/lyf-coder/job-opportunity-reminder/crawler"
)

func main() {
	var list []interface{}
	// 爬虫数组
	crawlers := []crawler.Crawler{
		// v2ex 爬虫
		crawler.V2exCrawler,
	}
	for _, c := range crawlers {
		list = append(list, c.Crawl())
	}
	out, _ := json.Marshal(list)
	fmt.Println(string(out))
}
