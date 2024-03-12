package main

import (
	"fmt"
	"sort"
)

type Text struct {
	id int
}

func filterUniqueSortedTexts(texts []Text) []Text {
	uniqueTexts := clearFromDuplicates(texts)

	sortTextsById(uniqueTexts)

	return uniqueTexts
}

func clearFromDuplicates(texts []Text) []Text {
	uniqueTexts := make([]Text, 0, len(texts))
	seen := make(map[int]bool)

	for _, text := range texts {
		if seen[text.id] == false {
			uniqueTexts = append(uniqueTexts, text)
			seen[text.id] = true
		}
	}

	return uniqueTexts
}

func sortTextsById(texts []Text) {
	sort.Slice(texts, func(i, j int) bool {
		return texts[i].id < texts[j].id
	})
}

func main() {
	texts := []Text{
		{3},
		{30},
		{2},
		{1},
		{34},
		{5},
		{3},
		{5},
		{353},
		{35},
	}

	fmt.Println(filterUniqueSortedTexts(texts))
}
