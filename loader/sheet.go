package loader

import "github.com/hajimehoshi/ebiten/v2"

type Sheet struct {
	Path        string
	Texture     *ebiten.Image
	CellWidth   int
	CellHeight  int
	Rows        int
	CellsPerRow int
}
