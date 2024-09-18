package player

import (
	"github.com/Ra1nz0r/iteco-1/internal/box"
	"github.com/Ra1nz0r/iteco-1/internal/services"
)

type PlayerWithOrderChoice struct {
	id            int
	limitAttempts int
	found         bool
}

func (p *PlayerWithOrderChoice) MakeAttempts(boxes [](*box.Casket)) bool {
	if boxes == nil {
		return false
	}

	checkID := p.id
	offset := 0
	for i := 0; i < p.limitAttempts; i++ {

		if checkID > len(boxes)-1 {
			return false
		}

		if p.id == boxes[checkID].Id+offset {
			p.found = true
			return true
		}

		checkID = boxes[checkID].Id - 1
		offset = -1
	}

	return false
}

func CreatePlayersWithOrder(size, attemptsLimit int) []Unit {
	shuffled := services.IntArrShuffled(size)
	initedPlayers := make([]Unit, size)
	for i, pId := range *shuffled {
		initedPlayers[i] = &PlayerWithOrderChoice{id: pId - 1, limitAttempts: attemptsLimit, found: false}
	}
	return initedPlayers
}
