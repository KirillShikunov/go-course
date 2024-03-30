package connection

import (
	"08_kahoot/player"
)

type Connection struct {
	question chan player.Question
	answerId chan int
}

func (c *Connection) Question() chan player.Question {
	return c.question
}

func (c *Connection) AnswerId() chan int {
	return c.answerId
}

func NewConnection() *Connection {
	return &Connection{
		make(chan player.Question),
		make(chan int, 1),
	}
}
