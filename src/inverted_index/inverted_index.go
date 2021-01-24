package inverted_index

import (
	read_dataset "github.com/kartik-dutt/Simple-Search-Engine/src/read_dataset"
	text_cleaner "github.com/kartik-dutt/Simple-Search-Engine/src/text_cleaner"
)

func InvertedIndex(index map[string][]int, docs read_dataset.Document) map[string][]int {
	for _, doc := range docs {
		for _, token := range text_cleaner.TextCleaner(doc.Text) {
			ids := index[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			index[token] = append(ids, doc.ID)
		}
	}

	return index
}
