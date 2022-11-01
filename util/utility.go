package util

import (
	"encoding/json"
	"fmt"
	"go-trial/entity"
	"io/ioutil"
	"os"
)

func ParseToString(data *entity.Models) {
	Path := "D:\\project\\go\\go-trial\\data\\data.json"
	jsonFile, err := os.Open(Path)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &data)

}

func WriteFile(newData *entity.Model) {
	Path := "D:\\project\\go\\go-trial\\data\\data.json"

	file, err := ioutil.ReadFile(Path)
	if err != nil {
		fmt.Println(err)
	}

	var data entity.Models

	json.Unmarshal(file, &data)

	data.Models = append(data.Models, *newData)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(Path, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteAll(newData entity.Model) {
	Path := "D:\\project\\go\\go-trial\\data\\data.json"

	// var data []entity.Model
	dataBytes, err := json.Marshal(newData)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(Path, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
