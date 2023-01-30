package main

import (
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/objects"
	"github.com/umi-l/open-mario-maker/physics"
	"github.com/umi-l/open-mario-maker/systems"
	"github.com/umi-l/open-mario-maker/types"
	"github.com/wfranczyk/ento"
)

func (game *Game) initECS() {
	// Create the world and register components
	world := ento.NewWorldBuilder().
		// Use "zero-values" as values are ignored when registering
		WithSparseComponents(physics.Velocity{}, objects.Object{}, geometry.Transform{}).
		WithSingletonComponents(types.DeltaTime{}, types.Screen{}).
		// Pre-allocate space for 256 entities (world can grow beyond that automatically)
		Build(256)

	// Add systems
	system := &systems.AnimationUpdateSystem{}
	world.AddSystems(system)
}
