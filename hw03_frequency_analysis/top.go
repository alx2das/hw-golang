package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// no comments
var regexpNormalize = regexp.MustCompile(`\p{L}+(-\p{L}+)*|-{2,}`)

type wordData struct {
	word  string
	count int
}

func Top10(input string) []string {
	// нормализуем данные
	input = strings.ToLower(input)
	words := regexpNormalize.FindAllString(input, -1)

	// считаем кол-во уникальных слов
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// в слайсоструктуру
	wordList := make([]wordData, 0, len(wordCount))
	for word, count := range wordCount {
		wordList = append(wordList, wordData{word, count})
	}

	// сортируем
	sort.Slice(wordList, func(i, j int) bool {
		// если кол-во одинаковао, сортируем лексикографически
		if wordList[i].count == wordList[j].count {
			return wordList[i].word < wordList[j].word
		}

		// сортировка по кол-ву
		return wordList[i].count > wordList[j].count
	})

	// подготавливаем результат
	result := make([]string, 0, 10)
	for i := 0; i < len(wordList) && i < 10; i++ {
		result = append(result, wordList[i].word)
	}

	return result
}
