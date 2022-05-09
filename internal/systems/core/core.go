package core

import (
	"github.com/jmoiron/sqlx"
)

// CoreProvider provides *Core
type CoreProvider struct {
	Core *CoreSystem
}

// Core is capable of providing core access
type CoreSystem struct {
	dbh sqlx.DB
}

// NewCoreProvider returns a new Core provider
func NewCoreProvider(dbh *sqlx.DB, urlPrefixBackend string) *CoreProvider {
	return &CoreProvider{
		&CoreSystem{
			dbh: *dbh,
		},
	}
}

func (b *CoreProvider) NewCore() *CoreSystem {
	return b.Core
}
