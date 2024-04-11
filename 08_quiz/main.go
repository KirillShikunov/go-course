package main

import (
	"08_kahoot/game"
	"08_kahoot/player"
	"08_kahoot/question"
)

func main() {
	var questions []*question.Question
	questions = append(
		questions,
		question.NewQuestion("Скільки місяців з кількістю днів 30?", []string{"1", "2", "3", "4"}, 3),
	)
	questions = append(
		questions,
		question.NewQuestion("Хто написав 'Гамлета'?", []string{"Чарльз Діккенс", "Вільям Шекспір", "Лео Толстой", "Марк Твен"}, 1),
	)

	questions = append(
		questions,
		question.NewQuestion("Столиця Франції?", []string{"Мадрид", "Лондон", "Берлін", "Париж"}, 3),
	)

	questions = append(
		questions,
		question.NewQuestion("Найбільший океан на Землі?", []string{"Атлантичний", "Індійський", "Тихий", "Північний Льодовитий"}, 2),
	)

	questions = append(
		questions,
		question.NewQuestion("Скільки континентів на Землі?", []string{"5", "6", "7", "8"}, 2),
	)

	questions = append(
		questions,
		question.NewQuestion("Який газ є основним у складі атмосфери Землі?", []string{"Оксиген", "Азот", "Вуглекислий газ", "Гелій"}, 1),
	)

	processor := question.NewProcessor(questions)
	newGame := game.NewGame(processor, game.NewScoreCalculator())

	newGame.AddPlayer(player.NewPlayer(1, "Андрій"))
	newGame.AddPlayer(player.NewPlayer(2, "Олександр"))
	newGame.AddPlayer(player.NewPlayer(3, "Інна"))

	newGame.Start()
}
