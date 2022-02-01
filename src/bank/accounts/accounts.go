package accounts

import "erros"

// Account private struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = erros.New("Can't Withdraw")

// Account의 메모리 주소를 새로운 account에 복사함
// NewAccount creates Account
func NewAccout(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// function과 메소드는 다르다
// accpimt.Deposit을 사용할수 있는 receive Method
// Deposit x amount on your account
// func (a Account) Deposit(amount int) { 를 하면 새로운 Account를 가져온다
func (a *Account) Deposit(amount int) {
	// Pointer receiver
	// * => Deposit 메소드를 호출한 Account를 사용하라는 구문
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x of youer account
func (a *Account) Whithdraw(amount int) error {
	// -10이 되기때문에 error handling을 해야한다
	if a.balance < amount {
		// return erros.New("Can't withdraw")
		return errNoMoney
	}
	// nil => null or none
	a.balance -= amount
	return nil
}
