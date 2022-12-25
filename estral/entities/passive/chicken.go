package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Chicken struct {
	entities.MobBase
}

func NewChicken(nametag string, pos mgl64.Vec3, hasai bool) *Chicken {
	z := &Chicken{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Chicken) Type() world.EntityType {
	return ChickenType{}
}

type ChickenType struct{}

func (ChickenType) EncodeEntity() string { return "minecraft:chicken" }
func (ChickenType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (ChickenType) DecodeNBT(data map[string]any) world.Entity {
	z := NewChicken(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (ChickenType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Chicken)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
