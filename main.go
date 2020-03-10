package main

import (
	"github.com/sluggishhackers/watch.watch.peoplepower/crawler"
	"github.com/sluggishhackers/watch.watch.peoplepower/exporter"
	"github.com/sluggishhackers/watch.watch.peoplepower/models"
	"github.com/sluggishhackers/watch.watch.peoplepower/store"
)

func main() {
	store := store.New()
	crawlerService := crawler.New(store)

	// 의원명단 크롤링
	// crawlerService.FetchPeople()

	// 본회의 && (optional) 표결현황 크롤링
	// crawlerService.FetchSessions(true)
	crawlerService.FetchTestSession()

	// 표결정보 CSV로 추출하기
	votes := store.GetVotes()
	ExportVotes(votes)
}

func ExportVotes(votes map[string]*models.Vote) {
	exporter := exporter.New()

	rows := [][]string{
		{"idx", "vote", "bill_idx", "bill_name", "turn", "status", "date", "name_kr", "party"},
	}

	for _, v := range votes {
		rows = append(rows, []string{v.ID, v.Result, v.BillID, v.BillName, v.SessionTurn, v.Status, v.Date, v.PersonName, v.PersonParty})
	}

	exporter.CSV(rows, "20.csv")
}
