package score

import "sync"

type Calculator struct {
	score map[int]int
	mu    sync.Mutex
}

func (c *Calculator) Increment(playerId int) {
	c.mu.Lock()
	c.score[playerId]++
	c.mu.Unlock()
}

func (c *Calculator) GetScore(playerId int) int {
	return c.score[playerId]
}

func NewScoreCalculator() *Calculator {
	return &Calculator{score: make(map[int]int)}
}
