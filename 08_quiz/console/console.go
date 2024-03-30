package console

import (
	"08_kahoot/player"
	"08_kahoot/question"
	"fmt"
)

type Console struct {
}

func (c Console) ShowQuestion(question *question.Question) Console {
	fmt.Printf("Питання: %s\n", question.Text())
	fmt.Printf("Варіанти відповідей: %v\n", question.Answers())

	return c
}

func (c Console) ShowPlayerAnswer(playerId int, answer string) Console {
	fmt.Printf("Гравець #%d відповів: %s\n", playerId, answer)

	return c
}

func (c Console) ShowCorrectAnswer(question *question.Question) Console {
	fmt.Printf("Правильна відповідь: %s\n", question.GetCorrectAnswer())

	return c
}

func (c Console) ShowPlayerScore(player *player.Player, score int) Console {
	fmt.Printf("Гравець %s набрав %d бал(ів)\n", player.Name(), score)

	return c
}

func (c Console) ShowTimeoutMessage(playerId int, questionText string) Console {
	fmt.Printf("Гравець #%d не встиг відповісти на питання '%s'\n", playerId, questionText)

	return c
}

func (c Console) AddBreakLine() Console {
	fmt.Println()

	return c
}

func (c Console) ShowTimeout() {
	fmt.Println("Час вийшов")
}
