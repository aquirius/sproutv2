package plant

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type PlantHandler struct{}

// LoginProvider provides *Login
type PlantProvider struct {
	PlantSystem *PlantSystem
}

// Login is capable of providing login access
type PlantSystem struct {
	db sqlx.DB
}

func (plant *PlantSystem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	switch {
	case r.Method == http.MethodPost:
		fmt.Println("hit")
		plant.GetPlantHandler(w, r)
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
		&PlantSystem{db: *db},
	}
}

func (b *PlantProvider) NewPlant() *PlantSystem {
	return b.PlantSystem
}
