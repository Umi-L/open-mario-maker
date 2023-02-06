package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
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

func (tile *SolidTile) GetObject() *Object {
	return &tile.Object
}

func (tile SolidTile) Update(dt float32) {
	//TODO implement me
	panic("implement me")
}

func (tile SolidTile) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (tile *SolidTile) Clone() ObjectInterface {
	u := *tile
	return &u
}

func (tile SolidTile) Draw(scale float32) drawstack.DrawCall {
	return func(screen *ebiten.Image) {

		op := ebiten.DrawImageOptions{}

		op.GeoM.Translate(float64(tile.ScreenPos.X), float64(tile.ScreenPos.Y))

		screen.DrawImage(tile.sprite.Image, &op)
	}
}

func NewSolidTile(sprite waloader.Sprite) *SolidTile {
	return &SolidTile{sprite: sprite}
}
