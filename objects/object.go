package objects

import "github.com/umi-l/open-mario-maker/physics"

type Object struct {
	Colliders []physics.Collider //array of property enum used to add to systems.

	//position in tile space
	X float32
	Y float32

	ZIndex int
}
