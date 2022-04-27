package core

import (
	"net/http"
	"reflect"
	"regexp"

	"github.com/jmoiron/sqlx"
)

var (
	validIdentifierName = regexp.MustCompile(`^[A-Za-z0-9-]+$`)
	matchMethodName     = regexp.MustCompile(`^(.+)V([0-9]+)$`)
)

// Server ...
type Server struct {
	systems map[reflect.Type]interface{}
	dbh     *sqlx.DB
}

// NewServer returns a new Server
func NewServer(dbh *sqlx.DB) *Server {
	return &Server{
		systems: map[reflect.Type]interface{}{},
		dbh:     dbh,
	}
}

// ListenAndServe will start listening on http on the given addr
func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}
