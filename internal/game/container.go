package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
)

type Container struct {
	entities map[int]Entifier
}

func (c *Container) Init() {
	c.entities = make(map[int]Entifier)
}

func (c Container) All(fn func(e Entifier) bool) bool {
	// Create a temporary slice of keys to iterate over
	// This avoids modifying the map during iteration
	keys := make([]int, 0, len(c.entities))
	for key := range c.entities {
		keys = append(keys, key)
	}

	// Iterate over the entities using the keys slice
	for _, key := range keys {
		entity, exists := c.entities[key]
		if exists {
			// Now you can access Entity-specific properties and methods
			if !fn(entity) {
				return false // Terminate early if the callback returns false
			}

		}
	}

	return true
}

func (c *Container) Update() {
	c.All(func(e Entifier) bool {
		e.Update()
		return true
	})
}

func (c *Container) Draw(screen *ebiten.Image) {
	for _, entity := range c.entities {
		entity.Draw(screen)
	}
}

func (c *Container) AddEntity(entity Entifier) {
	c.entities[entity.ID()] = entity
}

func (c *Container) RemoveEntity(entity Entifier) error {
	_, exists := c.entities[entity.ID()]
	if !exists {
		return errors.New("entity not found")
	}
	delete(c.entities, entity.ID())
	return nil
}
