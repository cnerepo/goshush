package goshush

// Account represents a user account.
type Account struct {
	Name     string
	Password string
}

// NewAccount creates a user account with a Name and Password.
func NewAccount(name, password string) Account {
	return Account{name, password}
}
