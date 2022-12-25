package register

import (
	"github.com/EstralMC/GoMine/estral/entities/hostile"
	"github.com/EstralMC/GoMine/estral/entities/passive"
	"github.com/EstralMC/GoMine/server/world"
)

// RegisterEntities registers various entities/models including Custom Entities
func RegisterEntities() {
	world.RegisterEntity(&hostile.Drowned{})
	world.RegisterEntity(&hostile.Guardian{})
	world.RegisterEntity(&hostile.Hoglin{})
	world.RegisterEntity(&hostile.Husk{})
	world.RegisterEntity(&hostile.Ravager{})
	world.RegisterEntity(&hostile.Vindicator{})
	world.RegisterEntity(&hostile.Wither{})
	world.RegisterEntity(&hostile.Zombie{})
	world.RegisterEntity(&passive.Allay{})
	world.RegisterEntity(&passive.Axolotl{})
	world.RegisterEntity(&passive.Bat{})
	world.RegisterEntity(&passive.Bee{})
	world.RegisterEntity(&passive.Cat{})
	world.RegisterEntity(&passive.Chicken{})
	world.RegisterEntity(&passive.Cod{})
	world.RegisterEntity(&passive.Cow{})
	world.RegisterEntity(&passive.Dolphin{})
	world.RegisterEntity(&passive.Donkey{})
	world.RegisterEntity(&passive.Fox{})
	world.RegisterEntity(&passive.Frog{})
	world.RegisterEntity(&passive.GlowSquid{})
	world.RegisterEntity(&passive.PolarBear{})
	world.RegisterEntity(&passive.Pufferfish{})
	world.RegisterEntity(&passive.Rabbit{})
	world.RegisterEntity(&passive.Sheep{})
	world.RegisterEntity(&passive.Squid{})
	world.RegisterEntity(&passive.Strider{})
	world.RegisterEntity(&passive.Villager{})
	world.RegisterEntity(&passive.VillagerV2{})
}
