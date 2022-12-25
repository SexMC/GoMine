package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Cow struct {
	entities.MobBase
}

func NewCow(nametag string, pos mgl64.Vec3, hasai bool) *Cow {
	z := &Cow{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Cow) Type() world.EntityType {
	return CowType{}
}

type CowType struct{}

func (CowType) EncodeEntity() string { return "estral:entity" }
func (CowType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (CowType) DecodeNBT(data map[string]any) world.Entity {
	z := NewCow(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (CowType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Cow)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
