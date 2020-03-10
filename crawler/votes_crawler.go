package crawler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sluggishhackers/watch.watch.peoplepower/models"
	"github.com/sluggishhackers/watch.watch.peoplepower/store"
	"github.com/sluggishhackers/watch.watch.peoplepower/utils"
)

type HTTPCrawler struct {
	store store.IStore
}

func (c *HTTPCrawler) Visit(link string, BillID string, ch chan bool) {
	var meetingbill_id string
	var sessionID string
	var BillName string

	pageRes, _ := http.Get(fmt.Sprintf("http://watch.peoplepower21.org/?mid=LawInfo&bill_no=%s", BillID))
	defer pageRes.Body.Close()

	var readerBuf bytes.Buffer
	clonendReader := io.TeeReader(pageRes.Body, &readerBuf)

	pageDoc, _ := goquery.NewDocumentFromReader(clonendReader)

	pageDoc.Find(".col-xs-12 .panel-body h1").Each(func(index int, title *goquery.Selection) {
		if index == 0 {
			BillName = strings.Trim(title.Text(), " ")
		}
	})

	sessionAnchor := pageDoc.Find("#collapseThree h4 a")
	URL, _ := sessionAnchor.Attr("href")
	sessionID = utils.QueryStringValueFromRawURL(URL, "meeting_id")

	// 표결현황을 가져오려면 별도의 API에 POST 요청을 보내야 하는데
	// 이때 보내는 `meetingbill_id`값이 javascript 안에 숨어있음
	pageBody, _ := ioutil.ReadAll(&readerBuf)
	str := string(pageBody)
	targetStr := "\"&meetingbill_id=\" + "
	lastIndex := strings.LastIndex(str, targetStr)

	// meetingbill_id를 찾을 수 없는 경우는 무시
	if lastIndex == -1 {
		ch <- false
	} else {
		bill := c.store.GetBill(BillID)
		session := c.store.GetSession(sessionID)

		fmt.Println(fmt.Sprintf("크롤링 시작: \"%s\" 본회의 - \"%s\" 안건", session.Date, BillName))

		meetingbill_id = strings.Trim(strings.Trim(str[lastIndex+len(targetStr):lastIndex+len(targetStr)+7], ""), ",")

		res, _ := http.PostForm(link, url.Values{"term_no": {"20"}, "meetingbill_id": {meetingbill_id}})
		defer res.Body.Close()

		votesDoc, _ := goquery.NewDocumentFromReader(res.Body)
		votesDoc.Find("table tr").Each(func(trIndex int, tr *goquery.Selection) {
			ResultSelector := tr.Find("td:nth-of-type(1) > span")
			result := ResultSelector.Text()

			tr.Find("td:nth-of-type(2)").Each(func(_i int, td *goquery.Selection) {
				var PartyName string
				td.Find("span").Each(func(_i int, spanSelector *goquery.Selection) {
					var PersonName string

					// 의원 정보
					// 1. "href" 에서 의원 ID를 빼냄
					// 2. Text 정보에서 의원명을 빼냄
					// 3. VoteID 조합
					if spanSelector.HasClass("session_attend_name") {
						personAnchor := spanSelector.Find("a")
						PersonName = strings.Trim(personAnchor.Text(), " ")
						URL, _ := personAnchor.Attr("href")
						PersonID := utils.QueryStringValueFromRawURL(URL, "member_seq")
						VoteID := models.CreateVoteID(BillID, PersonID)

						c.store.SaveVote(&models.Vote{ID: VoteID, BillID: bill.ID, Date: session.Date, BillName: BillName, PersonID: PersonID, PersonName: PersonName, PersonParty: PartyName, Result: result, SessionID: session.ID, SessionTurn: session.Turn, Status: bill.Status})
					} else {
						// 정당명
						PartyName = strings.Trim(spanSelector.Text(), "●  ")
					}

				})
			})

			// // 누락된 의원을 찾기 위한 작업 2
			// // 누락된 의원은 해당없음으로 채우기
			// for _, Person := range c.store.GetPeople() {
			// 	if sort.SearchStrings(personIDs, Person.ID) > -1 {
			// 		VoteID := models.CreateVoteID(BillID, Person.ID)
			// 		c.store.SaveVote(&models.Vote{ID: VoteID, BillID: bill.ID, Date: session.Date, BillName: BillName, PersonID: Person.ID, PersonName: Person.Name.Ko, PersonParty: PartyName, Result: "해당없음", SessionID: session.ID, SessionTurn: session.Turn, Status: bill.Status})
			// 	}
			// }
		})

		ch <- true
	}
}

var VotesCrawler *HTTPCrawler

func InitVotesCrawler(store store.IStore) {
	VotesCrawler = &HTTPCrawler{
		store: store,
	}
}
