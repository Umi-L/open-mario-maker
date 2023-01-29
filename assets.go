package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/umi-l/waloader"
)

// images
var playButtonImage *ebiten.Image

// animations
var marioSwimmingAnimation waloader.Animation

// sheets
var tilemapSheet waloader.Sheet

func (game *Game) InitAssets() {

	//load atlas
	game.Atlas = waloader.LoadAtlas("assets/atlases/", "atlas.xml")

	//resources
	charSheet := waloader.LoadSheet(game.Atlas["MarioSpriteSheet"], 16, 32)
	marioSwimmingAnimation = waloader.LoadAnimation(&charSheet, 0, 5, 0.1)

	tilemapSheet = waloader.LoadSheet(game.Atlas["tilemap"], 16, 16)
}
