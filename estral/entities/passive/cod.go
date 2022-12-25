package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Cod struct {
	entities.MobBase
}

func NewCod(nametag string, pos mgl64.Vec3, hasai bool) *Cod {
	z := &Cod{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Cod) Type() world.EntityType {
	return CodType{}
}

type CodType struct{}

func (CodType) EncodeEntity() string { return "minecraft:cod" }
func (CodType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (CodType) DecodeNBT(data map[string]any) world.Entity {
	z := NewCod(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (CodType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Cod)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
