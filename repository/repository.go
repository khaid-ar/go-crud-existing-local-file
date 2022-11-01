package repository

import (
	"encoding/json"
	"fmt"
	"go-trial/entity"
	"go-trial/util"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var dataModel entity.Model
var dataModels entity.Models

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CRUD RESTful API using GO within local file path!")
}

func Create(w http.ResponseWriter, r *http.Request) {
	util.ParseToString(&dataModels)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Data can't read")
	}
	json.Unmarshal(reqBody, &dataModel)
	dataModel.Id = strconv.FormatInt(int64((rand.Intn)(100)), 16)
	dataModels.Models = append(dataModels.Models, dataModel)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(dataModels)
	util.WriteFile(&dataModel)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	util.ParseToString(&dataModels)

	keyID := mux.Vars(r)["id"]

	for _, result := range dataModels.Models {
		if result.Id == keyID {
			json.NewEncoder(w).Encode(result)
		}
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	util.ParseToString(&dataModels)

	json.NewEncoder(w).Encode(&dataModels)

}

func Update(w http.ResponseWriter, r *http.Request) {
	keyID := mux.Vars(r)["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ID not found")
	}
	json.Unmarshal(reqBody, &dataModel)

	for i, result := range dataModels.Models {
		if result.Id == keyID {
			result.Name = dataModel.Name
			result.Status = dataModel.Status
			dataModels.Models = append(dataModels.Models[:i], result)
			json.NewEncoder(w).Encode(result)
		}
	}
	// util.WriteAll(dataModel)

}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	IDkey := mux.Vars(r)["id"]
	// util.ParseToString(&dataModels)
	for i, result := range dataModels.Models {
		if result.Id == IDkey {
			dataModels.Models = append(dataModels.Models[:i], dataModels.Models[i+1:]...)
			fmt.Fprintf(w, "Data with ID %v has been deleted successfully", IDkey)
			continue
		} else {
			util.WriteAll(result)
		}
	}
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	util.ParseToString(&dataModels)
	dataModel = entity.Model{}
	util.WriteAll(dataModel)
	fmt.Fprintf(w, "Success Truncate Data ....")
	json.NewEncoder(w).Encode(&dataModels)
}
