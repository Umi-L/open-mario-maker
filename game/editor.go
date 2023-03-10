package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/objects"
	"image/color"
	"math"
)

type Grid struct {
	Size        int
	DefaultSize int
	Offset      geometry.Point
}

type Editor struct {
	game *Game

	Objects []objects.ObjectInterface

	Grid Grid

	leftMouseDown bool

	KeysDown map[string]bool

	MousePos          geometry.Point
	LastFrameMousePos geometry.Point

	ScrollAmount      geometry.Point
	ScrollSensitivity float64
}

func (e *Editor) Update() {
	e.UpdateInputs()

	if e.leftMouseDown && e.IsKeyPressed("Shift") { // if shift is pressed and left mouse is down, scroll the grid
		e.Grid.Offset = e.Grid.Offset.Add(e.MousePos.Sub(e.LastFrameMousePos))
		fmt.Printf("scrolling to %+v \n", e.Grid.Offset)
	} else if e.leftMouseDown { // if just clicking add object to the grid
		newTile := objects.MakeObjectFromId(e.game.SelectedObject)

		tileObject := newTile.GetObject()

		tileObject.Pos = e.FindScreenPositionOnGrid(e.MousePos)

		e.Objects = append(e.Objects, newTile)
	} else if e.ScrollAmount.Y != 0 { // if scrolling scale the grid
		e.Grid.Size = int(e.ScrollAmount.Y) + e.Grid.DefaultSize

		if e.Grid.Size <= 0 {
			e.Grid.Size = 1
		}
	}

	for _, o := range e.Objects {
		o.GetObject().ScreenPos = e.GridSpacePointToScreenSpace(o.GetObject().Pos)

		// find scale that makes the object fit in the grid
		scale := float32(e.Grid.Size) / float32(e.Grid.DefaultSize)

		e.game.drawStack.Add(o.Draw(scale), o.GetObject().ZIndex)
	}

	// debug draw text of grid size with draw stack
	e.game.drawStack.Add(func(screen *ebiten.Image) {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("grid size: %d", e.Grid.Size))
	}, 0)
	//change last frame
	e.LastFrameMousePos = e.MousePos
}

func (e *Editor) FindScreenPositionOnGrid(point geometry.Point) geometry.Point {
	//get point in screen space and then convert to grid space

	//get point in screen space
	point = point.Sub(e.Grid.Offset)

	//convert to grid space
	point = point.DivF(float64(e.Grid.Size))

	//round to the nearest grid point
	point = geometry.Point{
		X: math.Round(point.X),
		Y: math.Round(point.Y),
	}

	return point
}

func (e *Editor) GridSpacePointToScreenSpace(point geometry.Point) geometry.Point {
	//convert to screen space
	point = point.MulF(float64(e.Grid.DefaultSize))

	//add offset
	point = point.Add(e.Grid.Offset)

	return point
}

func (e *Editor) DrawGrid() {
	// get screen size
	screenWidth, screenHeight := e.game.Screen.Screen.Size()

	// calculate the offset of a single grid square
	offset := geometry.Point{
		X: math.Mod(float64(e.Grid.Offset.X), float64(e.Grid.Size)),
		Y: math.Mod(float64(e.Grid.Offset.Y), float64(e.Grid.Size)),
	}

	// draw grid using vertical and horizontal lines using draw stack accounting for scale & offset
	e.game.drawStack.Add(func(screen *ebiten.Image) {

		for i := 0; i < screenWidth; i += e.Grid.Size {
			ebitenutil.DrawLine(screen, float64(i)+offset.X, offset.Y, float64(i)+offset.X, float64(screenHeight)+offset.Y, color.RGBA{255, 255, 255, 100})
		}

		for i := 0; i < screenHeight; i += e.Grid.Size {
			ebitenutil.DrawLine(screen, offset.X, float64(i)+offset.Y, float64(screenWidth)+offset.X, float64(i)+offset.Y, color.RGBA{255, 255, 255, 100})
		}
	}, 0)
}

func (e *Editor) UpdateInputs() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		e.leftMouseDown = true
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		e.leftMouseDown = false
	}

	mousePosX, mousePosY := ebiten.CursorPosition()

	e.MousePos = geometry.Point{X: float64(mousePosX), Y: float64(mousePosY)}

	wheelX, wheelY := ebiten.Wheel()

	e.ScrollAmount = e.ScrollAmount.Add(geometry.Point{X: wheelX * e.ScrollSensitivity, Y: wheelY * e.ScrollSensitivity})

	var pressed []ebiten.Key
	pressed = inpututil.AppendPressedKeys(pressed)

	e.KeysDown = make(map[string]bool)

	for _, key := range pressed {
		e.KeysDown[key.String()] = true
	}
}

func (e *Editor) IsKeyPressed(key string) bool {
	value, exists := e.KeysDown[key]

	if exists {
		return value
	}
	return false
}

func MakeEditor(game *Game) Editor {
	var editor = Editor{game: game}
	mousePosX, mousePosY := ebiten.CursorPosition()
	editor.LastFrameMousePos = geometry.Point{X: float64(mousePosX), Y: float64(mousePosY)}
	editor.Grid.Size = 16
	editor.Grid.DefaultSize = 16
	editor.ScrollSensitivity = 0.1
	return editor
}
