package main

import (
	"fmt"

	"github.com/Ra1nz0r/iteco-1/internal/player"
	"github.com/Ra1nz0r/iteco-1/internal/session"
)

func main() {
	size := 50
	attemptsPerPlayer := 25
	sessionsCount := 1000

	// Выбор варианта поведения игрока.
	var mode player.PlayerType = player.WithOrder

	Run(mode, size, attemptsPerPlayer, sessionsCount)
}

func Run(p player.PlayerType, size, attemptsPerPlayer, sessionsCount int) {
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

	diffOrder := sessionsCount - successedCount

	fmt.Printf("Successed: %d,\nFailed: %d,\nTotal: %d.\n", successedCount, diffOrder, sessionsCount)
}
