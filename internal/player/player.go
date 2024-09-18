package player

import "github.com/Ra1nz0r/iteco-1/internal/box"

type Unit interface {
	MakeAttempts(boxes [](*box.Casket)) bool
}
