package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Strider struct {
	entities.MobBase
}

func NewStrider(nametag string, pos mgl64.Vec3, hasai bool) *Strider {
	z := &Strider{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)
	return z
}

func (*Strider) Type() world.EntityType {
	return StriderType{}
}

type StriderType struct{}

func (StriderType) EncodeEntity() string { return "minecraft:strider" }
func (StriderType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (StriderType) DecodeNBT(data map[string]any) world.Entity {
	z := NewStrider(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (StriderType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Strider)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
