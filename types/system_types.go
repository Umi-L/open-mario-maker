package types

import "github.com/umi-l/open-mario-maker/animation"

type SystemIndex int32
type GameState int32

const (
	AnimationUpdate SystemIndex = 0
	Forces                      = 1
	Gravity                     = 2
)

const (
	MainMenu GameState = 0
	Playing            = 1
	Editing            = 2
)

type PossibleSystems interface {
	*animation.UpdateSystem
}
