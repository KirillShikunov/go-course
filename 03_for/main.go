package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	actManager := getActManager()

	scanner := bufio.NewScanner(os.Stdin)
	consoleManager := ConsoleManager{scanner}

	currentActId := actManager.getStartId()

	for {
		act := actManager.getAct(currentActId)
		consoleManager.showAct(act)

		optionIndex, err := consoleManager.getOptionIndex()

		if err != nil {
			fmt.Println("Помилка: Упсс... Щось пішло не запланом. Сталося наступна помилка:", err)
			break
		}

		if act.isExistOption(optionIndex) == false {
			fmt.Println("Помилка: Даного варіанта не існую, спробуйте ще раз.")
			continue
		}

		nextActId := act.getNextActId(optionIndex)
		if nextActId == 0 {
			fmt.Println("Нажаль цей варіант нікуди не веде, очікуйте продовження гри.")
			break
		}
		currentActId = nextActId
	}
}
