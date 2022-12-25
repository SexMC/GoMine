package hostile

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Drowned struct {
	entities.MobBase
}

func NewDrowned(nametag string, pos mgl64.Vec3, hasai bool) *Drowned {
	z := &Drowned{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Drowned) Type() world.EntityType {
	return DrownedType{}
}

type DrownedType struct{}

func (DrownedType) EncodeEntity() string { return "minecraft:drowned" }
func (DrownedType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (DrownedType) DecodeNBT(data map[string]any) world.Entity {
	z := NewDrowned(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (DrownedType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Drowned)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
