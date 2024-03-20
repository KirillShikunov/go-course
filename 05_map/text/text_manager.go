package text

import (
	"strings"
)

type TextManager struct {
	rows  []string
	index map[string]map[int]struct{}
}

func (textManager *TextManager) SearchWordByContains(searchWord string) []string {
	var matchedRows []string

	for _, row := range textManager.rows {
		if strings.Contains(strings.ToLower(row), strings.ToLower(searchWord)) {
			matchedRows = append(matchedRows, row)
		}
	}

	return matchedRows
}

func (textManager *TextManager) SearchWordByIndex(searchWord string) []string {
	var matchedRows []string

	searchWord = textManager.prepareSearchWord(searchWord)
	if textManager.index[searchWord] == nil {
		return matchedRows
	}

	for rowIndex, _ := range textManager.index[searchWord] {
		matchedRows = append(matchedRows, textManager.rows[rowIndex])
	}

	return matchedRows
}

func (textManager *TextManager) initIndex() {
	for rowIndex, row := range textManager.rows {
		words := strings.Split(row, " ")

		for _, word := range words {
			word = textManager.prepareSearchWord(word)
			if textManager.index[word] == nil {
				textManager.index[word] = make(map[int]struct{})
			}

			if _, ok := textManager.index[word][rowIndex]; ok {
				continue
			}

			textManager.index[word][rowIndex] = struct{}{}
		}
	}
}

func (textManager *TextManager) prepareSearchWord(word string) string {
	return strings.ToLower(strings.TrimSpace(strings.Trim(word, ".,!-?")))
}

func NewTextManager(lines []string) TextManager {
	textManager := TextManager{lines, make(map[string]map[int]struct{})}
	textManager.initIndex()

	return textManager
}
