package crawler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
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
			crawler := &V2exCrawler{
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
	// 模拟一个返回
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {

		switch strings.TrimSpace(r.URL.Path) {
		case "/":
			_, _ = io.WriteString(w, `
<div id="Main">
<div class="sep20"></div>
<div class="box" style="border-bottom: 0px;">
<div class="header"><div class="fr"><a href="/member/sky123488"><img src="https://cdn.v2ex.com/gravatar/b3dbce35ac6408c1a6faea38663aeba2?s=73&amp;d=retro" class="avatar" border="0" align="default" alt="sky123488"></a></div>
<a href="/">V2EX</a> <span class="chevron">&nbsp;›&nbsp;</span> <a href="/go/jobs">酷工作</a>
<div class="sep10"></div>
<h1>[Talentorg Joblist] 前端工程师/远程/全职</h1>
<div id="topic_923916_votes" class="votes">
<a href="javascript:" onclick="upVoteTopic(923916);" class="vote"><li class="fa fa-chevron-up"></li></a> &nbsp;<a href="javascript:" onclick="downVoteTopic(923916);" class="vote"><li class="fa fa-chevron-down"></li></a></div> &nbsp; <small class="gray"><a href="/member/sky123488">sky123488</a> · <span title="2023-03-14 15:12:55 +08:00">55 分钟前</span> · 86 次点击</small>
</div>
<div class="cell">
<div class="topic_content">招聘信息详细内容xxxxx</div>
</div>
</div>
<div class="sep20"></div>
<div id="no-comments-yet">
目前尚无回复
</div>
<div class="sep20"></div>
<div class="box">
<div class="inner"><i class="fa fa-tags fade"></i> <a href="/tag/远程" class="tag">远程</a><a href="/tag/talentorg" class="tag">talentorg</a><a href="/tag/佐玩" class="tag">佐玩</a><a href="/tag/工程师" class="tag">工程师</a></div>
</div>
<div class="sep20"></div>
<div class="sep20"></div>
<div style="box-sizing: border-box"><div class="wwads-cn wwads-horizontal" data-id="98" style="max-width: 100%; padding-top: 10px; margin-top: 0px; text-align: left; box-shadow: 0 2px 3px rgb(0 0 0 / 10%); border-bottom: 1px solid var(--box-border-color); background-color: var(--box-background-color); color: var(--box-foreground-color);"></div></div>
<style type="text/css">.wwads-cn { border-radius: 3px !important; } .wwads-text { color: var(--link-color) !important; }</style>
</div>
`)

		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))
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
					Url:     server.URL,
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
