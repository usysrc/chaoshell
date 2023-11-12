package systems

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/usysrc/chaoshell/internal/game/component"
)

type RenderSystem struct {
	Components *component.Manager
}

type RenderData struct {
	entity component.Entity
	render *component.Render
}

func (r *RenderSystem) Draw(screen *ebiten.Image) {
	var renderData []RenderData
	for e, render := range r.Components.Renders {
		renderData = append(renderData, RenderData{
			entity: e,
			render: render,
		})

	}

	sort.Slice(renderData, func(i, j int) bool {
		return renderData[i].render.Z < renderData[j].render.Z
	})

	for _, data := range renderData {
		if pos, ok := r.Components.Positions[data.entity]; ok {
			// Draw the entity's image at its position
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Reset()
			options.GeoM.Scale(data.render.Scale, data.render.Scale)
			options.GeoM.Translate(pos.X, pos.Y)
			screen.DrawImage(data.render.Image, options)
		}
	}
}
