package chapter3

type Point struct {
	// we set x and y with lowercase to
	// prevent them being exported and mutated
	x int
	y int
}

// using pointer every object stored in differene memory locations
// so the test will fail
// but by using value, we compare them using attribute value
func NewPoint(x, y int) Point {
	return Point{
		x: x,
	}
}

const (
	DirectionUnknown = iota
	DirectionNorth
	DirectionSouth
	DirectionEast
	DirectionWest
)

func TrackPlayer() {
	currLocation := NewPoint(3, 4)
	currLocation = Move(currLocation, DirectionNorth)
}

func Move(currLocation Point, direction int) Point {
	switch direction {
	case DirectionNorth:
		return NewPoint(currLocation.x, currLocation.y+1)
	case DirectionSouth:
		return NewPoint(currLocation.x, currLocation.y-1)
	case DirectionEast:
		return NewPoint(currLocation.x+1, currLocation.y)
	case DirectionWest:
		return NewPoint(currLocation.x-1, currLocation.y)
	default:
		// do a barrel roll
	}

	return currLocation
}
