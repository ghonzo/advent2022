package common

// Rect is an immutable data structure representing a rectangle.
type Rect struct {
	minPoint, maxPoint Point
}

func NewRect(x, y, w, h int) Rect {
	return Rect{Point{x, y}, Point{x + w, y + h}}
}

func NewRectByPoints(p1, p2 Point) Rect {
	return Rect{p1, p2}
}

func (r Rect) Width() int {
	return r.maxPoint.x - r.minPoint.x
}

func (r Rect) Height() int {
	return r.maxPoint.y - r.minPoint.y
}

func (r Rect) Location() Point {
	return r.minPoint
}

func (r Rect) MaxPoint() Point {
	return r.maxPoint
}

func (r Rect) Contains(p Point) bool {
	return p.x >= r.minPoint.x && p.x <= r.maxPoint.x && p.y >= r.minPoint.y && p.y <= r.maxPoint.y
}
