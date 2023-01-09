package animation

import (
	"github.com/EngoEngine/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/geometry"
)

type animationUpdateStorage struct {
	*ecs.BasicEntity
	*Animation
	*geometry.Transform
}

type UpdateSystem struct {
	Entities []animationUpdateStorage
}

func (a *UpdateSystem) Add(basic *ecs.BasicEntity, anim *Animation, trans *geometry.Transform) {
	a.Entities = append(a.Entities, animationUpdateStorage{basic, anim, trans})
}

func (a *UpdateSystem) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range a.Entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		a.Entities = append(a.Entities[:delete], a.Entities[delete+1:]...)
	}
}

func (a UpdateSystem) Update(dt float32) {
	for _, entity := range a.Entities {
		entity.updateTimer(dt)
	}
}

func (a UpdateSystem) Draw(screen *ebiten.Image) {
	for _, entity := range a.Entities {
		entity.Draw(screen, entity.Transform.Position.X, entity.Transform.Position.Y, entity.Transform.Rotation)
	}
}
