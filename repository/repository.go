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
var parseData entity.Models

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func Create(w http.ResponseWriter, r *http.Request) {
	util.ParseToString(&parseData)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Data can't read")
	}
	json.Unmarshal(reqBody, &dataModel)
	dataModel.Id = strconv.FormatInt(int64((rand.Intn)(100)), 16)
	parseData.Models = append(parseData.Models, dataModel)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(parseData)
	util.WriteFile(&dataModel)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	util.ParseToString(&parseData)

	keyID := mux.Vars(r)["id"]

	for _, result := range parseData.Models {
		if result.Id == keyID {
			json.NewEncoder(w).Encode(result)
		}
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	util.ParseToString(&parseData)

	json.NewEncoder(w).Encode(parseData)

}

func Update(w http.ResponseWriter, r *http.Request) {
	keyID := mux.Vars(r)["id"]
	var updateData entity.Model

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ID not found")
	}
	json.Unmarshal(reqBody, &updateData)

	for i, result := range parseData.Models {
		if result.Id == keyID {
			result.Id = updateData.Id
			result.Name = updateData.Name
			parseData.Models = append(parseData.Models[:i], result)
			json.NewEncoder(w).Encode(result)
		}
	}
	util.WriteFile(&dataModel)

}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	IDkey := mux.Vars(r)["id"]

	for i, result := range parseData.Models {
		if result.Id == IDkey {
			parseData.Models = append(parseData.Models[:i], parseData.Models[i+1:]...)
			fmt.Fprintf(w, "Data with ID %v has been deleted successfully", IDkey)
		}
	}

}
