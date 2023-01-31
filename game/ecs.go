package game

import (
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/objects"
	"github.com/umi-l/open-mario-maker/physics"
	"github.com/umi-l/open-mario-maker/systems"
	"github.com/umi-l/open-mario-maker/types"
	"github.com/umi-l/waloader"
	"github.com/wfranczyk/ento"
)

func (game *Game) InitECS() {
	// Create the world and register components
	game.World = ento.NewWorldBuilder().
		// Use "zero-values" as values are ignored when registering
		WithSparseComponents(physics.Velocity{}, objects.Object{}, geometry.Transform{}, waloader.Animation{}).
		WithSingletonComponents(types.DeltaTime{}, types.Screen{}).
		// Pre-allocate space for 256 entities (world can grow beyond that automatically)
		Build(256)

	// Add systems
	system := &systems.AnimationUpdateSystem{}
	game.World.AddSystems(system)

	var deltaTime *types.DeltaTime
	var screen *types.Screen

	singletons := game.World.AddEntity(types.DeltaTime{}, types.Screen{})

	singletons.Get(&deltaTime)
	singletons.Get(&screen)

	game.DeltaTime = deltaTime
	game.Screen = screen
}
