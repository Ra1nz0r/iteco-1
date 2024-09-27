package session

import (
	"fmt"

	"github.com/Ra1nz0r/iteco-1/internal/box"
	"github.com/Ra1nz0r/iteco-1/internal/player"
)

type GameSession struct {
	boxes   [](*box.Casket)
	players []player.Unit
}

// Фнукция-конструктор игровой сессии.
func NewGameSession(size int, units []player.Unit) (*GameSession, error) {
	gS := &GameSession{
		boxes:   box.CreateBoxes(size),
		players: units,
	}

	if gS.boxes == nil || gS.players == nil {
		return nil, fmt.Errorf("failed: nil dereference")
	}

	return gS, nil
}

// Запускает игровую сессию со всеми участниками и возвращает её результат.
func (gS *GameSession) PlaySession() (bool, error) {
	for _, player := range gS.players {
		ok, err := player.MakeAttempts(gS.boxes)
		if err != nil {
			return false, fmt.Errorf("%w", err)
		}

		if !ok {
			return false, nil
		}
	}
	return true, nil
}
