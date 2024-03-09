package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	actManager := getActManager()
	scanner := bufio.NewScanner(os.Stdin)
	currentActId := actManager.getStartId()

	for {
		act := actManager.getAct(currentActId)

		actManager.showAct(act)

		option, err := scanOption(scanner)
		if isExistOption(err, option, act) == false {
			fmt.Println("Даного варіанта не існую, спробуйте ще раз.")
			continue
		}

		nextActId := getNextActId(act, option)
		if nextActId == 0 {
			fmt.Println("Нажаль цей варіант нікуди не веде, очікуйте продовження гри.")
			break
		}
		currentActId = nextActId
	}
}

func getNextActId(act Act, option int) int {
	return act.options[option-1].nextActId
}

func isExistOption(err error, option int, act Act) bool {
	return err == nil && option >= 1 && option <= len(act.options)
}

func scanOption(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	input := scanner.Text()
	option, err := strconv.Atoi(input)
	return option, err
}
