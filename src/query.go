package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	inverted_index "github.com/kartik-dutt/Simple-Search-Engine/src/inverted_index"
	read_dataset "github.com/kartik-dutt/Simple-Search-Engine/src/read_dataset"
	text_cleaner "github.com/kartik-dutt/Simple-Search-Engine/src/text_cleaner"
)

type Dataset struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Datasets struct {
	DatasetsList []Dataset `json:"dataset"`
}

func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func search(index map[string][]int, text string) []int {
	res := make([]int, 0)
	for _, token := range text_cleaner.TextCleaner(text) {
		if idx, exists := index[token]; exists {
			if len(res) == 0 {
				res = idx
			} else {
				res = intersection(res, idx)
			}
		}
	}

	return res
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

	index := make(map[string][]int)
	index = inverted_index.InvertedIndex(index, docs)
	fmt.Println(search(index, "small wild cat"))
}
