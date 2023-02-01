package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/drawstack"
	gameui "github.com/umi-l/open-mario-maker/game_ui"
	"github.com/umi-l/open-mario-maker/types"
	"github.com/umi-l/waloader"
	"github.com/wfranczyk/ento"
	"image/color"
)

type Game struct {
	Gui   gameui.GameUI
	Atlas map[string]waloader.Sprite
	World *ento.World
	State GameState

	SelectedObject int

	DeltaTime *types.DeltaTime
	Screen    *types.Screen

	drawStack drawstack.DrawStack
}

// Update mainloop
func (game *Game) Update() error {
	RunStateMachine(game.State)

	game.Gui.Root.Update()

	//TODO: Make dt calculated per frame instead of assuming 60 fps
	game.DeltaTime.Dt = 0.01666666666

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

	game.Screen.Screen = screen

	// draw background color
	screen.Fill(color.RGBA{
		R: 97,
		G: 133,
		B: 251,
		A: 255,
	})

	//tiled.DrawMap(screen, testmap)

	game.World.Update()

	game.drawStack.Draw(screen)

	game.DrawUi(screen)
}

// Layout internal resolution
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 426, 240
}
