package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/physics"
	"github.com/umi-l/waloader"
)

type Player struct {
	ecs.BasicEntity
	waloader.Animation
	geometry.Transform
	physics.Velocity
}
