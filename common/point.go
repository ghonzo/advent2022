package common

// Point is an immutable data structure representing an X and Y coordinate pair.
//
// I know we could have used image.Point, but I wanted to enforce immutability and plus
// it's just fun to write.
type Point struct {
	x, y int
}

// NewPoint is the how we create a new Point. Can't use literal Point{X,Y} syntax outside this package.
func NewPoint(x, y int) Point {
	return Point{x, y}
}

// X returns the x-coordinate
func (p Point) X() int {
	return p.x
}

// Y returns the y-coordinate
func (p Point) Y() int {
	return p.y
}

// Add returns a new Point with x- and y-coordinates added
func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

// Sub returns a new point with (p.x-q.x, p.y-q.y)
func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
}

// Times scales each coordinate by m
func (p Point) Times(m int) Point {
	return Point{p.x * m, p.y * m}
}

// Left returns a new point that is rotated 90 degress counter-clockwise around the origin
func (p Point) Left() Point {
	return Point{p.y, -p.x}
}

// LeftAround returns a new point that is rotated 90 degress counter-clockwise around the given point
func (p Point) LeftAround(p2 Point) Point {
	return p.Sub(p2).Left().Add(p2)
}

// Right returns a new point that is rotated 90 degress clockwise around the origin
func (p Point) Right() Point {
	return Point{-p.y, p.x}
}

// RightAround returns a new point that is rotated 90 degress clockwise around the given point
func (p Point) RightAround(p2 Point) Point {
	return p.Sub(p2).Right().Add(p2)
}

// Reflect returns a new point that is reflected through the origin
func (p Point) Reflect() Point {
	return Point{-p.x, -p.y}
}

// ReflectAround returns a new pont that is reflected through the given point
func (p Point) ReflectAround(p2 Point) Point {
	return p.Sub(p2).Reflect().Add(p2)
}

// ManhattanDistance returns the sum of the distance of the x- and y-coordinates from the origin
func (p Point) ManhattanDistance() int {
	return Abs(p.x) + Abs(p.y)
}

// All of these directions (Up Down Left Right) assume "UP" and "LEFT" mean -1 while "DOWN" and "RIGHT" mean +1
var (
	UL = Point{-1, -1}
	U  = Point{0, -1}
	UR = Point{1, -1}
	L  = Point{-1, 0}
	/* skip 0,0 */
	R  = Point{1, 0}
	DL = Point{-1, 1}
	D  = Point{0, 1}
	DR = Point{1, 1}
)

// Can use compass directions as aliases to the above directions if you prefer
var (
	NW = UL
	N  = U
	NE = UR
	W  = L
	E  = R
	SW = DL
	S  = D
	SE = DR
)

// AllDirections is a slice of all the different offsets that repesent the different
// directions around a point. Does not include "center" or "zero"
var AllDirections = []Point{UL, U, UR, L, R, DL, D, DR}

// SurroundingPoints returns a channel of all the points (8 of them) that surround the given point
func (p Point) SurroundingPoints() <-chan Point {
	ch := make(chan Point)
	go func() {
		for _, d := range AllDirections {
			ch <- p.Add(d)
		}
		close(ch)
	}()
	return ch
}

// SurroundingCardinals returns a channel of all the points (4 of them) that surround the given point in cardinal directions
func (p Point) SurroundingCardinals() <-chan Point {
	ch := make(chan Point)
	go func() {
		for _, d := range []Point{R, D, L, U} {
			ch <- p.Add(d)
		}
		close(ch)
	}()
	return ch
}
