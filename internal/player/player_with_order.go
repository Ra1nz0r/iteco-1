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

// Метод игрока, который делает попытки найти соответствующую шкатулку, до тех пор пока не достигнет лимита или не закончится удачно.
func (p *PlayerWithOrderChoice) MakeAttempts(boxes [](*box.Casket)) bool {
	if boxes == nil {
		return false
	}

	checkID := p.id
	offset := 0

	// Выполняем попытки найти шкатулку в цикле.
	for i := 0; i < p.limitAttempts; i++ {

		if checkID > len(boxes)-1 {
			return false
		}

		// Сравниваем номер игрока с номером внутри шкатулку, если находим то завершаем функцию.
		if p.id == boxes[checkID].Id+offset {
			p.found = true
			return true
		}

		// Если попытка неудачная, то обновлняем переменные для цикла.
		checkID = boxes[checkID].Id - 1
		offset = -1
	}

	return false
}

// Создаем и инициализируем массив из игроков, добавляем им количество попыток и устанавливаем результат его игры на false.
func CreatePlayersWithOrder(size, attemptsLimit int) []Unit {
	shuffled := services.IntArrShuffled(size)
	initedPlayers := make([]Unit, size)
	for i, pId := range *shuffled {
		initedPlayers[i] = &PlayerWithOrderChoice{id: pId - 1, limitAttempts: attemptsLimit, found: false}
	}
	return initedPlayers
}
