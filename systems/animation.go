package systems

import (
	"github.com/umi-l/open-mario-maker/geometry"
	"github.com/umi-l/open-mario-maker/types"
	"github.com/umi-l/waloader"
	"github.com/wfranczyk/ento"
)

type AnimationUpdateSystem struct {
	Transform *geometry.Transform `ento:"required"`
	Animation *waloader.Animation `ento:"required"`

	//singleton
	DeltaTime *types.DeltaTime `ento:"required"`
	Screen    *types.Screen    `ento:"required"`
}

func (s *AnimationUpdateSystem) Update(entity *ento.Entity) {
	s.Animation.UpdateTimer(s.DeltaTime.Dt)
	s.Animation.Draw(s.Screen.Screen, s.Transform.X, s.Transform.Y, s.Transform.Rotation)
}
