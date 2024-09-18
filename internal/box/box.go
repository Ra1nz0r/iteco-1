package box

import (
	"fmt"

	"github.com/Ra1nz0r/iteco-1/internal/services"
)

type Casket struct {
	Id int
}

func FindByID(id int, boxes [](*Casket)) (*Casket, error) {
	for _, b := range boxes {
		if b == nil {
			return nil, fmt.Errorf("failed: nil dereference")
		}

		if id == b.Id {
			return b, nil
		}
	}

	return nil, fmt.Errorf("failed: id not found")
}

func SelectIds(boxes *[]int, count int) (*[]int, error) {
	if count > len(*boxes) {
		return nil, fmt.Errorf("failed: out of bounds")
	}

	boxesSelected := services.IntArrShuffled(len(*boxes))
	res := make([]int, count)
	for i := range res {
		res[i] = (*boxesSelected)[i]
	}

	return &res, nil
}

// Создаем и инициализируем шкатулки в случайном порядке номеров,
// добавляем им количество попыток и устанавливаем результат его игры на false.
func CreateBoxes(size int) []*Casket {
	shuffled := services.IntArrShuffled(size)
	initedBoxes := make([](*Casket), size)
	for i, bId := range *shuffled {
		initedBoxes[i] = &Casket{Id: bId}
	}
	return initedBoxes
}
