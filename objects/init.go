package objects

import (
	"github.com/umi-l/waloader"
)

var objects []ObjectInterface

func register(object ObjectInterface) {
	objects = append(objects, object)
}

func Init(atlas map[string]waloader.Sprite) {
	register(NewSolidTile(atlas["Tiles/SolidBrick"]))
	register(NewSolidTile(atlas["Tiles/Brick"]))
}
