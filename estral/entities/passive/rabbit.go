package passive

import (
	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/utils/nbtconv"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Rabbit struct {
	entities.MobBase
}

func NewRabbit(nametag string, pos mgl64.Vec3, hasai bool) *Rabbit {
	z := &Rabbit{}
	z.MobBase = entities.NewMobBase(z, pos, nametag, hasai)

	return z
}

func (*Rabbit) Type() world.EntityType {
	return RabbitType{}
}

type RabbitType struct{}

func (RabbitType) EncodeEntity() string { return "minecraft:rabbit" }
func (RabbitType) BBox(_ world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 2, 0.49)
}

func (RabbitType) DecodeNBT(data map[string]any) world.Entity {
	z := NewRabbit(nbtconv.Map[string](data, "Nametag"), nbtconv.MapVec3(data, "Pos"), nbtconv.Map[bool](data, "HasAI"))
	return z
}

func (RabbitType) EncodeNBT(e world.Entity) map[string]any {
	z := e.(*Rabbit)
	return map[string]any{
		"Nametag": z.MobBase.NameTag(),
		"Pos":     nbtconv.Vec3ToFloat32Slice(z.Position()),
		"HasAI":   z.MobBase.HasAi(),
	}
}
