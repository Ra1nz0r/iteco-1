package session

import (
	"github.com/Ra1nz0r/iteco-1/internal/box"
	"github.com/Ra1nz0r/iteco-1/internal/player"
)

type GameSession struct {
	boxes   [](*box.Casket)
	players []player.Unit
}

func NewGameSession(size, attemptsLimit int, units []player.Unit) *GameSession {
	gS := &GameSession{}
	gS.boxes = box.CreateBoxes(size)
	gS.players = units
	if gS.boxes == nil || gS.players == nil {
		return nil
	}
	return gS
}

func (gS *GameSession) PlaySession() bool {
	for _, player := range gS.players {
		if !player.MakeAttempts(gS.boxes) {
			return false
		}
	}
	return true
}
