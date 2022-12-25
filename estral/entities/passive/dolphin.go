package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Dolphin struct {
	entities.MobBase
}

func NewDolphin(nametag string, pos mgl64.Vec3, hasai bool) *Dolphin {
	z := &Dolphin{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Dolphin) Type() world.EntityType {
	return DolphinType{}
}

type DolphinType struct{}

func (DolphinType) EncodeEntity() string { return "minecraft:dolphin" }
func (DolphinType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (DolphinType) DecodeNBT(data map[string]any) world.Entity {
	z := NewDolphin(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (DolphinType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Dolphin)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
