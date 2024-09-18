package services

import (
	"math/rand/v2"
)

func Run() {

}

func IntArrShuffled(size int) *[]int {
	list := rand.Perm(size)
	for i := range list {
		list[i]++
	}
	return &list
}
