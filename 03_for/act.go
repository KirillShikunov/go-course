package main

import "fmt"

type Option struct {
	nextActId int
	text      string
}

type Act struct {
	options []Option
	text    string
}

type ActManager struct {
	acts map[int]Act
}

func (actManager ActManager) getStartId() int {
	return 1
}

func (actManager ActManager) getAct(actId int) Act {
	act := actManager.acts[actId]
	return act
}

type ActPrinter struct {
}

func (actPrinter ActPrinter) showAct(act Act) {
	fmt.Printf("Ситуація: %s \n", act.text)

	fmt.Printf("Варіанти дій: \n")
	for i, option := range act.options {
		fmt.Printf("#%d %s \n", i+1, option.text)
	}
}

func getActManager() ActManager {
	return ActManager{
		acts: map[int]Act{
			1: {
				text: "Ви опиняєтесь на початку темного лісу. Куди підете?",
				options: []Option{
					{nextActId: 2, text: "На північ, до глибин лісу."},
					{nextActId: 3, text: "На схід, в напрямку річки."},
					{nextActId: 4, text: "На захід, до старої хатини."},
				},
			},
			2: {
				text: "Ви натрапили на зграю вовків. Як діятимете?",
				options: []Option{
					{nextActId: 5, text: "Спробувати обійти їх непомітно."},
					{nextActId: 6, text: "Використати факел для відлякування."},
					{nextActId: 7, text: "Підгодувати їх частиною свого провіанту."},
				},
			},
			3: {
				text: "Ви знаходите міст через річку, але він старий і хиткий. Як перейдете?",
				options: []Option{
					{nextActId: 8, text: "Швидко перебігти, не зважаючи на ризик."},
					{nextActId: 9, text: "Обережно, крок за кроком."},
					{nextActId: 10, text: "Пошукати інший шлях."},
				},
			},
			4: {
				text: "У хатині ви знаходите стару книгу з закляттями. Читатимете її?",
				options: []Option{
					{nextActId: 11, text: "Так, можливо знайду щось корисне."},
					{nextActId: 12, text: "Ні, краще не ризикувати."},
					{nextActId: 13, text: "Взяти її з собою і читати по дорозі."},
				},
			},
			5: {
				text: "Ви обходите вовків непомітно і потрапляєте до галявини з дивовижними квітами. Що робитимете?",
				options: []Option{
					{nextActId: 0, text: "Назбирати квітів."},
					{nextActId: 0, text: "Відпочити на галявині."},
					{nextActId: 0, text: "Продовжити шлях."},
				},
			},
			6: {
				text: "Використовуючи факел, ви відлякуєте вовків. Попереду ви бачите темний печерний вхід. Ваші дії?",
				options: []Option{
					{nextActId: 0, text: "Увійти в печеру."},
					{nextActId: 0, text: "Обійти її та продовжити свій шлях."},
				},
			},
			7: {
				text: "Годуючи вовків, ви знаходите дружню зв'язок з ними. Вони ведуть вас до таємної стежки. Слідувати за ними?",
				options: []Option{
					{nextActId: 0, text: "Так, слідувати за вовками."},
					{nextActId: 0, text: "Ні, продовжити своїм шляхом."},
				},
			},
			8: {
				text: "Перебігаючи міст, ви ледь не впали, але утримались. Попереду - загадковий ліс.",
				options: []Option{
					{nextActId: 0, text: "Розшукати вхід до таємної галявини."},
					{nextActId: 0, text: "Розташувати табір тут, на ніч."},
				},
			},
			9: {
				text: "Обережно перейшовши міст, ви помічаєте старовинний артефакт при дорозі.",
				options: []Option{
					{nextActId: 0, text: "Вивчити артефакт."},
					{nextActId: 0, text: "Залишити його та продовжити шлях."},
				},
			},
			10: {
				text: "Пошуки іншого шляху зайняли багато часу, але ви нарешті знайшли безпечний спуск до річки, де можна безпечно перейти.",
				options: []Option{
					{nextActId: 0, text: "Перейти річку і продовжити свій шлях."},
					{nextActId: 0, text: "Розглянути можливість відпочинку біля річки перед подальшим шляхом."},
				},
			},
			11: {
				text: "Читаючи книгу, ви знаходите закляття, яке може захистити вас від небезпек лісу.",
				options: []Option{
					{nextActId: 0, text: "Використати закляття зараз і продовжити свій шлях."},
					{nextActId: 0, text: "Зберегти закляття на більш важливий момент."},
				},
			},
			12: {
				text: "Ви вирішили не ризикувати з читанням книги та залишити її. Продовжуючи свій шлях, ви натрапляєте на таємниче озеро.",
				options: []Option{
					{nextActId: 20, text: "Підійти ближче та розглянути озеро."},
					{nextActId: 20, text: "Обійти озеро стороною і продовжити свій шлях."},
				},
			},
			13: {
				text: "Взявши книгу з собою, ви вирушаєте далі по лісу. Раптово перед вами з'являється старий мудрець.",
				options: []Option{
					{nextActId: 20, text: "Показати мудрецю книгу і попросити пораду."},
					{nextActId: 20, text: "Запитати мудреця про найкращий шлях через ліс."},
				},
			},
		},
	}
}
