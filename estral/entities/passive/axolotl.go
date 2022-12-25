package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Axolotl struct {
	entities.MobBase
}

func NewAxolotl(nametag string, pos mgl64.Vec3, hasai bool) *Axolotl {
	z := &Axolotl{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Axolotl) Type() world.EntityType {
	return AxolotlType{}
}

type AxolotlType struct{}

func (AxolotlType) EncodeEntity() string { return "minecraft:axolotl" }
func (AxolotlType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (AxolotlType) DecodeNBT(data map[string]any) world.Entity {
	z := NewAxolotl(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (AxolotlType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Axolotl)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
