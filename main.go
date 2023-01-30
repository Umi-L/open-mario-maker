package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	gameui "github.com/umi-l/open-mario-maker/game_ui"
	"github.com/umi-l/open-mario-maker/gamestates"
	"github.com/umi-l/open-mario-maker/tiled"
	"github.com/umi-l/waloader"
	"image/color"
	_ "image/png"
	"log"
)

var testmap tiled.Map

//go:embed resources
var res embed.FS

// init
func (game *Game) init() {

	//map def
	mapdata, err := res.ReadFile("resources/testmap2.json")

	if err != nil {
		log.Fatal(err)
	}

	testmap = tiled.ParseJson(mapdata, tilemapSheet)
}

type Game struct {
	Gui   gameui.GameUI
	Atlas map[string]waloader.Sprite
	State gamestates.GameState
}

// Update mainloop
func (game *Game) Update() error {
	game.Gui.Root.Update()

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {

	// draw background color
	screen.Fill(color.RGBA{
		R: 97,
		G: 133,
		B: 251,
		A: 255,
	})

	tiled.DrawMap(screen, testmap)

	game.drawUi(screen)
}

// Layout internal resolution
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 426, 240
}

// entrypoint
func main() {

	//external resolution
	//ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowSize(1280, 720)

	//title
	ebiten.SetWindowTitle("Open Mario Maker")

	//define empty game
	game := Game{}

	//init all
	game.InitAssets()
	//game.init()
	game.initECS()
	game.initUI()

	//run game and handle errors
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
