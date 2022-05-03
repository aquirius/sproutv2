package server

import "github.com/jmoiron/sqlx"

// LoginProvider provides *Login
type ServerProvider struct {
	Server *Server
}

// Login is capable of providing login access
type Server struct {
	Sql sqlx.DB
}

// NewLoginProvider returns a new Login provider
func NewServerProvider() *ServerProvider {

	return &ServerProvider{
		&Server{},
	}
}

func (b *ServerProvider) NewServer() *Server {
	b.Server.Sql = *b.Server.connectSQL()
	return b.Server
}
