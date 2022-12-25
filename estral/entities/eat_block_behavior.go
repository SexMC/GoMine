package entities

import (
	"time"

	"github.com/EstralMC/GoMine/server/block"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/entity"
	"github.com/EstralMC/GoMine/server/world/particle"
	"github.com/go-gl/mathgl/mgl64"
)

type EatBlockBehavior struct {
	Mob         *MobBase
	BlockCoords mgl64.Vec3
	IsCompleted bool
}

func (b *EatBlockBehavior) Start() {
	z := b.Mob
	b.IsCompleted = false

	coords := z.Position()
	dir := coords.Normalize()

	blockCoords := mgl64.Vec3{coords.X() + dir.X(), coords.Y(), coords.Z() + dir.Z()}
	bk := z.World().Block(cube.PosFromVec3(blockCoords))

	if (bk != block.TallGrass{} && bk != block.DoubleTallGrass{}) {
		blockCoords = blockCoords.Add(mgl64.Vec3{0, -1, 0})
		bk = z.World().Block(cube.PosFromVec3(blockCoords))

		if (bk != block.Grass{}) {
			return
		}
	}

	b.BlockCoords = blockCoords

	b.onStart()
}

func (b *EatBlockBehavior) Completed() bool {
	return b.IsCompleted
}

func (b *EatBlockBehavior) onStart() {
	z := b.Mob

	for _, v := range z.Viewers() {
		v.ViewEntityAction(z.entity, entity.EatGrassAction{})
	}

	time.AfterFunc(time.Millisecond*1000, func() {
		b.onEnd()
	})
}

func (b *EatBlockBehavior) onEnd() {
	z := b.Mob
	w := z.World()

	blockCoords := b.BlockCoords
	bk := w.Block(cube.PosFromVec3(blockCoords))

	if (bk == block.DoubleTallGrass{}) {
		w.SetBlock(cube.PosFromVec3(blockCoords), block.TallGrass{}, nil)
	} else if (bk == block.TallGrass{}) {
		w.SetBlock(cube.PosFromVec3(blockCoords), block.Air{}, nil)
	} else if (bk == block.Grass{}) {
		w.SetBlock(cube.PosFromVec3(blockCoords), block.Dirt{}, nil)
	} else {
		return
	}

	blockBreakParticle := particle.BlockBreak{Block: bk}
	blockBreakParticle.Spawn(w, z.Position())

	b.IsCompleted = true
}
