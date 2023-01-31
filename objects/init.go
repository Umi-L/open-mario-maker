package objects

import (
	"github.com/umi-l/waloader"
)

var objects []ObjectInterface

func Init(atlas map[string]waloader.Sprite) {
	register(MakeSolidTile(atlas["SolidBrick"]))
}

func register(object ObjectInterface) {
	objects = append(objects, object)
}
