package hostile

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Hoglin struct {
	entities.MobBase
}

func NewHoglin(nametag string, pos mgl64.Vec3, hasai bool) *Hoglin {
	z := &Hoglin{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Hoglin) Type() world.EntityType {
	return HoglinType{}
}

type HoglinType struct{}

func (HoglinType) EncodeEntity() string { return "minecraft:hoglin" }
func (HoglinType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (HoglinType) DecodeNBT(data map[string]any) world.Entity {
	z := NewHoglin(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (HoglinType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Hoglin)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
