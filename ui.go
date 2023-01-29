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
	game.Gui.Root = yosui.MakeRootContainer(game.Layout(0, 0))

	//--main menu--
	game.Gui.MainMenu.Root = gui.NewContainer(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, true)

	trans := gui.MakeTransformWithImage(playButtonImage, gui.OriginCenter).SetXPercent(0.5).SetYPercent(0.5)

	//play button
	game.Gui.MainMenu.PlayButton = widgets.NewButton(playButtonImage, trans)

	//add to main menu
	game.Gui.MainMenu.Root.AddChild(&game.Gui.MainMenu.PlayButton)

	//add to gui
	game.Gui.Root.AddChild(&game.Gui.MainMenu.Root)

	//--Editor--
	game.Gui.Editor.Root = gui.NewContainer(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, false)

	// make top panel
	game.Gui.Editor.TopPanel = gui.MakeElement(topPanelImage)
	game.Gui.Editor.TopPanel.SetTransform(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 0.1})

	// add to Editor
	game.Gui.Editor.Root.AddChild(&game.Gui.Editor.TopPanel)

	// add Root to gui
	game.Gui.Root.AddChild(&game.Gui.Editor.Root)
}

func (game *Game) drawUi(screen *ebiten.Image) {
	//resize event
	w, h := screen.Size()
	game.Gui.Root.UpdateTransform(gui.Transform{X: 0, Y: 0, W: float32(w), H: float32(h)})

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
