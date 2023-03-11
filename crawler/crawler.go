package crawler

// Item 数据条目结构体
type Item struct {
	// 标题
	Title string `json:"title,omitempty"`
	// 内容
	Content string `json:"content"`
	// 详细链接
	Url string `json:"url,omitempty"`
}

// Crawler 爬虫
type Crawler interface {
	// Crawl 进行爬虫
	Crawl() []interface{}
}
