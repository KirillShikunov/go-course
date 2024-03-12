package main

import (
	"strings"
)

type TextManager struct {
	rows []string
}

func (textManager TextManager) searchRows(searchWord string) []string {
	var matchedRows []string
	for _, row := range textManager.rows {
		if strings.Contains(strings.ToLower(row), strings.ToLower(searchWord)) {
			matchedRows = append(matchedRows, row)
		}
	}

	return matchedRows
}

func getTextManager() (TextManager, error) {
	fileReader := FileReader{}
	lines, err := fileReader.read("texts.txt")

	return TextManager{lines}, err
}
