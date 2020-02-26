package main

import (
	entities "banking/entity"

)

func main(){

	account := entities.NewAccount(50)
	account.PrintAmount()

}