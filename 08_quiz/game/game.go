package game

import (
	"08_kahoot/config"
	"08_kahoot/connection"
	"08_kahoot/console"
	"08_kahoot/player"
	"08_kahoot/question"
	"08_kahoot/score"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Game struct {
	players           []*player.Player
	sessions          []*Session
	questionProcessor *question.Processor
	scoreCalculator   *score.Calculator
	console           console.Console
}

func (g *Game) AddPlayer(newPlayer *player.Player) {
	g.players = append(g.players, newPlayer)
	g.sessions = append(g.sessions, NewSession(newPlayer.Id(), connection.NewConnection()))
}

func (g *Game) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Гра була перервана. Завершення...")
			g.showScores()
			return
		default:
			currentQuestion, ok := g.questionProcessor.GetQuestion()
			if ok != nil {
				fmt.Println(ok)
				return
			}

			g.playRound(ctx, currentQuestion)

			if g.questionProcessor.IsLast() {
				g.showScores()
				return
			}

			g.questionProcessor.Increment()
		}
	}
}

func (g *Game) playRound(consoleCtx context.Context, question *question.Question) {
	ctx, cancel := context.WithTimeout(consoleCtx, config.PlayRoundTimeoutInSecond)
	defer cancel()

	g.console.ShowQuestion(question)

	go func() {
		select {
		case <-ctx.Done():
			g.console.ShowTimeout()
		}
	}()

	wg := &sync.WaitGroup{}
	for _, session := range g.sessions {
		wg.Add(1)

		go session.Start(ctx)

		go func(session *Session) {
			defer wg.Done()

			session.StreamQuestion(question)

			select {
			case <-ctx.Done():
				g.console.ShowTimeoutMessage(session.PlayerId(), question.Text())
			case answerId := <-session.GetAnswerId():
				if question.IsCorrect(answerId) {
					g.scoreCalculator.Increment(session.PlayerId())
				}
				g.console.ShowPlayerAnswer(session.PlayerId(), question.GetAnswer(answerId))
			}
		}(session)
	}

	wg.Wait()

	g.console.ShowCorrectAnswer(question).AddBreakLine()
}

func (g *Game) showScores() {
	fmt.Println("Кінець гри. Результати балів:")
	for _, p := range g.players {
		g.console.ShowPlayerScore(p, g.scoreCalculator.GetScore(p.Id()))
	}
}

func NewGame(processor *question.Processor, calculator *score.Calculator) *Game {
	return &Game{
		players:           []*player.Player{},
		questionProcessor: processor,
		scoreCalculator:   calculator,
	}
}
