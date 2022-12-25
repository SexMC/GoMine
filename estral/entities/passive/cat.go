package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Cat struct {
	entities.MobBase
}

func NewCat(nametag string, pos mgl64.Vec3, hasai bool) *Cat {
	z := &Cat{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Cat) Type() world.EntityType {
	return CatType{}
}

type CatType struct{}

func (CatType) EncodeEntity() string { return "minecraft:cat" }
func (CatType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (CatType) DecodeNBT(data map[string]any) world.Entity {
	z := NewCat(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (CatType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Cat)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
