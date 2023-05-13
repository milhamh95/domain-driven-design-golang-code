package chapter3

import "errors"

const METERS_TO_INCHES_RATIO = 39.3701

type Distance struct {
	meters float64
}

func NewDistance(meters float64) (Distance, error) {
	if meters < 0 {
		return Distance{}, errors.New("invalid distance")
	}

	return Distance{
		meters: meters,
	}, nil
}

func (d Distance) AsMeters() float64 {
	return d.meters
}

func (d Distance) AsInches() float64 {
	return d.meters * METERS_TO_INCHES_RATIO
}

func (d Distance) OfMeters(meters float64) Distance {
	return Distance{
		meters: meters,
	}
}

func (d Distance) OfInches(inches float64) Distance {
	return Distance{
		meters: inches / METERS_TO_INCHES_RATIO,
	}
}

func (d Distance) Equal(distance Distance) bool {
	return d == distance
}
