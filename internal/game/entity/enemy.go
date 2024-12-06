package entity

import (
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/usysrc/ebiten-boilerplate/internal/game/component"
)

func CreateEnemy(cm *component.Manager) component.Entity {
	e := cm.NewEntity()

	cm.Positions[e] = &component.Position{X: rand.Float64() * 800.0, Y: -32}

	img, _, err := ebitenutil.NewImageFromFile("internal/assets/enemy.png")
	if err != nil {
		log.Fatal(err)
	}
	cm.Velocities[e] = &component.Velocity{X: 0, Y: 1}
	cm.Renders[e] = &component.Render{Image: img, Scale: 1}
	cm.Tags[e] = &component.Tag{Name: "Enemy"}
	return e
}
