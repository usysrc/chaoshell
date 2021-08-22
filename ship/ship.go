package ship

import (
	// color "image/color"
	_ "image/png"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct{
	x, y float64
	img *ebiten.Image
	op *ebiten.DrawImageOptions

}

func (s *Ship) SetPos (x, y float64) {
	s.x, s.y = x, y
	s.op.GeoM.Translate(s.x, s.y)
}

func (s *Ship) Init() {
	var err error
	s.img, _, err = ebitenutil.NewImageFromFile("ship.png")
	if err != nil {
		log.Fatal(err)
	}
	s.op = &ebiten.DrawImageOptions{}
	s.op.GeoM.Scale(4, 4)
}

func (s *Ship) Update() {
	s.x, s.y = 0, 0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.x += 1
	}
	s.op.GeoM.Translate(s.x,s.y)
}

func (s *Ship) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.img, s.op)
}