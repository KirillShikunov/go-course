package game

import (
	"08_kahoot/config"
	"08_kahoot/connection"
	"08_kahoot/console"
	"08_kahoot/player"
	"context"
	"math/rand"
	"time"
)

type Session struct {
	playerId   int
	connection *connection.Connection
	console    console.Console
}

func (s *Session) PlayerId() int {
	return s.playerId
}

func (s *Session) Connection() *connection.Connection {
	return s.connection
}

func (s *Session) GetAnswerId() chan int {
	return s.connection.AnswerId()
}

func (s *Session) Start(ctx context.Context) {
	question := <-s.connection.Question()
	answerId := s.GetRandomAnswerId(question.Answers())

	select {
	case <-ctx.Done():
		// Показати користувачу помилку про те, що час вийшов
	default:
		s.connection.AnswerId() <- answerId
	}
}

func (s *Session) StreamQuestion(question player.Question) {
	s.connection.Question() <- question
}

func (s *Session) GetRandomAnswerId(answers []string) int {
	randomOffset := rand.Intn(5) - 3
	timeout := config.PlayRoundTimeoutInSecond + time.Duration(randomOffset)*time.Second

	if timeout > 0 {
		time.Sleep(timeout)
	}

	return rand.Intn(len(answers))
}

func NewSession(playerId int, connection *connection.Connection) *Session {
	return &Session{playerId: playerId, connection: connection}
}
