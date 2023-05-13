package chapter3

import "time"

type Birthday time.Time

func (b Birthday) IsYoungerThen(other time.Time) bool {
	return time.Time(b).After(other)
}

func (b Birthday) IsAdult() bool {
	return time.Time(b).AddDate(18, 0, 0).Before(time.Now())
}
