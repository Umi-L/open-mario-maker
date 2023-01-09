package physics

import (
	"github.com/EngoEngine/ecs"
	"github.com/umi-l/open-mario-maker/geometry"
)

type forcesStorage struct {
	*ecs.BasicEntity
	*geometry.Transform
	*Velocity
}

type ForcesSystem struct {
	Entities []forcesStorage
}

func (a *ForcesSystem) Add(basic *ecs.BasicEntity, trans *geometry.Transform, velo *Velocity) {
	a.Entities = append(a.Entities, forcesStorage{basic, trans, velo})
}

func (a *ForcesSystem) Remove(basic ecs.BasicEntity) {
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

func (a ForcesSystem) Update(dt float32) {
	for _, entity := range a.Entities {
		entity.Transform.Position.X += entity.Velocity.X
		entity.Transform.Position.Y += entity.Velocity.Y
	}
}
