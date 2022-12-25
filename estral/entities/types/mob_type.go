package types

import (
	"strings"

	"github.com/EstralMC/GoMine/estral/entities"
	"github.com/EstralMC/GoMine/estral/entities/hostile"
	"github.com/EstralMC/GoMine/estral/entities/passive"
	"github.com/go-gl/mathgl/mgl64"
)

// GetMobFromType returns a new Mob of that type. It returns nil if the type does not exist.
func GetMobFromType(nametag string, pos mgl64.Vec3, hasAi bool, mob string) entities.Mob {
	mob = strings.ToLower(mob)

	switch mob {
	case "drowned":
		m := hostile.NewDrowned(nametag, pos, hasAi)
		return m
	case "guardian":
		m := hostile.NewGuardian(nametag, pos, hasAi)
		return m
	case "hoglin":
		m := hostile.NewHoglin(nametag, pos, hasAi)
		return m
	case "husk":
		m := hostile.NewHusk(nametag, pos, hasAi)
		return m
	case "ravager":
		m := hostile.NewRavager(nametag, pos, hasAi)
		return m
	case "vindicator":
		m := hostile.NewVindicator(nametag, pos, hasAi)
		return m
	case "wither":
		m := hostile.NewWither(nametag, pos, hasAi)
		return m
	case "zombie":
		m := hostile.NewZombie(nametag, pos, hasAi)
		return m
	case "allay":
		m := passive.NewAllay(nametag, pos, hasAi)
		return m
	case "axolotl":
		m := passive.NewAxolotl(nametag, pos, hasAi)
		return m
	case "bat":
		m := passive.NewBat(nametag, pos, hasAi)
		return m
	case "bee":
		m := passive.NewBee(nametag, pos, hasAi)
		return m
	case "cat":
		m := passive.NewCat(nametag, pos, hasAi)
		return m
	case "chicken":
		m := passive.NewChicken(nametag, pos, hasAi)
		return m
	case "cod":
		m := passive.NewCod(nametag, pos, hasAi)
		return m
	case "cow":
		m := passive.NewCow(nametag, pos, hasAi)
		return m
	case "dolphin":
		m := passive.NewDolphin(nametag, pos, hasAi)
		return m
	case "donkey":
		m := passive.NewDonkey(nametag, pos, hasAi)
		return m
	case "fox":
		m := passive.NewFox(nametag, pos, hasAi)
		return m
	case "frog":
		m := passive.NewFrog(nametag, pos, hasAi)
		return m
	case "glow_squid":
		m := passive.NewGlowSquid(nametag, pos, hasAi)
		return m
	case "polar_bear":
		m := passive.NewPolarBear(nametag, pos, hasAi)
		return m
	case "pufferfish":
		m := passive.NewPufferfish(nametag, pos, hasAi)
		return m
	case "rabbit":
		m := passive.NewRabbit(nametag, pos, hasAi)
		return m
	case "sheep":
		m := passive.NewSheep(nametag, pos, hasAi)
		return m
	case "squid":
		m := passive.NewSquid(nametag, pos, hasAi)
		return m
	case "strider":
		m := passive.NewStrider(nametag, pos, hasAi)
		return m
	case "villager":
		m := passive.NewVillager(nametag, pos, hasAi)
		return m
	case "villager_v2":
		m := passive.NewVillagerV2(nametag, pos, hasAi)
		return m
	}

	return nil
}
