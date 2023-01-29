package gameui

import (
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/widgets"
)

type GameUI struct {
	Root     gui.Container
	MainMenu MainMenuUI
}

type MainMenuUI struct {
	PlayButton widgets.GuiButton
}
