package game

import (
	// color "image/color"

	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const DT = 1 / 60

type Ship struct {
	Entity
	container *element
}

func newShip(container *element, state *State) *Ship {
	s := &Ship{
		container: container,
	}
	s.state = state
	s.speed = 400

	var err error
	s.img, _, err = ebitenutil.NewImageFromFile("ship.png")
	if err != nil {
		log.Fatal(err)
	}

	s.scale = 4
	s.op = &ebiten.DrawImageOptions{}
	return nil
}

func (s *Ship) onInit(state *State) error {
	return nil
}

func (s *Ship) onUpdate() error {
	var x, y = 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y -= 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y += 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x -= 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x += 1.0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		var bullet = new(Bullet)
		bullet.Init(s.state)
		bullet.SetPos(s.x+16, s.y)
		s.state.AddEntity(bullet)
	}
	s.vx += x * DT * 5
	s.vy += y * DT * 5
	s.vx = math.Max(math.Min(s.vx, 1.0), -1.0)
	s.vy = math.Max(math.Min(s.vy, 1.0), -1.0)
	s.vx *= 0.95
	s.vy *= 0.95
	s.x += s.vx * DT * s.speed
	s.y += s.vy * DT * s.speed
	if s.x > 800-64 {
		s.x = 800 - 64
	}
	if s.x < 0 {
		s.x = 0
	}
	if s.y < 0 {
		s.y = 0
	}
	if s.y > 600-64 {
		s.y = 600 - 64
	}
	s.container.position.x = s.x
	s.container.position.y = s.y
	return nil
}

func (s *Ship) onDraw(screen *ebiten.Image) error {
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(s.scale, s.scale)
	s.op.GeoM.Translate(s.container.position.x, s.container.position.x)
	screen.DrawImage(s.img, s.op)
	return nil
}
