package game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {}

var (
	player *Ship
	bg *ebiten.Image
	myState *State
)

type Game struct{}

func (g *Game) Init() {
	myState = new(State)
	myState.Init()
	player = new(Ship)
	player.Init(myState)
	player.SetPos(360, 500)
	bg = ebiten.NewImage(800,600)
	bg.Fill(color.White)
}

func (g *Game) Update() error {
	player.Update()
	myState.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)
	player.Draw(screen)
	myState.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}