package read_dataset

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	error_handler "github.com/kartik-dutt/Simple-Search-Engine/src/error_handler"
)

// We expect the json file to have the following attributes :
// Title, URL and Abstract.
type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

type Documents struct {
	DocumentsList []Document `xml:"doc"`
}

func ReadDataset(filepath string) ([]Document, error) {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		error_handler.ErrorHandler("Error opening "+filepath, err)
	}

	defer xmlFile.Close()
	byteArr, err := ioutil.ReadAll(xmlFile)

	var docs Documents
	xml.Unmarshal(byteArr, &docs)
	for i, _ := range docs.DocumentsList {
		docs.DocumentsList[i].ID = i
	}

	return docs.DocumentsList, nil
}
