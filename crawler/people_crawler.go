package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

var PeoplePaginationCrawler *colly.Collector
var PeopleCrawler *colly.Collector

func init() {
	PeoplePaginationCrawler = colly.NewCollector()
	PeopleCrawler = colly.NewCollector()

	PeoplePaginationCrawler.OnHTML("ul.pagination", func(e *colly.HTMLElement) {
		pagesCount := e.DOM.Children().Length()

		for i := 0; i < pagesCount; i++ {
			fmt.Printf("의원명단 크롤링 시작: %d페이지\n", i+1)
			PeopleCrawler.Visit(fmt.Sprintf("http://watch.peoplepower21.org/?act=&mid=AssemblyMembers&vid=&mode=search&page=%d", (i + 1)))
		}
	})

	PeopleCrawler.OnHTML(".col-md-8 > .col-xs-6 a[href]", func(e *colly.HTMLElement) {
		PersonCrawler.Visit(fmt.Sprintf("http://watch.peoplepower21.org%s", e.Attr("href")))
	})
}
