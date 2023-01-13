package main

import (
	_ "image/png"
	"log"
	"time"

	"github.com/umi-l/open-mario-maker/animation"
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/gui"
	"github.com/umi-l/open-mario-maker/loader"
	"github.com/umi-l/open-mario-maker/physics"
	"github.com/umi-l/open-mario-maker/tiled"
	. "github.com/umi-l/open-mario-maker/types"
	"github.com/umi-l/open-mario-maker/utils"

	"embed"

	"github.com/EngoEngine/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var world ecs.World

var player Player

var dt func() time.Duration = utils.GetDt()

var charSheet loader.Sheet
var tilemapSheet loader.Sheet

var playButtonImage *ebiten.Image

var marioSwimmingAnimation animation.Animation

type Player struct {
	ecs.BasicEntity
	animation.Animation
	geometry.Transform
	physics.Velocity
}

var systems map[SystemIndex]interface{}

var testmap tiled.Map

//go:embed resources
var res embed.FS

// init
func init() {

	//systems map
	systems = make(map[SystemIndex]interface{})

	//resources
	charSheet = loader.LoadSheet("MarioSpriteSheet.png", 16, 32)
	marioSwimmingAnimation = animation.Load(&charSheet, 0, 5, 0.1)

	tilemapSheet = loader.LoadSheet("resources/tilemap.png", 16, 16)

	var err error
	playButtonImage, _, err = ebitenutil.NewImageFromFile("resources/MarioPlayButton.png")

	if err != nil{
		log.Fatal(err)
	}

	//ecs
	world = ecs.World{}

	//systems
	animationUpdateSystem := &animation.UpdateSystem{}
	world.AddSystem(animationUpdateSystem)
	systems[AnimationUpdate] = animationUpdateSystem

	ForcesSystem := &physics.ForcesSystem{}
	world.AddSystem(ForcesSystem)
	systems[Forces] = ForcesSystem

	gravitySystem := &physics.GravitySystem{}
	world.AddSystem(gravitySystem)
	systems[Gravity] = gravitySystem

	//debug entities
	player = Player{ecs.NewBasic(), marioSwimmingAnimation, geometry.NewEmptyTransform(), physics.NewEmptyVelocity()}

	//map def
	mapdata, err := res.ReadFile("resources/testmap2.json")


	if err != nil{
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

type Game struct{
	Gui gui.Container
}

// mainloop
func (g *Game) Update() error {
	world.Update(float32(dt().Seconds()))

	return nil
}

// draw
func (g *Game) Draw(screen *ebiten.Image) {
	//player.Draw(screen, 0.0, 0.0, 0.0)

	w, h := screen.Size()

	g.Gui.SetTransform(gui.Transform{X:0, Y: 0, W: float32(w), H: float32(h)})

	//get animation system and run draw
	systems[AnimationUpdate].(*animation.UpdateSystem).Draw(screen)

	tiled.DrawMap(screen, testmap)

	//draw GUI
	g.Gui.DrawTree(screen)
}

// internal resolution
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// entrypoint
func main() {

	//external resolution
	ebiten.SetWindowSize(640, 480)

	//title
	ebiten.SetWindowTitle("Open Mario Maker")

	game := Game{}

	elm := gui.MakeElement(playButtonImage)

	game.Gui.AddChild(&elm)

	trans := gui.MakeTransformWithImage(playButtonImage)

	trans.XPercent = 50
	trans.YPercent = 50

	elm.SetTransform(trans)

	//run game and handle errors
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
