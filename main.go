package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/usysrc/chaoshell/internal/game"
)

var myGame *game.Game

func init() {}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("ChaosHell")
	ebiten.SetVsyncEnabled(false)
	ebiten.SetTPS(60)

	myGame := &game.Game{}
	myGame.Init()
	if err := ebiten.RunGame(myGame); err != nil {
		if err == game.ErrTerminated {
			return
		}
		// Other termination
		log.Fatal(err)
	}
	// myGame.Init()
}
