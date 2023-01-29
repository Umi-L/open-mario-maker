package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/umi-l/open-mario-maker/animation"
	"github.com/umi-l/open-mario-maker/physics"

	. "github.com/umi-l/open-mario-maker/types"
)

func (game *Game) initECS() {
	//systems map
	systems = make(map[SystemIndex]interface{})

	//ecs
	world = ecs.World{}

	//systems
	animationUpdateSystem := &animation.UpdateSystem{}
	world.AddSystem(animationUpdateSystem)
	systems[AnimationUpdate] = animationUpdateSystem

	ForcesSystem := &physics.ForcesSystem{}
	world.AddSystem(ForcesSystem)
	systems[Forces] = ForcesSystem

	gravitySystem := &physics.GravitySystem{}
	world.AddSystem(gravitySystem)
	systems[Gravity] = gravitySystem

}
