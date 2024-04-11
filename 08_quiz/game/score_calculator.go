package game

import "sync"

type ScoreCalculator struct {
	score map[int]int
	mu    sync.Mutex
}

func (s *ScoreCalculator) Increment(playerId int) {
	s.mu.Lock()
	s.score[playerId]++
	s.mu.Unlock()
}

func (s *ScoreCalculator) GetScore(playerId int) int {
	return s.score[playerId]
}

func NewScoreCalculator() *ScoreCalculator {
	return &ScoreCalculator{score: make(map[int]int)}
}
