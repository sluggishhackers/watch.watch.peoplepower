package models

import "fmt"

type Vote struct {
	ID          string `JSON:"id"`
	BillID      string `JSON:"billId"`
	Date        string `JSON:"date"`
	BillName    string `JSON:"billName"`
	PersonID    string `JSON:"personId"`
	PersonName  string `JSON:"personName"`
	PersonParty string `JSON:"personParty"`
	SessionID   string `JSON:"sessionId"`
	SessionTurn string `JSON:"sessionTurn"`
	Status      string `JSON:"status"`
	Result      string `JSON:"result"`
}

func CreateVoteID(BillID string, PersonID string) string {
	return fmt.Sprintf("%s_%s", BillID, PersonID)
}
