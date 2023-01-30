package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/types"
	"github.com/umi-l/yosui-ui"
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/widgets"
)

func (game *Game) initUI() {
	game.Gui.Root = yosui.MakeRootContainer(game.Layout(0, 0))

	//--main menu--
	game.Gui.MainMenu.Trunk = gui.NewContainer(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, true)

	trans := gui.MakeTransformWithImage(playButtonImage, gui.OriginCenter).SetXPercent(0.5).SetYPercent(0.5)

	//play button
	game.Gui.MainMenu.PlayButton = widgets.NewButton(playButtonImage, trans, func() {
		game.State = types.Editing
		game.Gui.MainMenu.Trunk.Visible = false
		game.Gui.Editor.Trunk.Visible = true

		fmt.Print("Play Button Pressed\n")
	})

	//add to main menu
	game.Gui.MainMenu.Trunk.AddChild(&game.Gui.MainMenu.PlayButton)

	//add to gui
	game.Gui.Root.AddChild(&game.Gui.MainMenu.Trunk)

	//--Editor--
	game.Gui.Editor.Trunk = gui.NewContainer(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, false)

	// make top panel
	game.Gui.Editor.TopPanel = widgets.MakePanel(topPanelImage)
	game.Gui.Editor.TopPanel.SetTransform(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 0.1})

	// add to Editor
	game.Gui.Editor.Trunk.AddChild(&game.Gui.Editor.TopPanel)

	// add Root to gui
	game.Gui.Root.AddChild(&game.Gui.Editor.Trunk)
}

func (game *Game) drawUi(screen *ebiten.Image) {
	//resize event
	w, h := screen.Size()
	game.Gui.Root.UpdateTransform(gui.Transform{X: 0, Y: 0, W: float32(w), H: float32(h)})

	//draw GUI
	game.Gui.Root.DrawAsRoot(screen)
}
