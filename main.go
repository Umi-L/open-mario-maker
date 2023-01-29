package main

import (
	"image/color"
	_ "image/png"
	"log"
	"time"

	"embed"
	"github.com/EngoEngine/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/open-mario-maker/animation"
	"github.com/umi-l/open-mario-maker/entities"
	gameui "github.com/umi-l/open-mario-maker/game_ui"
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/physics"
	"github.com/umi-l/open-mario-maker/tiled"
	. "github.com/umi-l/open-mario-maker/types"
	"github.com/umi-l/open-mario-maker/utils"
	"github.com/umi-l/waloader"
)

var world ecs.World

var player entities.Player

var dt func() time.Duration = utils.GetDt()

var systems map[SystemIndex]interface{}

var testmap tiled.Map

//go:embed resources
var res embed.FS

// init
func (game *Game) init() {
	//debug entities
	player = entities.Player{BasicEntity: ecs.NewBasic(), Animation: marioSwimmingAnimation, Transform: geometry.NewEmptyTransform(), Velocity: physics.NewEmptyVelocity()}

	//map def
	mapdata, err := res.ReadFile("resources/testmap2.json")

	if err != nil {
		log.Fatal(err)
	}

	testmap = tiled.ParseJson(mapdata, tilemapSheet)

	//Loop over all Systems
	for _, system := range world.Systems() {

		// Use a type-switch to figure out which System is which
		switch sys := system.(type) {

		// Create a case for each System you want to use
		case *animation.UpdateSystem:
			sys.Add(&player.BasicEntity, &player.Animation, &player.Transform)
		case *physics.ForcesSystem:
			sys.Add(&player.BasicEntity, &player.Transform, &player.Velocity)
		case *physics.GravitySystem:
			sys.Add(&player.BasicEntity, &player.Transform, &player.Velocity)
		}
	}
}

type Game struct {
	Gui   gameui.GameUI
	Atlas map[string]waloader.Sprite
	State GameState
}

// mainloop
func (game *Game) Update() error {
	world.Update(float32(dt().Seconds()))

	game.Gui.Root.Update()

	return nil
}

// draw
func (game *Game) Draw(screen *ebiten.Image) {

	// draw background color
	screen.Fill(color.RGBA{
		R: 97,
		G: 133,
		B: 251,
		A: 255,
	})

	//get animation system and run draw
	systems[AnimationUpdate].(*animation.UpdateSystem).Draw(screen)

	tiled.DrawMap(screen, testmap)

	game.drawUi(screen)
}

// internal resolution
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
