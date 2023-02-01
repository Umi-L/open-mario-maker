package objects

func MakeObjectFromId(id int) ObjectInterface {
	return objects[id].Clone()
}
