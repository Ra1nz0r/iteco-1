package box

import (
	"github.com/Ra1nz0r/iteco-1/internal/services"
)

type Casket struct {
	Id int
}

// Поиск по массиву шкатулок, осуществляющийся по полю Id класса Casket и возвращает указатель на найденную.
func FindByID(id int, boxes [](*Casket)) *Casket {
	for _, b := range boxes {
		if b == nil {
			return nil
		}

		if id == b.Id {
			return b
		}
	}

	return nil
}

// Генерирует случайный порядок шкатулок, который будет открывать игрок.
func SelectIds(boxes *[]int, count int) *[]int {
	if count > len(*boxes) {
		return nil
	}

	boxesSelected := services.IntArrShuffled(len(*boxes))
	res := make([]int, count)
	for i := range res {
		res[i] = (*boxesSelected)[i]
	}

	return &res
}

// Создаем шкатулки, содержащие внутри номера в случайном порядке.
func CreateBoxes(count int) []*Casket {
	shuffled := services.IntArrShuffled(count)
	initedBoxes := make([](*Casket), count)
	for i, bId := range *shuffled {
		initedBoxes[i] = &Casket{Id: bId}
	}
	return initedBoxes
}
