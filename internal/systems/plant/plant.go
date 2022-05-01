package plant

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type PlantHandler struct{}

// LoginProvider provides *Login
type PlantProvider struct {
	Plant *Plant
}

// Login is capable of providing login access
type Plant struct {
	db sqlx.DB
}

func (plant *Plant) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		fmt.Println("hit")
		return
	case r.Method == http.MethodPost:
		return
	default:
		return
	}
}

// NewLoginProvider returns a new Login provider
func NewPlantProvider(db *sqlx.DB) *PlantProvider {
	return &PlantProvider{
		&Plant{db: *db},
	}
}

func (b *PlantProvider) NewPlant() *Plant {
	return b.Plant
}
