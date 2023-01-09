package loader

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

func LoadSheet(path string, cellWidth int, cellHeight int) Sheet {
	var err error
	var img *ebiten.Image

	img, _, err = ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()

	return Sheet{
		Path:        path,
		Texture:     img,
		CellWidth:   cellWidth,
		CellHeight:  cellHeight,
		Rows:        height / cellHeight,
		CellsPerRow: width / cellWidth,
	}
}
