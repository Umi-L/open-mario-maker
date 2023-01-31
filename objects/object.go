package objects

import "github.com/umi-l/open-mario-maker/physics"

type Object struct {
	Colliders []physics.Collider //array of property enum used to add to systems.

	//position in tile space
	x float32
	y float32
}
