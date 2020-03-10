package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/sluggishhackers/watch.watch.peoplepower/store"
)

var SessionsCrawler *colly.Collector

func InitSessionsCrawler(store store.IStore) {
	SessionsCrawler = colly.NewCollector()

	SessionsCrawler.OnHTML("select#meeting", func(e *colly.HTMLElement) {
		e.ForEach("option[id]", func(_index int, session *colly.HTMLElement) {
			SessionCrawler.Visit(fmt.Sprintf("http://watch.peoplepower21.org/index.php?mid=Session&meeting_id=%s", session.Attr("value")))
		})
	})
}
