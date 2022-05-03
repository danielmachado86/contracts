package data

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}

type Party struct {
	user     *User
	contract *Contract
}

func NewParty(user *User, contract *Contract) *Party {
	return &Party{user: user, contract: contract}
}

var user1 = NewUser("Daniel M")
var user2 = NewUser("Jimena L")

var Party1 = NewParty(user1, ContractInst)
var Party2 = NewParty(user2, ContractInst)
