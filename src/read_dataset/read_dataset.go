package read_dataset

import (
	"encoding/xml"
	"os"

	error_handler "github.com/kartik-dutt/Simple-Search-Engine/src/error_handler"
)

// We expect the json file to have the following attributes :
// Title, URL and Abstract.
type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func ReadDataset(filepath string) ([]document, error) {
	xmlFile, err := os.Open(filepath)
	error_handler.ErrorHandler("Error opening "+filepath, err)
	defer xmlFile.Close()

	dec := xml.NewDecoder(xmlFile)
	dump := struct {
		Documents []document `xml: "doc"`
	}{}

	if err := dec.Decode(&dump); err != nil {
		error_handler.ErrorHandler("Error decoding xml file.", err)
		return nil, err
	}

	docs := dump.Documents
	// ID is simple the index.
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
