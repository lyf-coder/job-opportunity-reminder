package crawler

import (
	"encoding/json"
	"testing"
)

func TestEleDuckCrawler_Crawl(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "测试爬取电鸭社区招聘信息",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EleDuckCrawler{}
			if got := e.Crawl(); len(got) > 0 {
				v, _ := json.Marshal(got)
				t.Log(string(v))
			} else {
				t.Errorf("Crawl() = %v, want len(got)>0", len(got))
			}
		})
	}
}
