package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Bee struct {
	entities.MobBase
}

func NewBee(nametag string, pos mgl64.Vec3, hasai bool) *Bee {
	z := &Bee{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Bee) Type() world.EntityType {
	return BeeType{}
}

type BeeType struct{}

func (BeeType) EncodeEntity() string { return "minecraft:bee" }
func (BeeType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (BeeType) DecodeNBT(data map[string]any) world.Entity {
	z := NewBee(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (BeeType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Bee)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
