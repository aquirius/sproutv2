package login

// LoginProvider provides *Login
type LoginProvider struct {
	Login *Login
}

// Login is capable of providing login access
type Login struct {
}

// NewLoginProvider returns a new Login provider
func NewLoginProvider() *LoginProvider {

	return &LoginProvider{
		&Login{},
	}
}

func (b *LoginProvider) NewLogin() *Login {
	return b.Login
}
