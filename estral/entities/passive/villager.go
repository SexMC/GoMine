package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Villager struct {
	entities.MobBase
}

func NewVillager(nametag string, pos mgl64.Vec3, hasai bool) *Villager {
	z := &Villager{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Villager) Type() world.EntityType {
	return VillagerType{}
}

type VillagerType struct{}

func (VillagerType) EncodeEntity() string { return "minecraft:villager" }
func (VillagerType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (VillagerType) DecodeNBT(data map[string]any) world.Entity {
	z := NewVillager(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (VillagerType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Villager)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
