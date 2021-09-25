package accounts

import "errors"

// account struct
type Account struct {
	owner   string
	balance int
}

var errnoMoney = errors.New("Can't withdraw")

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// 이런식으로 변화하는 경우에는 * 복사본이 아니라 원본
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errnoMoney
	}
	a.balance -= amount
	return nil
}

// Change of the account
func (a *Account) ChangeOwner(newOnwers string) {
	a.owner = newOnwers
}

// Owner
func (a Account) Owner() string {
	return a.owner
}
