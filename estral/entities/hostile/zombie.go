package hostile

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/item"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Zombie struct {
	entities.MobBase
}

func NewZombie(nametag string, pos mgl64.Vec3, hasai bool) *Zombie {
	z := &Zombie{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	z.Armour().SetHelmet(item.NewStack(item.Helmet{Tier: item.ArmourTierNetherite{}}, 1))
	z.Armour().SetChestplate(item.NewStack(item.Chestplate{Tier: item.ArmourTierNetherite{}}, 1))
	z.Armour().SetLeggings(item.NewStack(item.Leggings{Tier: item.ArmourTierNetherite{}}, 1))
	z.Armour().SetBoots(item.NewStack(item.Boots{Tier: item.ArmourTierNetherite{}}, 1))
	return z
}

func (*Zombie) Type() world.EntityType {
	return ZombieType{}
}

type ZombieType struct{}

func (ZombieType) EncodeEntity() string { return "minecraft:zombie" }
func (ZombieType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (ZombieType) DecodeNBT(data map[string]any) world.Entity {
	z := NewZombie(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (ZombieType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Zombie)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
