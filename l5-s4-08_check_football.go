package main

import (
	"fmt"
)

type Player struct {
    Name      string
    Goals     int
    Misses    int
    Assists   int
    Rating    float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
}

func main() {

}
/*
Программист Арсений заинтересовался футболом, но просто играть ему мало. 
Он решил проанализировать результаты своей любимой команды и помочь улучшить их.

Программа для анализа должна:
* Хранить данные о каждом футболисте: Имя, Голы, Промахи и Помощь.
* Рассчитывать рейтинг для каждого игрока по формуле: (Голы + Помощь / 2) / Промахи (если количество промахов равно нулю, то Голы + Помощь / 2).
Сортировка по:
	* Убыванию количества голов (функция goalsSort(players []Player) []Player)
	* Убыванию рейтинга (функция ratingSort(players []Player) []Player)
	* Убыванию отношения голов к промахам (функция gmSort(players []Player) []Player) Если рейтинг одинаковый, то во всех функциях необходимо сортировать по имени (Name) в алфавитном порядке.
Также нужно реализовать такую структуру:
type Player struct {
    Name      string
    Goals     int
    Misses    int
    Assists   int
    Rating    float64
}
и вспомогательный метод calculateRating() для расчёта поля Rating на основе входных данных. 
*Конструктор NewPlayer(name string, goals, misses, assists int) Player создаёт нового игрока и вычисляет его рейтинг с помощью calculateRating().

Примечания
    Для разработки программы крайне рекомендуется использовать GitHub.
*/
