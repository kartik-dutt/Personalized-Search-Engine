package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/**
* Download the dataset.
* @param filepath : File path with file name.
* @param url Url that will be fetched.
 */
func downloadDataset(filepath string, url string) {

	if _, err := os.Stat(filepath); err == nil {
		fmt.Println("File Already exists. Terminating dowload. To download dataset again, delete the existing file.")
		return
	}

	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Error in creating file!")
		panic(err)
	}

	defer out.Close()
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal("Recieved Code :", resp.StatusCode)
		panic(err)
	}

	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Dataset downloaded!")
}

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
	for _, dataset := range datasets.DatasetsList {
		fmt.Println(dataset)
	}
}
