package entity

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/usysrc/ebiten-boilerplate/internal/game/component"
)

func CreateBullet(cm *component.Manager) component.Entity {
	e := cm.NewEntity()

	cm.Positions[e] = &component.Position{X: 100, Y: 100}
	cm.Velocities[e] = &component.Velocity{}

	img, _, err := ebitenutil.NewImageFromFile("internal/assets/bullet.png")
	if err != nil {
		log.Fatal(err)
	}

	cm.Renders[e] = &component.Render{Image: img, Scale: 4, Z: 10}
	cm.Tags[e] = &component.Tag{Name: "Bullet"}

	return e
}
