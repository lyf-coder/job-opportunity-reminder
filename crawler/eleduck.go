package crawler

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"log"
	"time"
)

// 电鸭社区 招聘

const eleDuckJobUrl = "https://eleduck.com/categories/5?sort=new"
const detailUrl = "https://eleduck.com/posts/"
const detailApiUrl = "https://svc.eleduck.com/api/v1/posts/"

type EleDuckItem struct {
	Item
	Id string
}

type EleDuckCrawler struct {
	crawler
}

// 处理详细信息
func (item *EleDuckItem) crawlDetailPage() {
	var postData postData
	dataBytes, err := Get(detailApiUrl + item.Id)
	if err != nil {
		log.Println("请求接口失败！", item, err)
	}
	err = json.Unmarshal(dataBytes, &postData)
	if err != nil {
		log.Println("接口返回值转换格式失败！", item, err)
	}
	item.Title = postData.Post.Title
	item.Content = postData.Post.RawContent
	item.Flag = eleDuck
}

func (eleCrawler *EleDuckCrawler) Crawl() []interface{} {
	var list []interface{}
	c := colly.NewCollector()

	// id 为 TopicsNode 的 div
	c.OnHTML(`script[id=__NEXT_DATA__]`, func(e *colly.HTMLElement) {
		var d queryData
		err := json.Unmarshal([]byte(e.Text), &d)
		if err != nil {
			log.Println("格式转换失败！", err)
		}
		for _, p := range d.Props.InitialProps.PageProps.PostList.Posts {
			item := &EleDuckItem{
				Item: Item{
					Url:         detailUrl + p.Id,
					PublishTime: p.PublishedAt,
				},
				Id: p.Id,
			}
			// 只抓取时间范围内的详细页面
			if eleCrawler.isInDurationSec(&item.Item) {
				item.crawlDetailPage()
				list = append(list, item)
			}
		}
	})

	err := c.Visit(eleDuckJobUrl)
	if err != nil {
		log.Println(err)
	}
	log.Println("eleduck", len(list))
	return list
}

type queryData struct {
	Props struct {
		InitialProps struct {
			PageProps struct {
				PostList struct {
					Posts []struct {
						Id               string      `json:"id"`
						Title            string      `json:"title"`
						Summary          string      `json:"summary"`
						PublishedAt      string      `json:"published_at"`
						Deleted          bool        `json:"deleted"`
						Featured         bool        `json:"featured"`
						Pinned           bool        `json:"pinned"`
						PinnedInCategory bool        `json:"pinned_in_category"`
						ModifiedAt       time.Time   `json:"modified_at"`
						TouchedAt        time.Time   `json:"touched_at"`
						DeletedAt        interface{} `json:"deleted_at"`
						LastCommentAt    *time.Time  `json:"last_comment_at"`
						ViewsCount       int         `json:"views_count"`
						CommentsCount    int         `json:"comments_count"`
						Hide             bool        `json:"hide"`
						HidedAt          interface{} `json:"hided_at"`
						UpvotesCount     int         `json:"upvotes_count"`
						DownvotesCount   int         `json:"downvotes_count"`
						MarksCount       int         `json:"marks_count"`
						Tags             []struct {
							Id       int    `json:"id"`
							Name     string `json:"name"`
							TagGroup struct {
								Id   int    `json:"id"`
								Name string `json:"name"`
								Code string `json:"code"`
							} `json:"tag_group"`
						} `json:"tags"`
						Category struct {
							Id   int    `json:"id"`
							Name string `json:"name"`
							Code string `json:"code"`
						} `json:"category"`
						User struct {
							Id                   string   `json:"id"`
							Nickname             string   `json:"nickname"`
							Tagline              *string  `json:"tagline"`
							Roles                []string `json:"roles"`
							TagList              []string `json:"tag_list"`
							TagIdentity          *string  `json:"tag_identity"`
							AvatarUrl            string   `json:"avatar_url"`
							ContentsDeletedCount int      `json:"contents_deleted_count"`
						} `json:"user"`
						Contributors []interface{} `json:"contributors"`
						LastComment  struct {
							Id   string `json:"id"`
							User struct {
								Id       string `json:"id"`
								Nickname string `json:"nickname"`
							} `json:"user"`
						} `json:"last_comment,omitempty"`
					} `json:"posts"`
				} `json:"postList"`
			} `json:"pageProps"`
			Err interface{} `json:"err"`
		} `json:"initialProps"`
	} `json:"props"`
}

type postData struct {
	Post struct {
		Id               string      `json:"id"`
		Title            string      `json:"title"`
		PublishedAt      time.Time   `json:"published_at"`
		ModifiedAt       time.Time   `json:"modified_at"`
		Featured         bool        `json:"featured"`
		Pinned           bool        `json:"pinned"`
		PinnedInCategory bool        `json:"pinned_in_category"`
		ViewsCount       int         `json:"views_count"`
		CommentsCount    int         `json:"comments_count"`
		TouchedAt        time.Time   `json:"touched_at"`
		ChargesCount     int         `json:"charges_count"`
		ChargesEle       int         `json:"charges_ele"`
		RawContent       string      `json:"raw_content"`
		Hide             bool        `json:"hide"`
		HidedAt          interface{} `json:"hided_at"`
		UpvotesCount     int         `json:"upvotes_count"`
		DownvotesCount   int         `json:"downvotes_count"`
		MarksCount       int         `json:"marks_count"`
		Deleted          bool        `json:"deleted"`
		Content          string      `json:"content"`
	} `json:"post"`
}
