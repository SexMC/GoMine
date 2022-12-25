package entity

import (
	"math"
	"math/rand"

	"github.com/EstralMC/GoMine/server/block"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/internal/nbtconv"
	"github.com/EstralMC/GoMine/server/item"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/df-mc/atomic"
	"github.com/go-gl/mathgl/mgl64"
)

// FallingBlock is the entity form of a block that appears when a gravity-affected block loses its support.
type FallingBlock struct {
	transform

	block        world.Block
	fallDistance atomic.Float64

	c *MovementComputer
}

// NewFallingBlock creates a new FallingBlock entity.
func NewFallingBlock(block world.Block, pos mgl64.Vec3) *FallingBlock {
	b := &FallingBlock{
		block: block,
		c: &MovementComputer{
			Gravity:           0.04,
			Drag:              0.02,
			DragBeforeGravity: true,
		},
	}
	b.transform = newTransform(b, pos)
	return b
}

// Type returns FallingBlockType.
func (*FallingBlock) Type() world.EntityType {
	return FallingBlockType{}
}

// Block ...
func (f *FallingBlock) Block() world.Block {
	return f.block
}

// FallDistance ...
func (f *FallingBlock) FallDistance() float64 {
	return f.fallDistance.Load()
}

// damager ...
type damager interface {
	Damage() (damagePerBlock, maxDamage float64)
}

// breakable ...
type breakable interface {
	Break() world.Block
}

// landable ...
type landable interface {
	Landed(w *world.World, pos cube.Pos)
}

// Tick ...
func (f *FallingBlock) Tick(w *world.World, _ int64) {
	f.mu.Lock()
	m := f.c.TickMovement(f, f.pos, f.vel, 0, 0)
	f.pos, f.vel = m.pos, m.vel
	f.mu.Unlock()

	m.Send()

	distThisTick := f.vel.Y()
	if distThisTick < f.fallDistance.Load() {
		f.fallDistance.Sub(distThisTick)
	} else {
		f.fallDistance.Store(0)
	}

	pos := cube.PosFromVec3(m.pos)
	if pos[1] < w.Range()[0] {
		_ = f.Close()
	}

	if a, ok := f.block.(Solidifiable); (ok && a.Solidifies(pos, w)) || f.c.OnGround() {
		if d, ok := f.block.(damager); ok {
			damagePerBlock, maxDamage := d.Damage()
			if dist := math.Ceil(f.fallDistance.Load() - 1.0); dist > 0 {
				force := math.Min(math.Floor(dist*damagePerBlock), maxDamage)
				for _, e := range w.EntitiesWithin(f.Type().BBox(f).Translate(m.pos).Grow(0.05), f.ignores) {
					e.(Living).Hurt(force, block.DamageSource{Block: f.block})
				}
				if b, ok := f.block.(breakable); ok && force > 0.0 && rand.Float64() < 0.05+dist*0.05 {
					f.block = b.Break()
				}
			}
		}

		if l, ok := f.block.(landable); ok {
			l.Landed(w, pos)
		}

		b := w.Block(pos)
		if r, ok := b.(replaceable); ok && r.ReplaceableBy(f.block) {
			w.SetBlock(pos, f.block, nil)
		} else {
			if i, ok := f.block.(world.Item); ok {
				w.AddEntity(NewItem(item.NewStack(i, 1), pos.Vec3Middle()))
			}
		}

		_ = f.Close()
	}
}

// New creates and returns an FallingBlock with the world.Block and position provided. It doesn't spawn the FallingBlock
// by itself.
func (f *FallingBlock) New(bl world.Block, pos mgl64.Vec3) world.Entity {
	return NewFallingBlock(bl, pos)
}

// Explode ...
func (f *FallingBlock) Explode(mgl64.Vec3, float64, block.ExplosionConfig) {
	_ = f.Close()
}

// ignores returns whether the FallingBlock should ignore collision with the entity passed.
func (f *FallingBlock) ignores(entity world.Entity) bool {
	_, ok := entity.(Living)
	return !ok || entity == f
}

// Solidifiable represents a block that can solidify by specific adjacent blocks. An example is concrete
// powder, which can turn into concrete by touching water.
type Solidifiable interface {
	// Solidifies returns whether the falling block can solidify at the position it is currently in. If so,
	// the block will immediately stop falling.
	Solidifies(pos cube.Pos, w *world.World) bool
}

type replaceable interface {
	ReplaceableBy(b world.Block) bool
}

// FallingBlockType is a world.EntityType implementation for FallingBlock.
type FallingBlockType struct{}

func (FallingBlockType) EncodeEntity() string { return "minecraft:falling_block" }
func (FallingBlockType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.49, 0, -0.49, 0.49, 0.98, 0.49)
}

func (FallingBlockType) DecodeNBT(data map[string]any) world.Entity {
	b := nbtconv.MapBlock(data, "FallingBlock")
	if b == nil {
		return nil
	}
	n := NewFallingBlock(b, nbtconv.MapVec3(data, "Pos"))
	n.SetVelocity(nbtconv.MapVec3(data, "Motion"))
	n.fallDistance.Store(nbtconv.Map[float64](data, "FallDistance"))
	return n
}

func (FallingBlockType) EncodeNBT(e world.Entity) map[string]any {
	f := e.(*FallingBlock)
	return map[string]any{
		"UniqueID":     -rand.Int63(),
		"FallDistance": f.FallDistance(),
		"Pos":          nbtconv.Vec3ToFloat32Slice(f.Position()),
		"Motion":       nbtconv.Vec3ToFloat32Slice(f.Velocity()),
		"FallingBlock": nbtconv.WriteBlock(f.block),
	}
}
