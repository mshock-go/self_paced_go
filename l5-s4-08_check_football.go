package main

import (
	"fmt"
	"slices"
	"strings"
)

type Player struct {
    Name      string
    Goals     int
    Misses    int
    Assists   int
    Rating    float64
    gm        float64
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(p1, p2 Player) int {
		switch  {
		case p2.Goals < p1.Goals:
			return -1
		case p2.Goals > p1.Goals:
			return 1
		default:
			return strings.Compare(p1.Name, p2.Name)
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(p1, p2 Player) int {
		switch  {
		case p2.Rating < p1.Rating:
			return -1
		case p2.Rating > p1.Rating:
			return 1
		default:
			return strings.Compare(p1.Name, p2.Name)
		}
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(p1, p2 Player) int {		
		switch  {
		case p2.gm < p1.gm:
			return -1
		case p2.gm > p1.gm:
			return 1
		default:
			return strings.Compare(p1.Name, p2.Name)
		}
	})
	return players
}

func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals + p.Assists / 2.0)
	} else {
		p.Rating = float64(p.Goals + p.Assists / 2.0) / float64(p.Misses)
	}
}

func NewPlayer(name string, goals, misses, assists int) (p Player) {
	p = Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	p.calculateRating()
	p.gm = float64(p.Goals) / float64(p.Misses)
	return p
}

func main() {
	p1 := NewPlayer("Maradona", 16, 6, 12)
	p2 := NewPlayer("Pele", 28, 2, 16)
	p3 := NewPlayer("Beckham", 28, 2, 16)
	p4 := NewPlayer("Ronaldo", 18, 4, 21)
	p5 := NewPlayer("Zidane", 14, 8, 14)
	p6 := NewPlayer("Beckenbauer", 14, 8, 10)
	
	players := []Player {p1, p2, p3, p4, p5, p6}
	for _, p := range goalsSort(players) {
		fmt.Printf("%#v\n", p)
	}
	println()
	players = []Player {p1, p2, p3, p4, p5, p6}
	for _, p := range ratingSort(players) {
		fmt.Printf("%#v\n", p)
	}
	println()
	players = []Player {p1, p2, p3, p4, p5, p6}
	for _, p := range gmSort(players) {
		fmt.Printf("%#v\n", p)
	}
}
/*
Программист Арсений заинтересовался футболом, но просто играть ему мало. 
Он решил проанализировать результаты своей любимой команды и помочь улучшить их.

Программа для анализа должна:
+ Хранить данные о каждом футболисте: Имя, Голы, Промахи и Помощь.
  нужно реализовать такую структуру:
	type Player struct {
		Name      string
		Goals     int
		Misses    int
		Assists   int
		Rating    float64
	}
+ Рассчитывать рейтинг для каждого игрока по формуле: (Голы + Помощь / 2) / Промахи (если количество промахов равно нулю, то Голы + Помощь / 2).
  вспомогательный метод calculateRating() для расчёта поля Rating на основе входных данных. 
+ Конструктор NewPlayer(name string, goals, misses, assists int) Player создаёт нового игрока и вычисляет его рейтинг с помощью calculateRating().
Сортировка по:
	* Убыванию количества голов (функция goalsSort(players []Player) []Player)
	* Убыванию рейтинга (функция ratingSort(players []Player) []Player)
	* Убыванию отношения голов к промахам (функция gmSort(players []Player) []Player) 
Если рейтинг одинаковый, то во всех функциях необходимо сортировать по имени (Name) в алфавитном порядке.

Примечания
    Для разработки программы крайне рекомендуется использовать GitHub.
Ввод
func TestCalculateRating(t *testing.T) {
	tests := []struct {
		name           string
		player         Player
		expectedRating float64
	}{
		{
			name:           "No misses",
			player:         NewPlayer("Player1", 10, 0, 4),
			expectedRating: 12.0,
		},
		{
			name:           "Normal case",
			player:         NewPlayer("Player2", 8, 2, 4),
			expectedRating: 5.0,
		},
		{
			name:           "Zero assists",
			player:         NewPlayer("Player3", 5, 1, 0),
			expectedRating: 5.0,
		},
		{
			name:           "All stats",
			player:         NewPlayer("Player4", 12, 3, 6),
			expectedRating: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.player.Rating != tt.expectedRating {
				t.Errorf("Expected rating %.2f, got %.2f", tt.expectedRating, tt.player.Rating)
			}
		})
	}
}

func TestGoalsSort(t *testing.T) {
	players := []Player{
		NewPlayer("Player1", 10, 5, 3),
		NewPlayer("Player2", 15, 7, 2),
		NewPlayer("Player3", 8, 2, 5),
	}

	s... File is too long to be displayed fully
*/
