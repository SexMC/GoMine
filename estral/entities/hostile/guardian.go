package hostile

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Guardian struct {
	entities.MobBase
}

func NewGuardian(nametag string, pos mgl64.Vec3, hasai bool) *Guardian {
	z := &Guardian{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Guardian) Type() world.EntityType {
	return GuardianType{}
}

type GuardianType struct{}

func (GuardianType) EncodeEntity() string { return "minecraft:guardian" }
func (GuardianType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (GuardianType) DecodeNBT(data map[string]any) world.Entity {
	z := NewGuardian(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (GuardianType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Guardian)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
