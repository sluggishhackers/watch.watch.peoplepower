package crawler

import (
	"github.com/gocolly/colly"
	"github.com/sluggishhackers/watch.watch.peoplepower/store"
)

type ICrawler interface {
	FetchPeople()
	FetchSessions(withVotes bool)
	FetchTestSession()
}

type Crawler struct {
	personCrawler *colly.Collector
	store         store.IStore
}

func (c *Crawler) FetchSessions(withVotes bool) {
	if withVotes {
		SessionCrawler.OnRequest(func(r *colly.Request) {
			r.Ctx.Put("withVotes", "1")
		})
	}

	SessionsCrawler.Visit("http://watch.peoplepower21.org/Session")
	SessionCrawler.Wait()
}

// Test용으로 안건 2개가 처리된 본회의를 테스트해보자
func (c *Crawler) FetchTestSession() {
	SessionCrawler.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("withVotes", "1")
	})

	SessionCrawler.Visit("http://watch.peoplepower21.org/index.php?mid=Session&meeting_id=995")
}

func (c *Crawler) FetchPeople() {
	PeoplePaginationCrawler.Visit("http://watch.peoplepower21.org/?act=&mid=AssemblyMembers&vid=&mode=search")
}

func New(store store.IStore) ICrawler {
	InitVotesCrawler(store)
	InitPersonCrawler(store)
	InitSessionCrawler(store)
	InitSessionsCrawler(store)

	return &Crawler{
		store: store,
	}
}
