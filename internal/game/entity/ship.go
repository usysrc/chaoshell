package entity

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/usysrc/chaoshell/internal/game/component"
)

func CreateShip(cm *component.Manager) component.Entity {
	e := component.NewEntity()

	cm.Positions[e] = &component.Position{X: 100, Y: 100}
	cm.Velocities[e] = &component.Velocity{}

	img, _, err := ebitenutil.NewImageFromFile("internal/assets/ship.png")
	if err != nil {
		log.Fatal(err)
	}
	cm.Renders[e] = &component.Render{Image: img, Scale: 4}
	cm.Tags[e] = &component.Tag{Name: "Ship"}

	return e

}
