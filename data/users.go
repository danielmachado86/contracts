package data

import "time"

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}

type Party struct {
	user *User
}

func NewParty(user *User) *Party {
	return &Party{user: user}
}

var Signatures = map[*Party]*Signature{}

type Signature struct {
	Party *Party
	Date  time.Time
}

func NewSignature(party *Party) {
	Signatures[party] = &Signature{Party: party, Date: time.Now()}
}

func GetSignature(party *Party) *Signature {
	return Signatures[party]
}
