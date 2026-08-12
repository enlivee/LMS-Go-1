package main

import (
	"fmt"
	"strings"
)

func AnalyzeText(text string) {

	f := func(c rune) bool {
		return c == '?' || c == '!' || c == '.' || c == ',' || c == ' '
	}


	words := strings.FieldsFunc(text, f)
	length := len(words)
	fmt.Println("Количество слов: ", length)

	uniqueWords := make(map[string]int)
	for _, r := range words{
		r = strings.ToLower(r)
		uniqueWords[r]++
	}
	fmt.Printf("Количество уникальных слов: %d\n", len(uniqueWords))
	
	top1 := getTopWords(uniqueWords, 1)
	if len(top1) > 0 {
		fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", top1[0], uniqueWords[top1[0]])
	}
	top5 := getTopWords(uniqueWords, 5)

	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, w := range top5 {
		fmt.Printf("\"%s\": %d раз\n", w, uniqueWords[w])
	}
}


func getTopWords(wordMap map[string]int, n int) []string {
	result := make([]string, 0, n)
	tempMap := make(map[string]int)
    for k, v := range wordMap {
        tempMap[k] = v
    }
	for i := 0; i < n; i++ {
        var bestWord string
        bestCount := -1
		for w, c := range tempMap {
			if c > bestCount {
				bestCount = c
				bestWord = w
			}
		}
		if bestCount == -1 { 
			break
		}

		result = append(result, bestWord)
		delete(tempMap, bestWord)
	}
	return result
}