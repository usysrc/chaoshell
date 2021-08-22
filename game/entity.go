package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface{
	Update()
	Draw(screen *ebiten.Image)
}