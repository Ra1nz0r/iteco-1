package player

import (
	"fmt"

	"github.com/Ra1nz0r/iteco-1/internal/box"
	"github.com/Ra1nz0r/iteco-1/internal/services"
)

type PlayersWithRandomChoice struct {
	id            int
	limitAttempts int
	found         bool
}

// Метод игрока, который делает попытки найти соответствующую шкатулку, до тех пор пока не достигнет лимита или не закончится удачно.
func (p *PlayersWithRandomChoice) MakeAttempts(boxes [](*box.Casket)) (bool, error) {
	if boxes == nil {
		return false, fmt.Errorf("nil dereference")
	}

	// Создаёт массив с количеством элементов, равных количеству шкатулок в текущей сессии.
	// Наполняет их номерами на которые ссылаются эти шкатулки.
	var CurrentSessionBoxesRow []int
	for _, boxId := range boxes {
		if boxId == nil {
			continue
		}

		CurrentSessionBoxesRow = append(CurrentSessionBoxesRow, boxId.Id)
	}

	// Cоздаем массив из номеров шкатулок, которые игрок проверит в порядке очереди.
	selected, errSel := box.SelectIds(&CurrentSessionBoxesRow, p.limitAttempts)
	if errSel != nil {
		return false, fmt.Errorf("%w", errSel)
	}

	// Делаем проверку каждой коробки и совпадения номера внутри неё с номером игрока.
	// Если есть совпадение, то функция сразу завершается со значением true.
	for _, s := range *selected {
		p.limitAttempts--
		res, errRes := box.FindByID(s, boxes)
		if errRes != nil {
			return false, fmt.Errorf("%w", errRes)
		}

		if p.id == res.Id {
			p.found = true
			return true, nil
		}

	}
	return false, nil
}

// Создаем и инициализируем массив из игроков, добавляем им количество попыток и устанавливаем результат его игры на false.
func CreatePlayersWithRandom(size, attemptsLimit int) []Unit {
	shuffled := services.IntArrShuffled(size)
	initedPlayers := make([]Unit, size)
	for i, pId := range *shuffled {
		initedPlayers[i] = &PlayersWithRandomChoice{id: pId, limitAttempts: attemptsLimit, found: false}
	}
	return initedPlayers
}
