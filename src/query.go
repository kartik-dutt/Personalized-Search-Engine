package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	read_dataset "github.com/kartik-dutt/Simple-Search-Engine/src/read_dataset"
)

type Dataset struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Datasets struct {
	DatasetsList []Dataset `json:"dataset"`
}

func main() {
	jsonFile, err := os.Open("./../config/dataset.json")
	if err != nil {
		log.Fatal("Error in Opening File!")
		panic(err)
	}

	fmt.Println("Succesfully opened config/datset.json")
	defer jsonFile.Close()

	byteArr, _ := ioutil.ReadAll(jsonFile)
	var datasets Datasets
	json.Unmarshal(byteArr, &datasets)
	docs := make([]read_dataset.Document, 0)
	for _, dataset := range datasets.DatasetsList {
		fmt.Println("Opening ./../data/" + dataset.Name)
		doc, _ := read_dataset.ReadDataset("../data/" + dataset.Name)
		docs = append(docs, doc...)
		fmt.Println("Closed!")
	}
}
