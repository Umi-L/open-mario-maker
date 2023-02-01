package objects

import (
	"github.com/umi-l/waloader"
)

var objects []ObjectInterface

func Init(atlas map[string]waloader.Sprite) {
	register(NewSolidTile(atlas["MarioPlayButton"]))
}

func register(object ObjectInterface) {
	objects = append(objects, object)
}
