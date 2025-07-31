package internal

type Play struct {
	HighBid    int
	HighBidder string
	Trump      string
	Pot        []*Card
}

func (p *Play) ResetPot() {
	p.Pot = make([]*Card, 0)
}
