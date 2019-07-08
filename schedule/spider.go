package schedule

import (
	"fmt"
	"github.com/gocolly/colly"
)

// 爬取起点小说的书籍和作者的信息
// 代理池
// ip过滤
// 数据存储
// 多个项目并行
// 定时运行
// web监控
func InitSpider() {
	c := colly.NewCollector()
	c.AllowedDomains = []string{}
	c.AllowURLRevisit = true
	c.Async = true
	c.CacheDir = ""
	c.CheckHead = true
	c.DisallowedDomains = []string{}
	c.ID = 1
	c.SetProxyFunc(nil)
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("request")
		fmt.Println("Visiting", request.URL)
	})
	c.OnError(func(response *colly.Response, e error) {
		fmt.Println("error")
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("response")
	})
	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		fmt.Println("html")
	})
	c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
}
