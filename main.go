package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/game"
	"github.com/umi-l/open-mario-maker/objects"
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
	Game := game.Game{}

	//init all
	Game.InitAssets()
	Game.InitECS()
	Game.InitUI()
	game.InitStateMachine(&Game)
	objects.Init(Game.Atlas)

	//run game and handle errors
	if err := ebiten.RunGame(&Game); err != nil {
		log.Fatal(err)
	}
}
