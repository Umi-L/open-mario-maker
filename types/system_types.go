package types

import "github.com/umi-l/open-mario-maker/animation"

type SystemIndex int32

const (
	AnimationUpdate SystemIndex = 0
	Forces                      = 1
	Gravity                     = 2
)

type PossibleSystems interface {
	*animation.UpdateSystem
}
