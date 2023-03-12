package crawler

import (
	"os"
	"reflect"
	"testing"
)

func Test_v2exCrawler_Crawl(t *testing.T) {
	type fields struct {
		PagesNum int
		ProxyUrl string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// 本地执行测试-如果不能访问外网，需要设置代理，不然网络可能不通
		{
			name: "测试获取2个页面的招聘信息",
			fields: fields{
				PagesNum: 2,
				// 为了让本地可以执行测试-github上不需要代理。为了让代码方便使用，这里使用环境变量，本地临时在环境变量里设置一下代理地址以便测试用例通过
				ProxyUrl: os.Getenv("proxyUrl"),
			},
			want: 40,
		},
		{
			name: "测试获取3个页面的招聘信息",
			fields: fields{
				PagesNum: 3,
				ProxyUrl: os.Getenv("proxyUrl"),
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crawler := &v2exCrawler{
				PagesNum: tt.fields.PagesNum,
				ProxyUrl: tt.fields.ProxyUrl,
			}
			if got := crawler.Crawl(); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("Crawl() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_crawlPage(t *testing.T) {
	type args struct {
		pageNum  int
		proxyUrl string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// 本地执行测试-如果不能访问外网，需要设置代理，不然网络可能不通
		{
			name: "测试爬取第1页",
			args: args{
				pageNum: 1,
				// 为了让本地可以执行测试-github上不需要代理。为了让代码方便使用，这里使用环境变量，本地临时在环境变量里设置一下代理地址以便测试用例通过
				proxyUrl: os.Getenv("proxyUrl"),
			},
			want:    20, // 期望结果数量
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := crawlPage(tt.args.pageNum, tt.args.proxyUrl)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("crawlPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("crawlPage() len(got) = %v, want %v", len(got), tt.want)
			}

		})
	}
}

func TestV2exItem_crawlDetailPage(t *testing.T) {
	type fields struct {
		Item          Item
		LastReplyTime string
		ReplyCount    int
	}
	type args struct {
		proxyUrl string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "测试详情页面处理",
			fields: fields{
				Item: Item{
					Title:   "",
					Content: "",
					Url:     "https://www.v2ex.com/t/921667",
				},
				LastReplyTime: "",
				ReplyCount:    0,
			},
			args: args{
				proxyUrl: os.Getenv("proxyUrl"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v2exItem := &V2exItem{
				Item:          tt.fields.Item,
				LastReplyTime: tt.fields.LastReplyTime,
				ReplyCount:    tt.fields.ReplyCount,
			}
			v2exItem.crawlDetailPage(tt.args.proxyUrl)
			if len(v2exItem.Content) == 0 {
				t.Errorf("crawlPage() len(v2exItem.Content) = %v, want > 0", len(v2exItem.Content))
			}
		})
	}
}
