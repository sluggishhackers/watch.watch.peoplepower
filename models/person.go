package models

type PersonName struct {
	Ko    string `JSON:"ko"`
	HanJa string `JSON:"hanJa"`
}

type SangIm struct {
	Link string `JSON:"link"`
	Text string `JSON:"text"`
}

type Person struct {
	ID       string     `JSON:"id"`
	Name     PersonName `JSON:"name"`
	Party    string     `JSON:"party"`
	Precinct string     `JSON:"precinct"`
	SangIm   SangIm     `JSON:"sangIm"`
}
