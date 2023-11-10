package common

// Rect is an immutable data structure representing a rectangle.
type Rect struct {
	minPoint Point
	w, h     int
}

// NewRect creates a new Rect whose top-left corner is specified as (x, y) and whose width and height are specified.
// Width and height should both be positive
func NewRect(x, y, w, h int) Rect {
	return Rect{Point{x, y}, w, h}
}

// NewRectByPoints creates a new Rect with opposite corners specified by the given points
func NewRectByPoints(p1, p2 Point) Rect {
	// Don't assume positioning of the points
	minx := Min(p1.x, p2.x)
	miny := Min(p1.y, p2.y)
	return Rect{Point{minx, miny}, Abs(p2.x-p1.x) + 1, Abs(p2.y-p1.y) + 1}
}

// Width returns the maxPoint.x - minPoint.x
func (r Rect) Width() int {
	return r.w
}

// Width returns the maxPoint.y - minPoint.y
func (r Rect) Height() int {
	return r.h
}

// Location is the point of the top-left corner
func (r Rect) Location() Point {
	return r.minPoint
}

// MaxPoint is the point of the bottom-right corner
func (r Rect) MaxPoint() Point {
	return r.minPoint.Add(Point{r.w - 1, r.h - 1})
}

// Contains returns true if the given point is within this rectangle
func (r Rect) Contains(p Point) bool {
	return p.x >= r.minPoint.x && p.x < r.minPoint.x+r.w && p.y >= r.minPoint.y && p.y < r.minPoint.y+r.h
}
