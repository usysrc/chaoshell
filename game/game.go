package game

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {}

var (
	player *element
	bg     *ebiten.Image
	state  *State
	timer  *Timer
)

var ErrTerminated = errors.New("errTerminated")

type Game struct{}

func Spawn() {
	timer.After(rand.Float64()*3.0, func() {
		SpawnEnemy(state)
		Spawn()
	})
}

func (g *Game) Init() {
	rand.Seed(time.Now().UnixNano())

	timer = new(Timer)
	timer.Init()

	state = new(State)
	state.Init()
	player = &element{}
	player.addComponent(&Ship{})
	player.init(state)
	player.position.x = 360
	player.position.y = 500

	var err error
	bg, _, err = ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		log.Fatal(err)
	}
	Spawn()
}

func (g *Game) Update() error {
	player.update()
	state.Update()
	timer.Update()

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ErrTerminated
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)
	player.draw(screen)
	state.Draw(screen)
	timer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
