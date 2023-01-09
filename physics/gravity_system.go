package physics

import (
	"github.com/EngoEngine/ecs"
	"github.com/umi-l/open-mario-maker/geometry"
)

const G = 0.1

type gravityStorage struct {
	*ecs.BasicEntity
	*geometry.Transform
	*Velocity
}

type GravitySystem struct {
	Entities []forcesStorage
}

func (a *GravitySystem) Add(basic *ecs.BasicEntity, trans *geometry.Transform, velo *Velocity) {
	a.Entities = append(a.Entities, forcesStorage{basic, trans, velo})
}

func (a *GravitySystem) Remove(basic ecs.BasicEntity) {
	var del int = -1
	for index, entity := range a.Entities {
		if entity.ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		a.Entities = append(a.Entities[:del], a.Entities[del+1:]...)
	}
}

func (a GravitySystem) Update(dt float32) {
	for _, entity := range a.Entities {
		entity.Velocity.Y += G
	}
}
