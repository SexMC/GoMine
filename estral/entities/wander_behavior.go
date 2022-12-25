package entities

import (
	"math/rand"
	"time"

	"github.com/EstralMC/GoMine/estral/utils"
	"github.com/go-gl/mathgl/mgl64"
)

type WanderBehavior struct {
	Mob         *MobBase
	Target      mgl64.Vec3
	IsCompleted bool
}

func (b *WanderBehavior) Start() {
	b.IsCompleted = false
	z := b.Mob
	pos := utils.GetRoundedVector(z.Position())

	rand1 := rand.Intn(10) - 5
	rand2 := rand.Intn(10) - 5

	b.Target = pos.Add(mgl64.Vec3{float64(rand1), 0, float64(rand2)})
	path := GetPathBetween(z, b.Target)

	if path.Found {
		path.Walk(z)
	}

	time.AfterFunc(time.Second*10, func() {
		b.IsCompleted = true
	})
}

func (b *WanderBehavior) Completed() bool {
	return b.IsCompleted
}
