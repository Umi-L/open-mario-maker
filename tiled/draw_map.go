package tiled

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawMap(screen *ebiten.Image, m Map){

	mapX := 0
	mapY := 0

	for _, layer := range m.TileLayers{
		for _, value := range layer.TileData {

			if value != 0{
		
				tileLinePos := value - 1

				tileY := tileLinePos / (m.Sheet.CellsPerRow)
				tileX := tileLinePos - tileY * m.Sheet.CellsPerRow

				tile := m.Sheet.Texture.SubImage(image.Rect(
					tileX * m.Sheet.CellWidth,
					tileY * m.Sheet.CellHeight,
					tileX * m.Sheet.CellWidth + m.Sheet.CellWidth,
					tileY * m.Sheet.CellHeight + m.Sheet.CellHeight,
				)).(*ebiten.Image)

				op := &ebiten.DrawImageOptions{}

				op.GeoM.Translate(float64(m.Sheet.CellWidth*mapX), float64(m.Sheet.CellHeight*mapY))

				screen.DrawImage(tile, op)
			}
			
			mapX += 1
			if mapX > m.Sheet.CellsPerRow{
				mapX = 0
				mapY += 1
			}
		}
	}
}