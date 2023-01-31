package objects

import (
	"github.com/umi-l/open-mario-maker/drawstack"
	"github.com/umi-l/open-mario-maker/physics"
)

type ObjectInterface interface {
	IsColliding(collider physics.Collider) bool
	OnCollision(other ObjectInterface)
	GetObject() Object
	Update(dt float32)
	Destroy()
	Draw() drawstack.DrawCall
}
