package hostile

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Wither struct {
	entities.MobBase
}

func NewWither(nametag string, pos mgl64.Vec3, hasai bool) *Wither {
	z := &Wither{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Wither) Type() world.EntityType {
	return WitherType{}
}

type WitherType struct{}

func (WitherType) EncodeEntity() string { return "minecraft:wither" }
func (WitherType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (WitherType) DecodeNBT(data map[string]any) world.Entity {
	z := NewWither(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (WitherType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Wither)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
