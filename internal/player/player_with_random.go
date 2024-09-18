package player

import (
	"log"

	"github.com/Ra1nz0r/iteco-1/internal/box"
	"github.com/Ra1nz0r/iteco-1/internal/services"
)

type PlayerWithRandomChoice struct {
	id            int
	limitAttempts int
	found         bool
}

func (p *PlayerWithRandomChoice) MakeAttempts(boxes [](*box.Casket)) bool {
	if boxes == nil {
		return false
	}

	// превращаем массив коробок в массив из ID этих коробок чтобы пройти по ним
	var allIds []int
	for _, boxId := range boxes {
		if boxId == nil {
			continue
		}

		allIds = append(allIds, boxId.Id)
	}

	// создаем массив из ID шкатулок которые Player проверит в порядке очереди
	selected, errSelected := box.SelectIds(&allIds, p.limitAttempts)
	if errSelected != nil {
		log.Fatal(errSelected)
	}

	// проверяем ID выбранной коробки с ID плеера в порядке очереди
	for _, s := range *selected {
		p.limitAttempts--
		res, errRes := box.FindByID(s, boxes)
		if errRes != nil {
			log.Fatal(errRes)
		}

		if res == nil { //////
			return false
		}

		if p.id == res.Id {
			p.found = true
			return true
		}

	}

	return false
}

// Создаем и инициализируем очередь из игроков в случайном порядке номеров,
// добавляем им количество попыток и устанавливаем результат его игры на false.
func CreatePlayersWithRandom(size, attemptsLimit int) []Unit {
	shuffled := services.IntArrShuffled(size)
	initedPlayers := make([]Unit, size)
	for i, pId := range *shuffled {
		initedPlayers[i] = &PlayerWithRandomChoice{id: pId, limitAttempts: attemptsLimit, found: false}
	}
	return initedPlayers
}
