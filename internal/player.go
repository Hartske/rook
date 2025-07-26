package internal

type Player struct {
	Name     string
	Score    int
	Bid      int
	Hand     []*Card
	IsDealer bool
}
