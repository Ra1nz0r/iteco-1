package main

import (
	"fmt"

	"github.com/Ra1nz0r/iteco-1/internal/player"
	"github.com/Ra1nz0r/iteco-1/internal/session"

	"github.com/dariubs/percent"
)

func main() {
	size := 50
	attemptsPerPlayer := 25
	sessionsCount := 10000

	// Выбираем вариант со случайным выбором номера шкатулок.
	var mode player.PlayerType = player.WithRandom

	resRnd := Run(mode, size, attemptsPerPlayer, sessionsCount)
	fmt.Printf("Процент побед при случайном выборе, сессия из %d игр: %.0f%%.\n", sessionsCount, resRnd)

	// Переключаем на режим, где игроки договорились о способе выбора.
	mode = player.WithOrder

	resOrder := Run(mode, size, attemptsPerPlayer, sessionsCount)
	fmt.Printf("Процент побед при договорённости между игроками, сессия из %d игр: %.0f%%\n", sessionsCount, resOrder)

}

func Run(p player.PlayerType, size, attemptsPerPlayer, sessionsCount int) float64 {
	if attemptsPerPlayer > size {
		panic("GameSession initialization failed: attemptsPerPlayer > size")
	}

	var playersArr []player.Unit

	successedCount := 0
	// Запускаем цикл игровых сессий.
	for i := 0; i < sessionsCount; i++ {

		// В зависимости от Enum PlayerType выбираем реализацию интерфейса Unit и инициализируем переменную.
		switch p {
		case player.WithRandom:
			playersArr = player.CreatePlayersWithRandom(size, attemptsPerPlayer)
		case player.WithOrder:
			playersArr = player.CreatePlayersWithOrder(size, attemptsPerPlayer)
		}

		session := session.NewGameSession(size, playersArr)

		if session == nil {
			panic("GameSession initialization failed")
		}

		if session.PlaySession() {
			successedCount++
		}
	}

	return percent.PercentOf(successedCount, sessionsCount)
}
