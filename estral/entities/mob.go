package entities

import (
	"github.com/EstralMC/GoMine/server/entity"
	"github.com/EstralMC/GoMine/server/item"
	"github.com/EstralMC/GoMine/server/item/inventory"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// Mob represents a Living entities that usually has a set of ai also known as Artificial Intelligence.
// You can extend this interface to create a new type of mob
type Mob interface {
	entity.Living

	// NameTag returns the entity name tag
	NameTag() string
	// NameTagVisible returns if the nametag should be visible
	NameTagVisible() bool
	// SetNameTagVisible sets whether the entity's name tag should be visible or not.
	// This value is enabled by default
	SetNameTagVisible(v bool)

	// Targets returns a list of entities that this entities is targeting
	Targets() []entity.Living

	// Armour returns the Armour that the Mob is wearing
	Armour() *inventory.Armour
	// broadcastArmour broadcasts the mob's armor to the players
	broadcastArmour(_ int, before, after item.Stack)

	// Move moves the Mob
	Move(deltaPos mgl64.Vec3, deltaYaw, deltaPitch float64)

	// GetEffectiveDamage gets the effective damage after considering armor and effects
	GetEffectiveDamage(damage float64, src world.DamageSource) float64
	// damageItem damages the item such as tools, armor, etc.
	damageItem(s item.Stack, d int) item.Stack

	// HasAi returns whether this entities is required to use set of ai defined
	// If this is set to return false then any ai passed will be ignored
	HasAi() bool

	// AddHealth adds an amount of health to the mob
	AddHealth(h float64)
	// Kill kills the mob
	Kill()

	// GetHeadDirection returns the head direction vector
	GetHeadDirection() mgl64.Vec3
}
