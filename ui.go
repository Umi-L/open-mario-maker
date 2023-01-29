package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/yosui-ui"
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/widgets"

	. "github.com/umi-l/open-mario-maker/types"
)

func (game *Game) initUI() {
	//load assets
	playButtonImage = game.Atlas["MarioPlayButton"].Image

	//fmt.Printf("PlayButton: %+v\n", game.Atlas["MarioPlayButton"])

	game.Gui.Root = yosui.MakeRootContainer(game.Layout(0, 0))

	//--main menu--
	game.Gui.MainMenu.Root = gui.NewContainer(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, true)

	trans := gui.MakeTransformWithImage(playButtonImage, gui.OriginCenter)

	trans.XPercent = 0.5
	trans.YPercent = 0.5

	//play button
	game.Gui.MainMenu.PlayButton = widgets.NewButton(playButtonImage, trans)

	//add to main menu
	game.Gui.MainMenu.Root.AddChild(&game.Gui.MainMenu.PlayButton)

	//add to gui
	game.Gui.Root.AddChild(&game.Gui.MainMenu.Root)
}

func (game *Game) drawUi(screen *ebiten.Image) {
	//resize event
	w, h := screen.Size()
	game.Gui.Root.SetTransform(gui.Transform{X: 0, Y: 0, W: float32(w), H: float32(h)})

	//on play button click
	if game.Gui.MainMenu.PlayButton.IsPressed() && game.Gui.MainMenu.PlayButton.IsVisible() {
		game.State = Editing
		game.Gui.MainMenu.Root.Visible = false
		game.Gui.Editor.Root.Visible = true

		fmt.Print("Play Button Pressed\n")
	}

	//draw GUI
	game.Gui.Root.DrawTree(screen)
}
