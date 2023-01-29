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
	Root       gui.Container
	PlayButton widgets.GuiButton
}

type EditorUI struct {
	Root gui.Container
}
