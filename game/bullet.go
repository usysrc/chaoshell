package game

import (
	// color "image/color"
	"math"
	_ "image/png"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct{
	x, y float64
	vx, vy float64
	speed float64
	img *ebiten.Image
	op *ebiten.DrawImageOptions
	scale float64
	state *State
}

func (bullet *Bullet) SetPos (x, y float64) {
	bullet.x, bullet.y = x, y
	bullet.op.GeoM.Reset()
	bullet.op.GeoM.Scale(bullet.scale, bullet.scale)
	bullet.op.GeoM.Translate(bullet.x, bullet.y)
}

func (bullet *Bullet) Init(myState *State) {
	bullet.state = myState
	bullet.speed = 10
	var err error
	bullet.img, _, err = ebitenutil.NewImageFromFile("bullet.png")
	if err != nil {
		log.Fatal(err)
	}
	bullet.scale = 2
	bullet.op = &ebiten.DrawImageOptions{}
	bullet.op.GeoM.Reset()
	bullet.op.GeoM.Scale(bullet.scale, bullet.scale)
	bullet.vx = 0
	bullet.vy = -1.0
}

func (bullet *Bullet) Update() {
	bullet.x += bullet.vx * DT * bullet.speed
	bullet.y += bullet.vy * DT * bullet.speed
	if bullet.y < 0 {
		bullet.state.RemoveEntity(bullet)
	}
	bullet.op.GeoM.Reset()
	bullet.op.GeoM.Scale(bullet.scale, bullet.scale)
	bullet.op.GeoM.Translate(bullet.x, bullet.y)
}

func (s *Bullet) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.img, s.op)
}