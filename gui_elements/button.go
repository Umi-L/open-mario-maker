package gui_elements

import (
	"github.com/umi-l/open-mario-maker/gui"
	"github.com/umi-l/open-mario-maker/gui_update_params"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GuiButton struct {
	gui.Element

	OnClick func(gui_update_params.ButtonUpdateParams)
}

func (b GuiButton) Update(updateParams gui_update_params.UpdateParams) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		var x, y = ebiten.CursorPosition()
		if b.Rect.Contains((float32)(x), (float32)(y)) {
			b.OnClick(updateParams.ButtonUpdateParams)
		}
	}
}

func NewButton(image *ebiten.Image, onClick func(gui_update_params.ButtonUpdateParams)) GuiButton {
	return GuiButton{
		Element: gui.MakeElement(image),
		OnClick: onClick,
	}
}
