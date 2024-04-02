package player

type Question interface {
	Text() string
	Answers() []string
}

type Player struct {
	id   int
	name string
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Id() int {
	return p.id
}

func NewPlayer(id int, name string) *Player {
	return &Player{id, name}
}
