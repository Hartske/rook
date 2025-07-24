package main

func commandDraw(ctx *GameContext, args ...string) error {
	ctx.decideDealer()
	return nil
}

func commandDeal(ctx *GameContext, args ...string) error {
	ctx.deal()
	return nil
}
