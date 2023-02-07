package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/objects"
	"math"
)

type Grid struct {
	Size        int
	DefaultSize int
	Offset      geometry.Point
}

func (g Grid) GetSizeValue() float64 {
	return float64(g.Size) / float64(g.DefaultSize)
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

	if e.leftMouseDown && e.IsKeyPressed("Shift") {
		e.Grid.Offset = e.Grid.Offset.Add(e.MousePos.Sub(e.LastFrameMousePos))
		fmt.Printf("scrolling to %+v \n", e.Grid.Offset)
	} else if e.leftMouseDown {
		newTile := objects.MakeObjectFromId(e.game.SelectedObject)

		tileObject := newTile.GetObject()

		tileObject.Pos = e.FindScreenPositionOnGrid(e.MousePos)

		e.Objects = append(e.Objects, newTile)
	} else if e.ScrollAmount.Y != 0 {
		e.Grid.Size = int(e.ScrollAmount.Y) + e.Grid.DefaultSize

		if e.Grid.Size <= 0 {
			e.Grid.Size = 1
		}
	}

	for _, o := range e.Objects {
		e.SetObjectScreenPos(o)

		e.game.drawStack.Add(o.Draw(1/float32(e.Grid.GetSizeValue())), o.GetObject().ZIndex)
	}

	//change last frame
	e.LastFrameMousePos = e.MousePos
}

func (e *Editor) FindScreenPositionOnGrid(point geometry.Point) geometry.Point {
	// tile position is mouse position (accounting for offset) / tile size floored.

	gridPos := geometry.Point{
		X: math.Floor((point.X/(1/e.Grid.GetSizeValue()) - e.Grid.Offset.X) / float64(e.Grid.Size)),
		Y: math.Floor((point.Y/(1/e.Grid.GetSizeValue()) - e.Grid.Offset.Y) / float64(e.Grid.Size)),
	}

	fmt.Printf("gridPos: %+v \n mousePos: %+v \n", gridPos, e.MousePos)

	return gridPos
}

func (e *Editor) SetObjectScreenPos(oInterface objects.ObjectInterface) {
	object := oInterface.GetObject()

	object.ScreenPos = geometry.Point{
		X: object.Pos.X*float64(e.Grid.Size) + e.Grid.Offset.X,
		Y: object.Pos.Y*float64(e.Grid.Size) + e.Grid.Offset.Y,
	}
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
