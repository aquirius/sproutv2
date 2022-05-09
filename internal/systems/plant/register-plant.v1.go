package plant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Plant
type RegisterPlant struct {
	UUID        string `json:"id"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

type RegisterPlantV1Params struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RegisterPlantV1Result struct {
}

//CreateCustomerV1 creates a customer object with given arguments
func (l *PlantSystem) RegisterPlantV1(p *RegisterPlantV1Params, res *RegisterPlantV1Result) error {
	fmt.Println(p)
	if p == nil {
		return fmt.Errorf("no params")
	}
	_, err := l.db.Exec(`insert into plant (uuid, name, description) values (?,?,?)`, uuid.New(), p.Name, p.Description)
	if err != nil {
		log.Print("error: ", err)
		return err
	}
	return nil
}

func (l *PlantSystem) GetRegisterPlantHandler(w http.ResponseWriter, r *http.Request) {
	req := &RegisterPlantV1Params{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, req)
	fmt.Println(req)
	if err != nil {
		log.Fatal("error in json")
		return
	}
	err = l.RegisterPlantV1(req, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonBytes, err := json.Marshal("error")
		if err != nil {
			log.Fatal("error in json")
			return
		}
		w.Write(jsonBytes)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
