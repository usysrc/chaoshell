package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct {
	entities []Entity
}

func (s *State) Init() {
	s.entities = make([]Entity, 0)
}

func (s *State) Update() {
	// backwards because we might need to remove objects while iterating
	for i := len(s.entities) - 1; i >= 0; i-- {
		// we need to check, because length might change during iteration
		if i < len(s.entities) {
			s.entities[i].Update()
		}
	}
}

func (s *State) Draw(screen *ebiten.Image) {
	for _, entity := range s.entities {
		entity.Draw(screen)
	}
}

func (s *State) AddEntity(entity Entity) {
	s.entities = append(s.entities, entity)
}

func (s *State) RemoveEntity(entity Entity) {
	var index, found = 0, false
	for i, e := range s.entities {
		if e == entity {
			found = true
			index = i
			break
		}
	}
	if found {
		s.entities = append(s.entities[:index], s.entities[index+1:]...)
	}
}
