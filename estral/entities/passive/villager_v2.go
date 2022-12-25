package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type VillagerV2 struct {
	entities.MobBase
}

func NewVillagerV2(nametag string, pos mgl64.Vec3, hasai bool) *VillagerV2 {
	z := &VillagerV2{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*VillagerV2) Type() world.EntityType {
	return VillagerV2Type{}
}

type VillagerV2Type struct{}

func (VillagerV2Type) EncodeEntity() string { return "minecraft:villager_v2" }
func (VillagerV2Type) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (VillagerV2Type) DecodeNBT(data map[string]any) world.Entity {
	z := NewVillagerV2(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (VillagerV2Type) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*VillagerV2)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
