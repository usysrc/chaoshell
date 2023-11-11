package game

import (
	"errors"
	"log"
	"math/rand"

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

var ErrTerminated = errors.New("errTerminated")

type Game struct{}

func Spawn() {
	timer.After(rand.Float64()*3.0, func() {
		SpawnEnemy(myState)
		Spawn()
	})
}

func (g *Game) Init() {
	timer = new(Timer)
	timer.Init()

	myState = &State{}
	myState.Init()

	player = &Ship{}
	player.Init(myState)
	player.SetPos(360, 500)

	var err error
	bg, _, err = ebitenutil.NewImageFromFile("internal/assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	Spawn()
}

func (g *Game) Update() error {
	player.Update()
	myState.Update()
	timer.Update()

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ErrTerminated
	}
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
