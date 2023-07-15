package game

import (
	// color "image/color"
	// "math"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	Entity
}

func (e *Enemy) X() float64 {
	return e.x
}

func (e *Enemy) Y() float64 {
	return e.y
}

func (enemy *Enemy) SetPos(x, y float64) {
	enemy.x, enemy.y = x, y
	enemy.op.GeoM.Reset()
	enemy.op.GeoM.Scale(enemy.scale, enemy.scale)
	enemy.op.GeoM.Translate(enemy.x, enemy.y)
}

func (enemy *Enemy) Init(state *State) {
	enemy.state = state
	enemy.speed = 200
	var err error
	enemy.img, _, err = ebitenutil.NewImageFromFile("enemy.png")
	if err != nil {
		log.Fatal(err)
	}
	enemy.scale = 4
	enemy.op = &ebiten.DrawImageOptions{}
	enemy.vx = 0
	enemy.vy = 1
	enemy.UpdatePosition()
}

func (enemy *Enemy) UpdatePosition() {
	enemy.x += enemy.vx * DT * enemy.speed
	enemy.y += enemy.vy * DT * enemy.speed
	if enemy.y > 600 {
		SpawnEnemy(enemy.state)
		enemy.state.RemoveEntity(enemy)
	}

}

func (enemy *Enemy) Update() {
	enemy.UpdatePosition()
}

func (enemy *Enemy) Type() string {
	return "enemy"
}

func (s *Enemy) Draw(screen *ebiten.Image) {
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(s.scale, s.scale)
	s.op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.img, s.op)
}

func SpawnEnemy(state *State) {
	var enemy = &Enemy{}
	enemy.y = -64
	enemy.x = rand.Float64() * 800.0
	enemy.Init(state)
	state.AddEntity(enemy)
}
