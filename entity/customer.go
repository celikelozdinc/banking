package entity


type Customer struct {
	Name string
	Surname string
	Accounts *[]Account
}


func NewCustomer(name string, surname string ) *Customer{
	return &Customer{Name:name, Surname:surname}
}