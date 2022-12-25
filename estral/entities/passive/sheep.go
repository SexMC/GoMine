package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Sheep struct {
	entities.MobBase
}

func NewSheep(nametag string, pos mgl64.Vec3, hasai bool) *Sheep {
	z := &Sheep{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Sheep) Type() world.EntityType {
	return SheepType{}
}

type SheepType struct{}

func (SheepType) EncodeEntity() string { return "minecraft:sheep" }
func (SheepType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (SheepType) DecodeNBT(data map[string]any) world.Entity {
	z := NewSheep(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (SheepType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Sheep)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
