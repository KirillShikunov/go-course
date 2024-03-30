package game

import (
	"08_kahoot/connection"
	"08_kahoot/console"
	"08_kahoot/player"
	"08_kahoot/random"
	"context"
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
	answerId := random.GetRandomAnswerId(question.Answers())

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

func NewSession(playerId int, connection *connection.Connection) *Session {
	return &Session{playerId: playerId, connection: connection}
}
