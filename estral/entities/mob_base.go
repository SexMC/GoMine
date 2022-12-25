package entities

import (
	"math"
	"sync"
	"time"

	"github.com/EstralMC/GoMine/estral/entities/ai"

	"github.com/EstralMC/GoMine/estral/console"
	"github.com/EstralMC/GoMine/server/entity"
	"github.com/EstralMC/GoMine/server/entity/effect"
	"github.com/EstralMC/GoMine/server/item"
	"github.com/EstralMC/GoMine/server/item/enchantment"
	"github.com/EstralMC/GoMine/server/item/inventory"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/EstralMC/GoMine/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// MobBase holds an instance of Mob class with various properties that the entity possesses
type MobBase struct {
	entity             Mob
	velocity, position mgl64.Vec3

	health, maxHealth, speed, yaw, pitch float64
	hasAi, attackImmune                  bool

	effects []effect.Effect
	targets []entity.Living

	armour *inventory.Armour

	nametag        string
	nametagVisible bool
	immunity       int64
	behaviorTree   ai.Tree
}

// NewMobBase creates a new instance of MobBase
func NewMobBase(e Mob, pos mgl64.Vec3, nametag string, hasAi bool) MobBase {
	z := MobBase{}

	z = MobBase{
		entity:         e,
		position:       pos,
		velocity:       mgl64.Vec3{},
		targets:        []entity.Living{},
		nametag:        nametag,
		nametagVisible: false,

		hasAi:        hasAi,
		attackImmune: false,
		health:       100,
		maxHealth:    100,
		immunity:     time.Now().UnixMilli(),
		yaw:          0,
		pitch:        0,
		armour:       inventory.NewArmour(z.broadcastArmour),
	}

	return z
}

// Name returns the nametag
func (z *MobBase) Name() string {
	return z.nametag
}

// NameTag returns the nametag only when the Name Tag Visibility is toggled
func (z *MobBase) NameTag() string {
	if z.NameTagVisible() {
		return z.nametag
	}

	return ""
}

// NameTagVisible returns whether the Name Tag is visible or not
func (z *MobBase) NameTagVisible() bool {
	return z.nametagVisible
}

// SetNameTagVisible sets whether the name tag should be visible or not
func (z *MobBase) SetNameTagVisible(v bool) {
	z.nametagVisible = v
}

// Armour returns the armour that the entity is wearing
func (z *MobBase) Armour() *inventory.Armour {
	return z.armour
}

// broadcastArmour is used to broadcast any armour changes for the entity to it's viewers
func (z *MobBase) broadcastArmour(_ int, before, after item.Stack) {
	if before.Comparable(after) && before.Empty() == after.Empty() {
		return
	}

	for _, viewer := range z.Viewers() {
		viewer.ViewEntityArmour(z.entity)
	}
}

// HasAi returns whether the Ai is enabled. This is enabled by default.
// Disabling this would not lead to removal of Behaviors or Target Behaviors from the tree
func (z *MobBase) HasAi() bool {
	return z.hasAi
}

func (z *MobBase) Targets() []entity.Living {
	return z.targets
}

func (z *MobBase) Dead() bool {
	return z.Health() <= mgl64.Epsilon
}

func (z *MobBase) Heal(health float64, _ world.HealingSource) {
	if z.Dead() || health < 0 || health > z.MaxHealth() {
		return
	}

	z.AddHealth(health)
}

func (z *MobBase) Health() float64 {
	return z.health
}

func (z *MobBase) MaxHealth() float64 {
	return z.maxHealth
}

func (z *MobBase) SetMaxHealth(h float64) {
	z.maxHealth = h
}

func (z *MobBase) SetHealth(h float64) {
	z.health = h
}

func (z *MobBase) AddHealth(h float64) {
	z.health = z.Health() + h
}

func (z *MobBase) SetImmunity(duration int64) {
	z.immunity = duration + time.Now().UnixMilli()
}

func (z *MobBase) AttackImmune() bool {
	return time.Now().UnixMilli() < z.immunity
}

func (z *MobBase) AddEffect(e effect.Effect) {
	z.effects = append(z.effects, e)
}

func (z *MobBase) RemoveEffect(_ effect.Type) {
	// TODO
}

func (z *MobBase) Effects() []effect.Effect {
	return z.effects
}

func (z *MobBase) Speed() float64 {
	return z.speed
}

func (z *MobBase) SetSpeed(s float64) {
	z.speed = s
}

func (z *MobBase) Velocity() mgl64.Vec3 {
	return z.velocity
}

func (z *MobBase) SetVelocity(v mgl64.Vec3) {
	z.velocity = v

	for _, viewer := range z.Viewers() {
		viewer.ViewEntityVelocity(z.entity, v)
	}
}

func (z *MobBase) Position() mgl64.Vec3 {
	return z.position
}

func (z *MobBase) GetHeadDirection() mgl64.Vec3 {
	yaw, pitch := z.Rotation()

	yaw = yaw * (3.14 / 180)
	pitch = pitch * (3.14 / 180)

	return mgl64.Vec3{-math.Sin(yaw) * -math.Cos(pitch), -math.Sin(pitch), math.Cos(yaw) * math.Cos(pitch)}
}

func (z *MobBase) Rotation() (float64, float64) { return z.yaw, z.pitch }

func (z *MobBase) World() *world.World {
	w, _ := world.OfEntity(z.entity)
	return w
}

func (z *MobBase) Close() error {
	w, _ := world.OfEntity(z.entity)
	w.RemoveEntity(z.entity)
	return nil
}

func (z *MobBase) Viewers() []world.Viewer {
	viewers := z.World().Viewers(z.Position())
	return viewers
}

func (z *MobBase) Hurt(damage float64, src world.DamageSource) (n float64, vulnerable bool) {
	if z.Dead() || damage > z.Health() {
		z.Kill()
		return 0, true
	}

	dmg := z.GetEffectiveDamage(damage, src)
	z.AddHealth(-dmg)

	console.Log.Printf("Hurt %f total health %f", dmg, z.Health())

	if src.ReducedByArmour() {
		armourDamage := int(math.Max(math.Floor(dmg/4), 1))
		for slot, it := range z.armour.Slots() {
			_ = z.armour.Inventory().SetItem(slot, z.damageItem(it, armourDamage))
		}
	}

	for _, v := range z.Viewers() {
		v.ViewEntityAction(z.entity, entity.HurtAction{})
	}

	z.SetImmunity(500)

	return damage, true
}

func (z *MobBase) Kill() {
	for _, viewer := range z.Viewers() {
		viewer.ViewEntityAction(z.entity, entity.DeathAction{})
	}

	time.AfterFunc(time.Millisecond*5000, func() {
		z.World().RemoveEntity(z.entity)
	})
}

func (z *MobBase) KnockBack(src mgl64.Vec3, force, height float64) {
	velocity := z.Position().Sub(src)
	velocity[1] = 0

	velocity = velocity.Normalize().Mul(force)
	velocity[1] = height

	var resistance float64
	for _, i := range z.Armour().Items() {
		if a, ok := i.Item().(item.Armour); ok {
			resistance += a.KnockBackResistance()
		}
	}

	z.SetVelocity(velocity.Mul(1 - resistance))
}

func (z *MobBase) Move(newPos mgl64.Vec3, newYaw, newPitch float64) {
	if z.Dead() {
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		var (
			pos        = z.Position()
			yaw, pitch = z.Rotation()
		)

		subPos := newPos.Sub(pos)
		subYaw := newYaw - yaw
		subPitch := newPitch - pitch

		deltaPos := mgl64.Vec3{subPos.X() / 50, subPos.Y() / 50, subPos.Z() / 50}
		deltaYaw := subYaw / 50
		deltaPitch := subPitch / 50

		for i := 1; i < 50; i++ {
			finalPos := z.Position().Add(deltaPos)
			finalYaw := z.yaw + deltaYaw
			finalPitch := z.pitch + deltaPitch

			for _, v := range z.Viewers() {
				v.ViewEntityMovement(z.entity, finalPos, finalYaw, finalPitch, true)
			}

			z.position = finalPos
			z.yaw = finalYaw
			z.pitch = finalPitch

			time.Sleep(10 * time.Millisecond)
		}

		wg.Done()
	}()

	wg.Wait()
}

func (z *MobBase) GetEffectiveDamage(damage float64, src world.DamageSource) float64 {
	damage = math.Max(damage, 0)

	damage -= z.Armour().DamageReduction(damage, src)
	return damage
}

func (z *MobBase) damageItem(s item.Stack, d int) item.Stack {
	if d == 0 || s.MaxDurability() == -1 {
		return s
	}

	if e, ok := s.Enchantment(enchantment.Unbreaking{}); ok {
		d = (enchantment.Unbreaking{}).Reduce(s.Item(), e.Level(), d)
	}

	if s = s.Damage(d); s.Empty() {
		z.World().PlaySound(z.Position(), sound.ItemBreak{})
	}

	return s
}

func (z *MobBase) Tick(_ *world.World, _ int64) {
}
