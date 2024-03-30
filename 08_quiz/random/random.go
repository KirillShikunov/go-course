package random

import (
	"08_kahoot/config"
	"math/rand"
	"time"
)

func GetRandomAnswerId(answers []string) int {
	randomOffset := rand.Intn(5) - 3
	timeout := config.PlayRoundTimeoutInSecond + time.Duration(randomOffset)*time.Second

	if timeout > 0 {
		time.Sleep(timeout)
	}

	return rand.Intn(len(answers))
}
