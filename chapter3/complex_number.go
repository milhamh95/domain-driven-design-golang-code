package chapter3

type ComplexNumber struct {
	realPart    float64
	complexPart float64
}

func NewComplexNumber(realPart, complexPart float64) ComplexNumber {
	return ComplexNumber{
		realPart:    realPart,
		complexPart: complexPart,
	}
}

func (c ComplexNumber) Add(complexNumber ComplexNumber) ComplexNumber {
	return ComplexNumber{
		realPart:    c.realPart + complexNumber.realPart,
		complexPart: c.complexPart + complexNumber.complexPart,
	}
}
