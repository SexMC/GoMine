package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Frog struct {
	entities.MobBase
}

func NewFrog(nametag string, pos mgl64.Vec3, hasai bool) *Frog {
	z := &Frog{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Frog) Type() world.EntityType {
	return FrogType{}
}

type FrogType struct{}

func (FrogType) EncodeEntity() string { return "minecraft:frog" }
func (FrogType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (FrogType) DecodeNBT(data map[string]any) world.Entity {
	z := NewFrog(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (FrogType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Frog)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
