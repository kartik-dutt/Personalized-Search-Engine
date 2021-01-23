package read_dataset

import (
	"os"

	error_handler "github.com/kartik-dutt/Simple-Search-Engine/error_handler"
)

// We expect the json file to have the following attributes :
// Title, URL and Abstract.
type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int32
}

func ReadDataset(filepath string) ([]document, error) {
	xmlFile, err := os.Open(filepath)
	error_handler.ErrorHandler("Error opening "+filepath, err)

}
