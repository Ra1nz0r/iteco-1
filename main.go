package main

import (
	"fmt"
	"math/rand/v2"
)

type Box struct {
	id int
}

type Player struct {
	id            int
	limitAttempts int
	found         bool
}

func (p *Player) MakeAttempts(boxes [](*Box)) bool {
	if boxes == nil {
		return false
	}

	var allIds []int

	for _, boxId := range boxes {
		if boxId == nil {
			continue
		}

		allIds = append(allIds, boxId.id)
	}

	selected := SelectBoxesIds(&allIds)

	for _, s := range *selected {
		fBox := FindBoxByID(s, boxes)

		if fBox == nil {
			return false
		}

		fmt.Println("Player id: ", p.id, "selected box id:", s)

		if p.id == fBox.id {
			p.found = true
			fmt.Println("=================GOOD ATTEMPT==============")
			return true
		}

	}

	return false
}

func FindBoxByID(id int, boxes [](*Box)) *Box {
	for _, b := range boxes {

		if b == nil {
			return nil
		}

		if id == b.id {
			return b
		}
	}
	return nil
}

func SelectBoxesIds(boxes *[]int) *[]int {
	boxesSelected := IntArrShuffled(len(*boxes))
	res := make([]int, len(*boxesSelected)/2)
	for i := range res {
		res[i] = (*boxesSelected)[i]
	}
	return &res
}

type GameSession struct {
	boxes   [](*Box)
	players [](*Player)
}

func NewGameSession(size, attemptsLimit int) *GameSession {
	gS := &GameSession{}
	gS.createBoxes(size)
	gS.createPlayer(size, attemptsLimit)

	return gS
}

func (gS *GameSession) PlaySession() bool {
	for _, player := range gS.players {
		if !player.MakeAttempts(gS.boxes) {
			return false
		}
	}

	return true
}

func (gS *GameSession) createBoxes(size int) {
	shuffled := IntArrShuffled(size)
	initedBoxes := make([](*Box), size)
	for i, bId := range *shuffled {
		initedBoxes[i] = &Box{id: bId}
	}
	gS.boxes = initedBoxes
}

func (gS *GameSession) createPlayer(size, attemptsLimit int) {
	shuffled := IntArrShuffled(size)
	initedPlayers := make([](*Player), size)
	for i, pId := range *shuffled {
		initedPlayers[i] = &Player{}
		initedPlayers[i].id = pId
		initedPlayers[i].limitAttempts = attemptsLimit
		initedPlayers[i].found = false
	}
	gS.players = initedPlayers
}

func IntArrShuffled(size int) *[]int {
	list := rand.Perm(size)
	for i := range list {
		list[i]++
	}
	return &list
}

func main() {
	size := 10
	attemptsPerPlayer := 5
	session := NewGameSession(size, attemptsPerPlayer)
	fmt.Println(session.PlaySession())
}
