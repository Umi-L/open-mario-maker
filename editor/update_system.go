package editor

import (
	"github.com/EngoEngine/ecs"
	"github.com/umi-l/open-mario-maker/objects"
)

type editorUpdateStorage struct {
	*ecs.BasicEntity
	*objects.Object
}

type UpdateSystem struct {
	Entities []editorUpdateStorage
}

func (updateSystem *UpdateSystem) Add(basic *ecs.BasicEntity, object *objects.Object) {
	updateSystem.Entities = append(updateSystem.Entities, editorUpdateStorage{basic, object})
}

func (updateSystem *UpdateSystem) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range updateSystem.Entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		updateSystem.Entities = append(updateSystem.Entities[:delete], updateSystem.Entities[delete+1:]...)
	}
}

func (updateSystem UpdateSystem) Update(dt float32) {

}
