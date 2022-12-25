package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type PolarBear struct {
	entities.MobBase
}

func NewPolarBear(nametag string, pos mgl64.Vec3, hasai bool) *PolarBear {
	z := &PolarBear{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*PolarBear) Type() world.EntityType {
	return PolarBearType{}
}

type PolarBearType struct{}

func (PolarBearType) EncodeEntity() string { return "minecraft:polar_bear" }
func (PolarBearType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (PolarBearType) DecodeNBT(data map[string]any) world.Entity {
	z := NewPolarBear(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (PolarBearType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*PolarBear)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
