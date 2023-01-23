package gameui

import (
	"github.com/umi-l/open-mario-maker/gui"
	"github.com/umi-l/open-mario-maker/gui_elements"
)

type GameUI struct {
	Root     gui.Container
	MainMenu MainMenuUI
}

type MainMenuUI struct {
	PlayButton gui_elements.GuiButton
}
