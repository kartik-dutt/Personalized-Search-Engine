package text_cleaner

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	error_handler "github.com/kartik-dutt/Simple-Search-Engine/src/error_handler"
	snowballeng "github.com/kljensen/snowball/english"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func ToLowerCase(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}

	return tokens
}

func ReadWordsFromTxt(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func RemoveStopWords(tokens []string) []string {
	stopWords, err := ReadWordsFromTxt("../data/stop_words.txt")
	if err != nil {
		error_handler.ErrorHandler("Error reading stop words", err)
	}

	stopWordsSet := make(map[string]struct{})
	for _, stopWord := range stopWords {
		stopWordsSet[stopWord] = struct{}{}
	}

	res := make([]string, 0)
	for _, token := range tokens {
		if _, exists := stopWordsSet[token]; !exists {
			res = append(res, token)
		}
	}

	return res
}

func StemTokens(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}

	return r
}

func TextCleaner(text string) []string {
	tokens := Tokenize(text)
	tokens = ToLowerCase(tokens)
	tokens = RemoveStopWords(tokens)
	tokens = StemTokens(tokens)
	return tokens
}
