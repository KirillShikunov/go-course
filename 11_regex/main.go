package main

import (
	"11_regex/search"
	"fmt"
)

func main() {
	findPhones()
	findVowelConstantWords()
}

func findVowelConstantWords() {
	const regex = `(?m)(?:^|\s|[,.!?-])([АЕЄИІЇОУЮЯаеєиіїоуюя][а-яА-ЯїЇіІєЄ]*[бвгґджзйклмнпрстфхцчшщБВГҐДЖЗЙКЛМНПРСТФХЦЧШЩ])(?:$|\s|[,.!?-])`

	err, words := search.FindInFile("file/text.txt", regex)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Слова, що починаються на голосну і закінчуються на приголосну:")
	fmt.Println(words)
}

func findPhones() {
	const regex = `(?m)(?m)^\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`

	err, phones := search.FindInFile("file/phones.txt", regex)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Номери телефонів:")
	fmt.Println(phones)
}
