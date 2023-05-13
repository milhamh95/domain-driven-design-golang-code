package chapter3

import (
	"errors"
	"fmt"
)

type Money struct {
	value    float64
	currency Currency
}

type Currency struct {
	name string
}

func (m Money) ToHTML() string {
	return fmt.Sprintf(`%.2f`, m.value)
}

func (m Money) Equal(other Money) bool {
	return m == other
}

func (c Currency) Equal(other Currency) bool {
	return c == other
}

// wrong way to change the state inside value object
func (m *Money) AddAmount(amount float64) {
	m.value += amount
}

// right way to return new value objets with new state
func (m Money) WithAmount(value float64) Money {
	return Money{
		value:    m.value + value,
		currency: m.currency,
	}
}

// wrong way to change the state inside value object
func (m *Money) Deduct(other Money) {
	m.value -= other.value
}

// right way to return new value objects with new state
func (m Money) DeductedWith(other Money) Money {
	return Money{
		value:    m.value - other.value,
		currency: m.currency,
	}
}

// bind functions to values instead of reference (pointer) of value objects
// to be sure that we never change the innter state
// this immutability means that we should not validate Value Object during its whole lifetime.
// but only on creation like.
// when we want to create a new value object
// we must always perform validation and return errors if business invariants are not fulfilled
// and make value object only if it is valid
func (m Money) Deduct2(other Money) (Money, error) {
	if !m.currency.Equal(other.currency) {
		return Money{}, errors.New("currencies must be identical")
	}

	if m.value > other.value {
		return Money{}, errors.New("not enough amount to deduct")
	}

	return Money{
		value:    m.value - other.value,
		currency: m.currency,
	}, nil
}

type Salutation string

func (s Salutation) IsPerson() bool {
	return s != "company"
}

type Color struct {
	red   int
	green int
	blue  int
}

func (c Color) ToCSS() string {
	return fmt.Sprintf(`rgb(%d, %d, %d)`, c.red, c.green, c.blue)
}

func (c Color) EqualTo(color Color) bool {
	return c.red == color.red &&
		c.green == color.green &&
		c.blue == color.blue
}

func (c *Color) KeepOnlyGreen() {
	c.red = 0
	c.blue = 0
}

func (c Color) WithOnlyGreen() Color {
	return Color{
		red:   0,
		green: c.green,
		blue:  0,
	}
}

func (c Color) ToBrighter() Color {
	return Color{
		red:   Min(255, c.red+10),
		green: Min(255, c.green+10),
		blue:  Min(255, c.blue+10),
	}
}

func (c Color) ToDarker() Color {
	return Color{
		red:   Max(0, c.red-10),
		green: Max(0, c.green-10),
		blue:  Max(0, c.blue-10),
	}
}

func (c Color) Combine(other Color) Color {
	return Color{
		red:   Min(255, c.red+other.red),
		green: Min(255, c.green+other.green),
		blue:  Min(255, c.blue+other.blue),
	}
}

func (c Color) IsRed() bool {
	return c.red == 255 && c.green == 0 && c.blue == 0
}

func (c Color) IsYellow() bool {
	return c.red == 255 && c.green == 0 && c.blue == 255
}

func (c Color) IsMagenta() bool {
	return c.red == 255 && c.green == 0 && c.blue == 255
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
