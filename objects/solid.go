package objects

import (
	"github.com/umi-l/open-mario-maker/drawstack"
	"github.com/umi-l/open-mario-maker/physics"
	"github.com/umi-l/waloader"
)

type SolidTile struct {
	sprite waloader.Sprite

	Object
}

func (tile SolidTile) IsColliding(collider physics.Collider) bool {
	//TODO implement me
	panic("implement me")
}

func (tile SolidTile) OnCollision(other ObjectInterface) {
	//TODO implement me
	panic("implement me")
}

func (tile SolidTile) GetObject() Object {
	//TODO implement me
	panic("implement me")
}

func (tile SolidTile) Update(dt float32) {
	//TODO implement me
	panic("implement me")
}

func (tile SolidTile) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (tile SolidTile) Draw() drawstack.DrawCall {
	//TODO implement me
	panic("implement me")
}

func MakeSolidTile(sprite waloader.Sprite) SolidTile {
	return SolidTile{sprite: sprite}
}
