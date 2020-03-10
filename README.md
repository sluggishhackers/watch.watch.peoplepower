watch.watch.peoplepower
===
watch.peoplepower21.org를 `watch`하는 프로젝트로 참여연대 열려라국회 사이트 크롤러입니다<br />
이 프로젝트의 기원은 코드포서울 [Congress Report](https://github.com/codeforseoul/congress-report)입니다.

Sluggish Hackers
---
- Sluggish Hackers 2번째 프로젝트로 ***x팀과 협업
- [협업 내용은 이곳에서 확인할 수 있습니다](https://github.com/sluggishhackers/collaborators/issues/2)

How To Use
---
0. 초기화
```golang
store := store.New()
crawlerService := crawler.New(store)
```

1. 국회의원 목록 수집하기
```golang
crawlerService.FetchPeople()
```

2. 본회의 목록 & 표결정보(optional) 수집하기
```golang
crawlerService.FetchSessions(true)
crawlerService.FetchSessions(false) // 표결정보 제외하고 수집하기
```

3. 수집 후 표결정보 CSV로 꺼내기
```golang
exporter := exporter.New()
votes := store.GetVotes()

// Column Headers
rows := [][]string{
   {"idx", "vote", "bill_idx", "bill_name", "turn", "status", "date", "name_kr", "party"},
}

for _, v := range votes {
   rows = append(rows, []string{v.ID, v.Result, v.BillID, v.BillName, v.SessionTurn, v.Status, v.Date, v.PersonName, v.PersonParty})
}

exporter.CSV(rows, "20.csv")

// File "data/20.csv" is created!
```

Features
---

#### 크롤링 대상
1. [x] 국회의원별 표결 정보
2. [x] 국회의원 목록 & 세부정보
3. [ ] 본회의 목록 & 세부정보
    * *표결 정보를 위해 일부만 작업되어 있음*
4. [ ] 상임위원회 목록 & 세부정보

#### Export
* 파일 유형
  * [x] CSV
  * [ ] JSON

Development
---

#### 1. Crawler

1. **국회의원 관련**
    1. `people_crawler`를 통해 페이지 수를 파악한 후 의원 목록을 수집
    2. `person_crawler`를 통해 의원별 상세 정보를 수집
2. **본회의  관련**
    1. `sessions_crawler`를 통해 현재까지의 본회의 목록을 수집
    2. `session_crawler`를 통해 본회의 상세 정보를 수집
3. **표결 관련**
    1. `session_crawler`를 통해 수집된 본회의 정보에서 본회의에서 처리된 안건 목록을 수집
    2. `votes_crawler`를 통해 안건별 표결 정보를 수집. 표결 정보는 HTML 형태가 아닌 별도의 API를 통해서 수집. 이 떄 사용되는 `meetingbill_id`는 어떤 ID를 의미하는 것인지 특정하지 못해서 `Javascript` 코드 내에서 찾아내 활용하는 중.

#### 2. Exporter
1. **CSV**

Open Sources
---
- [colly](https://github.com/gocolly/colly)
- [goquery](https://github.com/PuerkitoBio/goquery)

Credit
---
- [춘식(Sluggish Hackers 운영자)](https://github.com/the6thm0nth)
- [배여운(데이터 저널리스트)](https://github.com/the6thm0nth)
- [ftto(크루)](https://ftto.kr)
- 희진([Congress Report](https://github.com/codeforseoul/congress-report) 개발자)

LICENSE
---
[MIT](LICENSE)