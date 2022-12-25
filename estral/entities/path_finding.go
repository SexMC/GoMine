package entities

import (
	"math"
	"sync"
	"time"

	"github.com/EstralMC/GoMine/estral/utils"
	"github.com/EstralMC/GoMine/server/block"
	"github.com/EstralMC/GoMine/server/block/cube"
	"github.com/EstralMC/GoMine/server/item"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/beefsack/go-astar"
	"github.com/go-gl/mathgl/mgl64"
)

var Target *Tile

var ClimbableBlocks = [17]world.Block{block.Dirt{}, block.Wood{}, block.Log{}, block.Planks{}, block.Stairs{}, block.Stone{}, block.Glass{}, block.Bricks{}, block.Sand{}, block.Sandstone{}, block.Quartz{}, block.QuartzPillar{}, block.Terracotta{}, block.DirtPath{}, block.Mud{}, block.Coral{}, block.Sponge{}}
var PassableBlocks = [6]world.Block{block.Air{}, block.Grass{}, block.TallGrass{}, block.Carpet{}, block.Slab{}, block.WoodTrapdoor{}}

// Tile holds the Position and Cost of a Tile
type Tile struct {
	Position mgl64.Vec3
	Cost     float64
	World    *world.World
}

// PathNeighbors returns a slice of tiles in four directions (N, W, E, S)
func (t *Tile) PathNeighbors() []astar.Pather {
	directions := [8]mgl64.Vec3{{0, 0, 1}, {0, 0, -1}, {1, 0, 0}, {-1, 0, 0}, {-1, 0, -1}, {1, 0, 1}, {-1, 0, 1}, {1, 0, -1}}

	var wg sync.WaitGroup
	var neighbors []astar.Pather
	ch := make(chan *Tile)

	for _, d := range directions {
		wg.Add(1)
		go get(t.Position.Add(d), t, &wg, ch)
	}

	go func() {
		for v := range ch {
			neighbors = append(neighbors, v)
		}
	}()

	wg.Wait()
	close(ch)

	return neighbors
}

// PathNeighborCost is the effective cost of moving from one tile to another
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	tile := to.(*Tile)
	return tile.Cost
}

// PathEstimatedCost returns the ManhattanDistance between two tiles
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	return t.ManhattanDistance(to)
}

// ManhattanDistance calculates orthogonal distance between two tiles
func (t *Tile) ManhattanDistance(to astar.Pather) float64 {
	tile := to.(*Tile)
	return math.Abs(t.Position.X()-tile.Position.X()) + math.Abs(t.Position.Z()-tile.Position.Z())
}

// Path holds the Tiles, Distance and World of a Path created between any two tiles
type Path struct {
	Tiles    []astar.Pather
	Distance float64
	Found    bool
	World    *world.World
}

// Show shows the path by replacing all the tiles in the path by redwool
func (p *Path) Show() {
	for _, t := range p.Tiles {
		pos := t.(*Tile).Position
		p.World.SetBlock(cube.PosFromVec3(pos.Add(mgl64.Vec3{0, 5, 0})), block.Wool{Colour: item.ColourRed()}, nil)
	}
}

// Walk makes the entity walk on the Path
func (p *Path) Walk(mob *MobBase) {
	go func() {
		var path []*Tile

		for i := range p.Tiles {
			n := p.Tiles[len(p.Tiles)-1-i]
			path = append(path, n.(*Tile))
		}

		for _, p := range path {
			pos := p.Position
			delta := pos.Sub(mob.Position())
			yaw, pitch := getYawPitch(delta)

			mob.Move(pos, yaw, pitch)

			time.Sleep(10 * time.Millisecond)
		}
	}()
}

// GetPathBetween returns a list of *Tile between the start & the end
func GetPathBetween(src *MobBase, target mgl64.Vec3) *Path {
	pos := utils.GetRoundedVector(src.Position())

	Source := &Tile{
		Position: pos,
		Cost:     0,
		World:    src.World(),
	}

	Target = &Tile{
		Position: target,
		Cost:     1.00,
		World:    src.World(),
	}

	pathers, distance, found := astar.Path(Source, Target)

	return &Path{
		Tiles:    pathers,
		Distance: distance,
		Found:    found,
		World:    src.World(),
	}
}

// getYawPitch gets the Yaw and the Pitch of the entity for a delta position
func getYawPitch(delta mgl64.Vec3) (float64, float64) {
	pitch := mgl64.RadToDeg(-math.Atan2(delta.Y(), mgl64.Vec2{delta.X(), delta.Z()}.Len()))
	yaw := mgl64.RadToDeg(math.Atan2(delta.Z(), delta.X())) - 90
	if yaw < 0 {
		yaw += 360
	}

	return yaw, pitch
}

// get gets a tile in a direction (used for retrieving neighbors)
func get(direction mgl64.Vec3, t *Tile, wg *sync.WaitGroup, ch chan *Tile) {
	tile := getEffTilePos(direction, t.World)

	if tile.Cost != -1.00 {
		ch <- tile
	}

	wg.Done()
}

// getEffTilePos gets the effective tile position, returns the Target Pointer if the position is matched with the goal
func getEffTilePos(dir mgl64.Vec3, w *world.World) *Tile {
	effPos := mgl64.Vec3{}
	cost := -1.00

	down := isEntityMoveableThrough(dir, -1, w)

	if !down {
		same := isEntityMoveableThrough(dir, 0, w)

		if !same {
			up := isEntityMoveableThrough(dir, 1, w)

			if up {
				pos := dir.Add(mgl64.Vec3{0, 1, 0})
				blk := w.Block(cube.PosFromVec3(pos))

				for _, b := range ClimbableBlocks {
					if b == blk {
						effPos = pos
						cost = 3
						break
					}
				}
			}
		} else {
			effPos = dir.Add(mgl64.Vec3{0, 0, 0})
			cost = 1
		}
	} else {
		effPos = dir.Add(mgl64.Vec3{0, -1, 0})
		cost = 2
	}

	if effPos == Target.Position {
		return Target
	}

	return &Tile{
		Position: effPos,
		World:    w,
		Cost:     cost,
	}
}

// isEntityMoveableThrough returns whether the block pairs are passable by the entity
func isEntityMoveableThrough(dir mgl64.Vec3, y int, w *world.World) bool {
	v1 := dir.Add(mgl64.Vec3{0, float64(y), 0})
	v2 := dir.Add(mgl64.Vec3{0, float64(y + 1), 0})

	blk1 := w.Block(cube.PosFromVec3(v1))
	blk2 := w.Block(cube.PosFromVec3(v2))

	b1State := false
	b2State := false

	for _, b := range PassableBlocks {
		if b == blk1 {
			b1State = true
		}

		if b == blk2 {
			b2State = true
		}
	}

	if b1State && b2State {
		return true
	} else {
		return false
	}
}
