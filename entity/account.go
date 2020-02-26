package entity

import "fmt"

type Account struct{
	Amount int
}

//constructor
func NewAccount(amount int) *Account{
	return &Account{Amount:amount}
}

func (account *Account) PrintAmount() {
	fmt.Printf("Total amount of the account is : %d\n",account.Amount)
}