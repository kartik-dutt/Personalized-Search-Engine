package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"

	"fyne.io/fyne"
	app "fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/kartik-dutt/Simple-Search-Engine/src/inverted_index"
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

func DocToString(doc read_dataset.Document) string {
	s := doc.Title + "\n" + doc.URL + "\n" + doc.Text + "\n"
	return s
}

func main() {
	appl := app.New()
	win := appl.NewWindow("Your-Personal-Search-Engine")
	disp := "Loading Dataset Configurations!\n"
	win.SetContent(widget.NewLabel(disp))
	win.Resize(fyne.NewSize(800, 800))
	go func() {
		jsonFile, err := os.Open("./../config/dataset.json")
		if err != nil {
			log.Fatal("Error in Opening File!")
			panic(err)
		}

		disp += "Succesfully opened config/datset.json\n"
		disp += "=====================================\n\n"
		log.Println("Succesfully opened config/datset.json")
		defer jsonFile.Close()
		win.SetContent(widget.NewLabel(disp))

		byteArr, _ := ioutil.ReadAll(jsonFile)
		var datasets Datasets
		json.Unmarshal(byteArr, &datasets)
		docs := make([]read_dataset.Document, 0)
		disp += "Loading Dataset! \nThis may take 1 - 2 minutes as nearly 1 million urls are being loaded.\n"
		win.SetContent(widget.NewLabel(disp))
		for _, dataset := range datasets.DatasetsList {
			log.Println("Opening ./../data/" + dataset.Name)
			disp = disp + "Opening ./../data/" + dataset.Name
			win.SetContent(widget.NewLabel(disp))
			doc, _ := read_dataset.ReadDataset("../data/" + dataset.Name)
			docs = append(docs, doc...)
			disp += "\nClosed!\n"
			win.SetContent(widget.NewLabel(disp))
			log.Println("Closed!")
		}

		disp += "Loaded Dataset!\n"
		disp += "=====================================\n\n"
		win.SetContent(widget.NewLabel(disp))

		index := make(map[string][]int)
		disp += "Indexing the dataset!"
		win.SetContent(widget.NewLabel(disp))
		index = inverted_index.InvertedIndex(index, docs)
		disp += "\nIndexed Dataset!\n"
		disp += "================================\n\n"
		win.SetContent(widget.NewLabel(disp))

		input := widget.NewEntry()
		input.SetPlaceHolder("Search : ")
		searchRes := widget.NewLabel("Search Results")
		content := container.NewVBox(input, widget.NewButton("Search", func() {
			log.Println("Content was:", input.Text)
			txt := ""
			if input.Text != "" {
				res := search(index, input.Text)
				if len(res) > 0 {
					res = res[:int(math.Max(float64(len(res)), 10))]
				}
				for _, idx := range res {
					txt += ("\n" + DocToString(docs[idx]))
				}

				searchRes.SetText(txt)
				log.Println(txt)
			}
		}), searchRes)

		win.SetContent(widget.NewScrollContainer(content))
	}()

	win.ShowAndRun()
	return
}
