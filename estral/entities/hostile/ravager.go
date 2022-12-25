package hostile

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Ravager struct {
	entities.MobBase
}

func NewRavager(nametag string, pos mgl64.Vec3, hasai bool) *Ravager {
	z := &Ravager{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Ravager) Type() world.EntityType {
	return RavagerType{}
}

type RavagerType struct{}

func (RavagerType) EncodeEntity() string { return "minecraft:ravager" }
func (RavagerType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (RavagerType) DecodeNBT(data map[string]any) world.Entity {
	z := NewRavager(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (RavagerType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Ravager)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
