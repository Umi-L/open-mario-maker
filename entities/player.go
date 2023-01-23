package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/umi-l/open-mario-maker/animation"
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/physics"
)

type Player struct {
	ecs.BasicEntity
	animation.Animation
	geometry.Transform
	physics.Velocity
}
