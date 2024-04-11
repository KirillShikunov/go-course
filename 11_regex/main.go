package main

import (
	"11_regex/search"
	"fmt"
	"regexp"
)

func findPhones() {
	regex := regexp.MustCompile(`(?m)(?m)^\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`)

	phones, err := search.FindInFile("file/phones.txt", regex)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Номери телефонів:")
	fmt.Println(phones)
}

func findVowelConstantWords() {
	regex := regexp.MustCompile(`(?m)(?:^|\s|[,.!?-])([АЕЄИІЇОУЮЯаеєиіїоуюя][а-яА-ЯїЇіІєЄ]*[бвгґджзйклмнпрстфхцчшщБВГҐДЖЗЙКЛМНПРСТФХЦЧШЩ])(?:$|\s|[,.!?-])`)

	words, err := search.FindInFile("file/text.txt", regex)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Слова, що починаються на голосну і закінчуються на приголосну:")
	fmt.Println(words)
}

func main() {
	findPhones()
	findVowelConstantWords()
}
