package objects

import (
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/physics"
)

type Object struct {
	Colliders []physics.Collider //array of property enum used to add to systems.

	//position in tile space
	Pos geometry.Point

	ScreenPos geometry.Point

	ZIndex int
}
