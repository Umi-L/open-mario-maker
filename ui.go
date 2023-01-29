package main

import (
	"github.com/umi-l/yosui-ui"
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/widgets"
)

func (game *Game) initUI() {
	//load assets
	playButtonImage = game.Atlas["MarioPlayButton"].Image

	//fmt.Printf("PlayButton: %+v\n", game.Atlas["MarioPlayButton"])

	game.Gui.Root = yosui.MakeRootContainer(game.Layout(0, 0))

	//--main menu--
	mainMenu := gui.NewContainer(gui.Transform{X: 0, Y: 0, WPercent: 1, HPercent: 1}, true)

	trans := gui.MakeTransformWithImage(playButtonImage, gui.OriginCenter)

	trans.XPercent = 0.5
	trans.YPercent = 0.5

	//play button
	playButton := widgets.NewButton(playButtonImage, trans)

	//add to main menu
	mainMenu.AddChild(&playButton)

	//add to gui
	game.Gui.Root.AddChild(&mainMenu)
}
