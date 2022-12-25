package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Donkey struct {
	entities.MobBase
}

func NewDonkey(nametag string, pos mgl64.Vec3, hasai bool) *Donkey {
	z := &Donkey{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Donkey) Type() world.EntityType {
	return DonkeyType{}
}

type DonkeyType struct{}

func (DonkeyType) EncodeEntity() string { return "minecraft:donkey" }
func (DonkeyType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (DonkeyType) DecodeNBT(data map[string]any) world.Entity {
	z := NewDonkey(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (DonkeyType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Donkey)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
