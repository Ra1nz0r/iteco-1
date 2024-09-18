package session

import (
	"fmt"

	"github.com/Ra1nz0r/iteco-1/internal/player"
)

func Run(size, attemptsPerPlayer, sessionsCount int) {

	var successedWithRandom, successedWithOrder int

	for i := 0; i < sessionsCount; i++ {
		sessionRandom := NewGameSession(size, attemptsPerPlayer, player.CreatePlayersWithRandom(size, attemptsPerPlayer))

		sessionOrder := NewGameSession(size, attemptsPerPlayer, player.CreatePlayersWithOrder(size, attemptsPerPlayer))

		if sessionRandom == nil || sessionOrder == nil {
			panic("GameSession initialization failed")
		}

		if sessionRandom.PlaySession() {
			successedWithRandom++
		}

		if sessionOrder.PlaySession() {
			successedWithOrder++
		}
	}

	difRnd := sessionsCount - successedWithRandom
	difOrder := sessionsCount - successedWithOrder

	fmt.Printf("Successed: %d, Failed: %d.\n", successedWithRandom, difRnd)
	fmt.Printf("Successed: %d, Failed: %d.\n", successedWithOrder, difOrder)
}
