package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
)

type State struct {
	entities map[int]IEntity
}

func (s *State) Init() {
	s.entities = make(map[int]IEntity)
}

func (s State) All(fn func(e IEntity) bool) bool {
	// Create a temporary slice of keys to iterate over
	// This avoids modifying the map during iteration
	keys := make([]int, 0, len(s.entities))
	for key := range s.entities {
		keys = append(keys, key)
	}

	// Iterate over the entities using the keys slice
	for _, key := range keys {
		entity, exists := s.entities[key]
		if exists {
			if !fn(entity) {
				return false // Terminate early if the callback returns false
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
	s.entities[entity.ID()] = entity
}

func (s *State) RemoveEntity(entity IEntity) error {
	_, exists := s.entities[entity.ID()]
	if !exists {
		return errors.New("entity not found")
	}
	delete(s.entities, entity.ID())
	return nil
}
