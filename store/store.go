package store

import (
	"github.com/sluggishhackers/watch.watch.peoplepower/models"
)

type IStore interface {
	GetBill(ID string) *models.Bill
	GetBills() map[string]*models.Bill
	GetPerson(ID string) *models.Person
	GetPeople() map[string]*models.Person
	GetSession(ID string) *models.Session
	GetSessions() map[string]*models.Session
	GetVote(ID string) *models.Vote
	GetVotes() map[string]*models.Vote
	SaveBill(b *models.Bill)
	SavePerson(p *models.Person)
	SaveSession(s *models.Session)
	SaveVote(v *models.Vote)
}

type Store struct {
	bills    map[string]*models.Bill
	people   map[string]*models.Person
	sessions map[string]*models.Session
	votes    map[string]*models.Vote
}

func (s *Store) GetBill(ID string) *models.Bill {
	return s.bills[ID]
}

func (s *Store) GetBills() map[string]*models.Bill {
	return s.bills
}

func (s *Store) GetPerson(ID string) *models.Person {
	return s.people[ID]
}

func (s *Store) GetPeople() map[string]*models.Person {
	return s.people
}

func (s *Store) GetSession(ID string) *models.Session {
	return s.sessions[ID]
}

func (s *Store) GetSessions() map[string]*models.Session {
	return s.sessions
}

func (s *Store) GetVote(ID string) *models.Vote {
	return s.votes[ID]
}

func (s *Store) GetVotes() map[string]*models.Vote {
	return s.votes
}

func (s *Store) SaveBill(b *models.Bill) {
	s.bills[b.ID] = b
}

func (s *Store) SavePerson(p *models.Person) {
	s.people[p.ID] = p
}

func (s *Store) SaveSession(session *models.Session) {
	s.sessions[session.ID] = session
}

func (s *Store) SaveVote(v *models.Vote) {
	s.votes[v.ID] = v
}

func New() IStore {
	return &Store{
		bills:    make(map[string]*models.Bill),
		people:   make(map[string]*models.Person),
		sessions: make(map[string]*models.Session),
		votes:    make(map[string]*models.Vote),
	}
}
