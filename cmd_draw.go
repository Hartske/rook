package main

func commandDraw(ctx *GameContext, args ...string) error {
	ctx.decideDealer()
	return nil
}
