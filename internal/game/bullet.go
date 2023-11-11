package game

import (
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

func (b *Bullet) SetPos(x, y float64) {
	b.x, b.y = x, y
	b.op.GeoM.Reset()
	b.op.GeoM.Scale(b.scale, b.scale)
	b.op.GeoM.Translate(b.x, b.y)
}

func (b *Bullet) Init(myState *State) {
	b.id = getNextID()
	b.state = myState
	b.speed = 1000
	var err error
	b.img, _, err = ebitenutil.NewImageFromFile("internal/assets/bullet.png")
	if err != nil {
		log.Fatal(err)
	}
	b.scale = 2
	b.op = &ebiten.DrawImageOptions{}
	b.op.GeoM.Reset()
	b.op.GeoM.Scale(b.scale, b.scale)
	b.vx = 0
	b.vy = -1.0
}

func (b *Bullet) Update() {
	b.x += b.vx * DT * b.speed
	b.y += b.vy * DT * b.speed
	if b.y < 0 {
		_ = b.state.RemoveEntity(b)
	}
	w := 32.0
	h := 32.0
	found := false
	b.state.All(func(e IEntity) bool {
		if e != b && e.Type() != b.Type() {
			if b.x > e.X() && b.x < e.X()+w && b.y > e.Y() && b.y < e.Y()+h {
				_ = b.state.RemoveEntity(e)
				found = true
				return false
			}
		}
		return true
	})

	if found {
		_ = b.state.RemoveEntity(b)
	}
}

func (b *Bullet) Type() string {
	return "b"
}
