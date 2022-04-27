package login

import "fmt"

// LoginProvider provides *Login
type LoginProvider struct {
	Login *Login
}

// Login is capable of providing login access
type Login struct {
	customer *Customer
}

// NewLoginProvider returns a new Login provider
func NewLoginProvider() *LoginProvider {
	customer := CreateCustomerV1()
	fmt.Println(customer)
	return &LoginProvider{
		&Login{customer: customer},
	}
}

func (b *LoginProvider) NewLogin() *Login {
	return b.Login
}
