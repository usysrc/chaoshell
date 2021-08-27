package game

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {}

var (
	player  *Ship
	bg      *ebiten.Image
	myState *State
	timer   *Timer
)

type Game struct{}

func (g *Game) Init() {
	rand.Seed(time.Now().UnixNano())

	timer = new(Timer)
	timer.Init()

	myState = new(State)
	myState.Init()
	player = new(Ship)
	player.Init(myState)
	player.SetPos(360, 500)
	var err error
	bg, _, err = ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		timer.After(rand.Float64()*3.0, func() {
			SpawnEnemy(myState)
		})
	}
}

func (g *Game) Update() error {
	player.Update()
	myState.Update()
	timer.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)
	player.Draw(screen)
	myState.Draw(screen)
	timer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
