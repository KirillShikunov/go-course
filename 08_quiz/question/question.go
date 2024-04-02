package question

type Question struct {
	text      string
	answers   []string
	correctId int
}

func (q *Question) IsCorrect(answerId int) bool {
	return q.correctId == answerId
}

func (q *Question) Answers() []string {
	return q.answers
}

func (q *Question) Text() string {
	return q.text
}

func (q *Question) GetAnswer(Id int) string {
	return q.answers[Id]
}

func (q *Question) GetCorrectAnswer() string {
	return q.answers[q.correctId]
}

func NewQuestion(question string, answers []string, correctId int) *Question {
	return &Question{question, answers, correctId}
}
