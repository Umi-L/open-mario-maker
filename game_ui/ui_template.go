package gameui

import (
	"github.com/umi-l/yosui-ui/gui"
	"github.com/umi-l/yosui-ui/widgets"
)

type GameUI struct {
	Root     gui.Container
	MainMenu MainMenuUI
	Editor   EditorUI
}

type MainMenuUI struct {
	Trunk      gui.Container
	PlayButton widgets.GuiButton
}

type EditorUI struct {
	Trunk      gui.Container
	TopPanel   widgets.Panel
	TestButton widgets.GuiButton
}
