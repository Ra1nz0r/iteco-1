package services

import (
	"math/rand/v2"
)

// Возвращает массив из целых чисел, заполненный случайными образом, от 1 до size включительно.
func IntArrShuffled(size int) *[]int {
	list := rand.Perm(size)
	for i := range list {
		list[i]++
	}
	return &list
}
