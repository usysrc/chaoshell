package game

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type vector struct {
	x float64
	y float64
}

type component interface {
	onInit(*State) error
	onUpdate() error
	onDraw(*ebiten.Image) error
}

type element struct {
	position   vector
	components []component
}

func (e *element) addComponent(new component) {
	e.components = append(e.components, new)
}

func (e *element) getComponent(get component) *component {
	searchFor := reflect.TypeOf(get)
	for _, component := range e.components {
		if searchFor == reflect.TypeOf(component) {
			return &component
		}
	}
	return nil
}

func (e *element) init(state *State) {
	for _, component := range e.components {
		err := component.onInit(state)
		if err != nil {
			panic(err)
		}
	}
}

func (e *element) update() {
	for _, component := range e.components {
		err := component.onUpdate()
		if err != nil {
			panic(err)
		}
	}
}

func (e *element) draw(screen *ebiten.Image) {
	for _, component := range e.components {
		err := component.onDraw(screen)
		if err != nil {
			panic(err)
		}
	}
}
