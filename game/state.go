package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct {
	entities []IEntity
}

func (s *State) Init() {
	s.entities = make([]IEntity, 0)
}

// return true means that iteration was complete, false means that iteration ended prematurely
// the callback function returns false then iteration is ended immediatly
func (s State) All(fn func(s IEntity) bool) bool {
	// backwards because we might need to remove objects while iterating
	for i := len(s.entities) - 1; i >= 0; i-- {
		// we need to check, because length might change during iteration if we remove an item
		if i < len(s.entities) {
			ret := fn(s.entities[i])
			if ret == false {
				return false
			}
		}
	}
	return true
}

func (s *State) Update() {
	s.All(func(e IEntity) bool {
		e.Update()
		return true
	})
}

func (s *State) Draw(screen *ebiten.Image) {
	for _, entity := range s.entities {
		entity.Draw(screen)
	}
}

func (s *State) AddEntity(entity IEntity) {
	s.entities = append(s.entities, entity)
}

func (s *State) RemoveEntity(entity IEntity) {
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
