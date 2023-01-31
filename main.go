package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/game"
	_ "image/png"
	"log"
)

// entrypoint
func main() {

	//external resolution
	//ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowSize(1280, 720)

	//title
	ebiten.SetWindowTitle("Open Mario Maker")

	//define empty game
	game := game.Game{}

	//init all
	game.InitAssets()
	//game.init()
	game.InitECS()
	game.InitUI()

	//run game and handle errors
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
