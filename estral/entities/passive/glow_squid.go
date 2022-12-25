package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type GlowSquid struct {
	entities.MobBase
}

func NewGlowSquid(nametag string, pos mgl64.Vec3, hasai bool) *GlowSquid {
	z := &GlowSquid{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*GlowSquid) Type() world.EntityType {
	return GlowSquidType{}
}

type GlowSquidType struct{}

func (GlowSquidType) EncodeEntity() string { return "minecraft:glow_squid" }
func (GlowSquidType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (GlowSquidType) DecodeNBT(data map[string]any) world.Entity {
	z := NewGlowSquid(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (GlowSquidType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*GlowSquid)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
