package grid

import (
	"arrangio-core/entity"
	"arrangio-core/geometry"
)

// gridNode is an element in a flat array `nods`
// it is a node is a singly-linked list, links one entity to the next one in the same cell of the grid
type gridNode struct {
	entity *entity.Entity
	next   int32 // index of the next node in the same cell chain (-1 if end of the list)
}

type Grid struct {
	shiftBits uint8 // determines the size of a single grid cell as a power of two
	sizeX     int64
	sizeY     int64
	sizeZ     int64
	strideY   int64 // precomputed offsets used to flatten 3D grid coordinates
	strideZ   int64 // into a single-dimensional `heads` slice index

	queryID uint64 // ID for deduplication in `QueryBuf` -- unique timestamp for current query session

	heads     []int32    // maps each 3D cell to the the starting index of its linked list inside `nodes`
	nodes     []gridNode // stores all linked list nodes for all cells
	freeNodes []int32    // tracks which indices in the `nodes` are currently empty
}

// `maxObjectsPerCell` is a hard limit for total entity-to-cell linkages allowed
// rule of thumb: `maxObjectsPerCell` = N * 8, where are N is a number of entities
// assuming the worst of case of all 8 neighbor shapes overlapping at one cell
func NewGrid(shiftBits uint8, worldW, worldH, worldD int64, maxObjectsPerCell int) *Grid {
	cellsX := worldW >> shiftBits
	cellsY := worldH >> shiftBits
	cellsZ := worldD >> shiftBits

	totalCells := cellsX * cellsY * cellsZ

	heads := make([]int32, totalCells)
	for i := range heads {
		heads[i] = -1
	}

	nodes := make([]gridNode, maxObjectsPerCell)
	freeNodes := make([]int32, maxObjectsPerCell)
	for i := 0; i < maxObjectsPerCell; i++ {
		freeNodes[i] = int32(i)
	}

	return &Grid{
		shiftBits: shiftBits,
		sizeX:     cellsX,
		sizeY:     cellsY,
		sizeZ:     cellsZ,
		strideY:   cellsX,
		strideZ:   cellsX * cellsY,
		queryID:   0,
		heads:     heads,
		nodes:     nodes,
		freeNodes: freeNodes,
	}
}

func (g *Grid) getIndex(cx, cy, cz int64) int64 {
	return cx + (cy * g.strideY) + (cz * g.strideZ)
}

func (g *Grid) Insert(e *entity.Entity) {
	minBounds, maxBounds := e.Footprint.WorldBounds()

	minX, minY, minZ := minBounds.X>>g.shiftBits, minBounds.Y>>g.shiftBits, minBounds.Z>>g.shiftBits
	maxX, maxY, maxZ := maxBounds.X>>g.shiftBits, maxBounds.Y>>g.shiftBits, maxBounds.Z>>g.shiftBits

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				if x < 0 || x >= g.sizeX || y < 0 || y >= g.sizeY || z < 0 || z >= g.sizeZ {
					continue
				}

				cellIdx := g.getIndex(x, y, z)

				if len(g.freeNodes) == 0 {
					panic("Grid: out of memory in nodes pool! Increase maxObjectsPerCell")
				}

				// use index from the end of slice
				nodeIdx := g.freeNodes[len(g.freeNodes)-1]
				// pop element from the end
				g.freeNodes = g.freeNodes[:len(g.freeNodes)-1]

				g.nodes[nodeIdx] = gridNode{
					entity: e,
					next:   g.heads[cellIdx],
				}
				g.heads[cellIdx] = nodeIdx
			}
		}
	}
}

func (g *Grid) Remove(e *entity.Entity) {
	minBounds, maxBounds := e.Footprint.WorldBounds()

	minX, minY, minZ := minBounds.X>>g.shiftBits, minBounds.Y>>g.shiftBits, minBounds.Z>>g.shiftBits
	maxX, maxY, maxZ := maxBounds.X>>g.shiftBits, maxBounds.Y>>g.shiftBits, maxBounds.Z>>g.shiftBits

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				if x < 0 || x >= g.sizeX || y < 0 || y >= g.sizeY || z < 0 || z >= g.sizeZ {
					continue
				}

				cellIdx := g.getIndex(x, y, z)

				currentNodeIdx := g.heads[cellIdx]
				var prevNodeIdx int32 = -1

				// iterate through the list of nodes
				for currentNodeIdx != -1 {
					node := g.nodes[currentNodeIdx]

					if node.entity.ID == e.ID {
						if prevNodeIdx == -1 {
							g.heads[cellIdx] = node.next
						} else {
							g.nodes[prevNodeIdx].next = node.next
						}

						g.freeNodes = append(g.freeNodes, currentNodeIdx)
						break
					}

					prevNodeIdx = currentNodeIdx
					currentNodeIdx = node.next
				}
			}
		}
	}
}

// return all entities from searchMin point to searchMax point
func (g *Grid) QueryBuf(searchMin, searchMax geometry.Point64, buffer []*entity.Entity) []*entity.Entity {
	g.queryID++
	result := buffer[:0]

	minX, minY, minZ := searchMin.X>>g.shiftBits, searchMin.Y>>g.shiftBits, searchMin.Z>>g.shiftBits
	maxX, maxY, maxZ := searchMax.X>>g.shiftBits, searchMax.Y>>g.shiftBits, searchMax.Z>>g.shiftBits

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				if x < 0 || x >= g.sizeX || y < 0 || y >= g.sizeY || z < 0 || z >= g.sizeZ {
					continue
				}

				cellIdx := g.getIndex(x, y, z)
				nodeIdx := g.heads[cellIdx]

				for nodeIdx != -1 {
					node := g.nodes[nodeIdx]
					e := node.entity

					if e.LastQueryID != g.queryID {
						e.LastQueryID = g.queryID

						eMin, eMax := e.Footprint.WorldBounds()
						if eMax.X >= searchMin.X && eMin.X <= searchMax.X &&
							eMax.Y >= searchMin.Y && eMin.Y <= searchMax.Y &&
							eMax.Z >= searchMin.Z && eMax.Z <= searchMax.Z {
							result = append(result, e)
						}
					}
					nodeIdx = node.next
				}
			}
		}
	}
	return result
}
