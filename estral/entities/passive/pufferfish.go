package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Pufferfish struct {
	entities.MobBase
}

func NewPufferfish(nametag string, pos mgl64.Vec3, hasai bool) *Pufferfish {
	z := &Pufferfish{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Pufferfish) Type() world.EntityType {
	return PufferfishType{}
}

type PufferfishType struct{}

func (PufferfishType) EncodeEntity() string { return "minecraft:pufferfish" }
func (PufferfishType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (PufferfishType) DecodeNBT(data map[string]any) world.Entity {
	z := NewPufferfish(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (PufferfishType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Pufferfish)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
