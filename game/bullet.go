package game

import (
	// color "image/color"
	// "math"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	Entity
}

func (b *Bullet) X() float64 {
	return b.x
}
func (b *Bullet) Y() float64 {
	return b.y
}

func (bullet *Bullet) SetPos(x, y float64) {
	bullet.x, bullet.y = x, y
	bullet.op.GeoM.Reset()
	bullet.op.GeoM.Scale(bullet.scale, bullet.scale)
	bullet.op.GeoM.Translate(bullet.x, bullet.y)
}

func (bullet *Bullet) Init(myState *State) {
	bullet.state = myState
	bullet.speed = 200
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
	w := 32.0
	h := 32.0
	found := false
	bullet.state.All(func(e IEntity) bool {
		if e != bullet && e.Type() != bullet.Type() {
			if bullet.x > e.X() && bullet.x < e.X()+w && bullet.y > e.Y() && bullet.y < e.Y()+h {
				bullet.state.RemoveEntity(e)
				found = true
				return false
			}
		}
		return true
	})

	if found {
		bullet.state.RemoveEntity(bullet)
	}
}

func (bullet *Bullet) Type() string {
	return "bullet"
}
