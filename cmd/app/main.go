package main

import (
	"github.com/Ra1nz0r/iteco-1/internal/session"
)

func main() {
	size := 50
	attemptsPerPlayer := 25
	sessionsCount := 1000

	session.Run(size, attemptsPerPlayer, sessionsCount)
}
