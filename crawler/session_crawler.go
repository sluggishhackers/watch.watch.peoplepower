package crawler

import (
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/sluggishhackers/watch.watch.peoplepower/models"
	"github.com/sluggishhackers/watch.watch.peoplepower/store"
	"github.com/sluggishhackers/watch.watch.peoplepower/utils"
)

var SessionCrawler *colly.Collector

func InitSessionCrawler(store store.IStore) {
	SessionCrawler = colly.NewCollector()
	// when visiting links which domains' matches "*httpbin.*" glob
	SessionCrawler.Limit(&colly.LimitRule{
		Parallelism: 5,
		RandomDelay: 5 * time.Second,
	})

	SessionCrawler.OnRequest(func(r *colly.Request) {
		id := utils.QueryStringValueFromURL(r.URL, "meeting_id")
		r.Ctx.Put("meetingId", id)
	})

	SessionCrawler.OnHTML("select > option[selected]", func(e *colly.HTMLElement) {
		title := strings.Split(e.Text, "  ")
		e.Request.Ctx.Put("date", title[0])
		e.Request.Ctx.Put("turn", title[1])
	})

	SessionCrawler.OnHTML("#collapseTwo .panel-body", func(e *colly.HTMLElement) {
		meetingID := e.Request.Ctx.Get("meetingId")
		date := e.Request.Ctx.Get("date")
		turn := e.Request.Ctx.Get("turn")

		store.SaveSession(&models.Session{
			ID:   meetingID,
			Date: date,
			Turn: turn,
		})

		if e.Request.Ctx.Get("withVotes") == "1" {
			count := e.ChildText("h2 > span")
			switch count {
			case "0":
				// 안건이 0개인 본회의는 Skip
				break
			default:
				// TODO: 어떻게 Timer를 돌리지?
				c := make(chan bool)
				len := e.DOM.Find("table tbody tr").Length()
				bills := make([]*models.Bill, len)

				e.ForEach("table tbody tr", func(index int, tr *colly.HTMLElement) {
					link := tr.ChildAttr("td:nth-of-type(1) > a", "href")
					billID := utils.QueryStringValueFromRawURL(link, "bill_no")
					Status := tr.ChildText("td:nth-of-type(3)")

					bill := &models.Bill{ID: billID, Status: Status}
					store.SaveBill(bill)
					bills[index] = bill
				})

				go VotesCrawler.Visit("http://watch.peoplepower21.org/opages/Lawinfo/_vote_table.php", bills[0].ID, c)

				docs := make([]bool, len)
				for i, _ := range docs {
					docs[i] = <-c

					if i < len-1 {
						go VotesCrawler.Visit("http://watch.peoplepower21.org/opages/Lawinfo/_vote_table.php", bills[i+1].ID, c)
					}
				}
				break
			}
		}
	})
}
