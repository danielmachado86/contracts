package data

import (
	"time"
)

type Signature struct {
	Party *Party
	Date  time.Time
}

func NewSignature(party *Party, date time.Time) *Signature {
	return &Signature{Party: party, Date: date}
}

func GetSignature(party *Party) *Signature {
	return Signatures[party]
}

var Signatures = map[*Party]*Signature{
	Party1: NewSignature(Party1, time.Date(2022, 4, 26, 15, 0, 0, 0, time.UTC)),
	Party2: NewSignature(Party2, time.Date(2022, 4, 27, 15, 0, 0, 0, time.UTC)),
}
