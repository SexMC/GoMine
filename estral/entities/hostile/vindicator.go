package hostile

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Vindicator struct {
	entities.MobBase
}

func NewVindicator(nametag string, pos mgl64.Vec3, hasai bool) *Vindicator {
	z := &Vindicator{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Vindicator) Type() world.EntityType {
	return VindicatorType{}
}

type VindicatorType struct{}

func (VindicatorType) EncodeEntity() string { return "minecraft:vindicator" }
func (VindicatorType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (VindicatorType) DecodeNBT(data map[string]any) world.Entity {
	z := NewVindicator(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (VindicatorType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Vindicator)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
