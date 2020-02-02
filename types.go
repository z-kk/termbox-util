package termboxUtil

type Cursor struct {
	x int
	y int
	isShow bool
}

func (c Cursor) GetPoint() (int, int) {
	return c.x, c.y
}

func (c *Cursor) SetPoint(xp, yp int) {
	c.x = xp
	c.y = yp
}

type Window struct {
	posx, posy int
	width, height int
	cursor Cursor
	header, footer []string
	strings []string
	showStrings []string
	showStringsIndex []int
	headLine int
}
