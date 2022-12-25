package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Fox struct {
	entities.MobBase
}

func NewFox(nametag string, pos mgl64.Vec3, hasai bool) *Fox {
	z := &Fox{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Fox) Type() world.EntityType {
	return FoxType{}
}

type FoxType struct{}

func (FoxType) EncodeEntity() string { return "minecraft:fox" }
func (FoxType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (FoxType) DecodeNBT(data map[string]any) world.Entity {
	z := NewFox(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (FoxType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Fox)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
