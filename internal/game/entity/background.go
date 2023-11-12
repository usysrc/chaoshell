package entity

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/usysrc/chaoshell/internal/game/component"
)

func CreateBackground(cm *component.Manager) component.Entity {
	e := component.NewEntity()
	cm.Positions[e] = &component.Position{X: 0, Y: 0}

	img, _, err := ebitenutil.NewImageFromFile("internal/assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	cm.Renders[e] = &component.Render{Image: img, Scale: 1, Z: -1000}
	return e

}
