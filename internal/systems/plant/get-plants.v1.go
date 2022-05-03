package plant

import (
	"encoding/json"
	"log"
	"net/http"
)

// User
type Plant struct {
	UUID        string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetPlantsV1Params struct {
}

type GetPlantsV1Result struct {
	Plants []Plant `json:"plants"`
}

func (l *PlantSystem) GetPlantHandler(w http.ResponseWriter, r *http.Request) {
	var mockupPlants []Plant

	mockupPlants = append(mockupPlants,
		Plant{UUID: "testuuid", Name: "tomato", Description: "yeah good tomato"},
		Plant{UUID: "testuuid", Name: "pepper", Description: "yeah good peppers"},
		Plant{UUID: "testuuid", Name: "potato", Description: "yeah good potatoes"},
		Plant{UUID: "testuuid", Name: "cabbage", Description: "yeah good cabbages"},
		Plant{UUID: "testuuid", Name: "carrots", Description: "yeah good carrots"})
	res := &GetPlantsV1Result{Plants: mockupPlants}

	w.WriteHeader(http.StatusOK)
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		log.Fatal("error in json")
		return
	}
	w.Write(jsonBytes)
}
