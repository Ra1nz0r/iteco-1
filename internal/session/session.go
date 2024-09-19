package session

import (
	"github.com/Ra1nz0r/iteco-1/internal/box"
	"github.com/Ra1nz0r/iteco-1/internal/player"
)

type GameSession struct {
	boxes   [](*box.Casket)
	players []player.Unit
}

// Фнукция-конструктор игровой сессии.
func NewGameSession(size int, units []player.Unit) *GameSession {
	gS := &GameSession{
		boxes:   box.CreateBoxes(size),
		players: units,
	}

	if gS.boxes == nil || gS.players == nil {
		return nil
	}

	return gS
}

// Запускает игровую сессию со всеми участниками и возвращает её результат.
func (gS *GameSession) PlaySession() bool {
	for _, player := range gS.players {
		if !player.MakeAttempts(gS.boxes) {
			return false
		}
	}
	return true
}
