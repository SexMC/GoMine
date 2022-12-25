package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Bat struct {
	entities.MobBase
}

func NewBat(nametag string, pos mgl64.Vec3, hasai bool) *Bat {
	z := &Bat{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Bat) Type() world.EntityType {
	return BatType{}
}

type BatType struct{}

func (BatType) EncodeEntity() string { return "minecraft:bat" }
func (BatType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (BatType) DecodeNBT(data map[string]any) world.Entity {
	z := NewBat(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (BatType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Bat)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
