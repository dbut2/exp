package grid

type Grid[T any] map[[2]int]T

func (g *Grid[T]) Set(x, y int, cell T) {
	g.ensureGrid()
	(*g)[[2]int{x, y}] = cell
}

func (g *Grid[T]) Get(x, y int) T {
	g.ensureGrid()
	return (*g)[[2]int{x, y}]
}

func (g *Grid[T]) ensureGrid() {
	if g == nil {
		*g = make(Grid[T])
	}
}

type Cube[T any] map[[3]int]T

func (c *Cube[T]) Set(x, y, z int, cell T) {
	c.ensureCube()
	(*c)[[3]int{x, y, z}] = cell
}

func (c *Cube[T]) Get(x, y, z int) T {
	c.ensureCube()
	return (*c)[[3]int{x, y, z}]
}

func (c *Cube[T]) ensureCube() {
	if c == nil {
		*c = make(Cube[T])
	}
}
