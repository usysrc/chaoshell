package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/usysrc/chaoshell/game"
)

var myGame *game.Game

func init() {}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetVsyncEnabled(false)
	ebiten.SetMaxTPS(60)

	myGame.Init()
	if err := ebiten.RunGame(myGame); err != nil {
		log.Fatal(err)
	}
	// myGame.Init()
}
