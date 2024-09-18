package player

import "github.com/Ra1nz0r/iteco-1/internal/box"

type Unit interface {
	MakeAttempts(boxes [](*box.Casket)) bool
}

// Enum для выбора типа игрока.
type PlayerType int

const (
	WithRandom PlayerType = iota
	WithOrder
)
