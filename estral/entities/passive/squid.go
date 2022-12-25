package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Squid struct {
	entities.MobBase
}

func NewSquid(nametag string, pos mgl64.Vec3, hasai bool) *Squid {
	z := &Squid{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Squid) Type() world.EntityType {
	return SquidType{}
}

type SquidType struct{}

func (SquidType) EncodeEntity() string { return "minecraft:squid" }
func (SquidType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (SquidType) DecodeNBT(data map[string]any) world.Entity {
	z := NewSquid(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (SquidType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Squid)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
