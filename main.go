package main

import (
	color "image/color"
	_ "image/png"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/usysrc/chaoshell/ship"
)

var player *ship.Ship
var bg *ebiten.Image

func init() {
	player = new(ship.Ship)
	
	player.Init()
	player.SetPos(360, 500)
	bg = ebiten.NewImage(800,600)
	bg.Fill(color.RGBA{0xff,0xff,0xff,0xff})
}

type Game struct{}

func (g *Game) Update() error {
	player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)
	player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}


func main() {

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
