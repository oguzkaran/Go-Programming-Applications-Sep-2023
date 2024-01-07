package numeric

import (
	"fmt"
	"math"
)

type Fraction struct {
	a, b int
}

func (f *Fraction) simplify() {
	minValue := int(min(math.Abs(float64(f.a)), float64(f.b)))

	for i := minValue; i >= 2; i-- {
		if f.a%i == 0 && f.b%i == 0 {
			f.a /= i
			f.b /= i
			break
		}
	}
}

func (f *Fraction) set(a, b int) {
	f.a = a
	f.b = b

	if a != 0 {
		f.setSign()
		f.simplify()
	} else {
		f.b = 1
	}
}

func (f *Fraction) setSign() {
	if f.b < 0 {
		f.a = -f.a
		f.b = -f.b
	}
}

func check(a, b int) int {
	if b == 0 {
		if a == 0 {
			return -1
		}

		return 0
	}

	return 1
}

func add(a1, b1, a2, b2 int) Fraction {
	p, _ := New(a1*b2+a2*b1, b1*b2)

	return *p
}

func subtract(a1, b1, a2, b2 int) Fraction {
	return add(a1, b1, -a2, b2)
}

func multiply(a1, b1, a2, b2 int) Fraction {
	p, _ := New(a1*a2, b1*b2)

	return *p
}

func divide(a1, b1, a2, b2 int) (*Fraction, int) {
	return New(a1*b2, a2*b1)
}

func New(a, b int) (*Fraction, int) {
	status := check(a, b)

	if status != 1 {
		return nil, status
	}

	var f *Fraction = &Fraction{}

	f.set(a, b)

	return f, 1
}

func (f *Fraction) SetNumerator(val int) {
	f.set(val, f.b)
}

func (f *Fraction) SetDenominator(val int) bool {
	if check(f.a, val) != 1 {
		return false
	}

	f.set(f.a, val)

	return true
}

func (f *Fraction) GetNumerator() int {
	return f.a
}

func (f *Fraction) GetDenominator() int {
	return f.b
}

func (f *Fraction) GetRealValue() float64 {
	return float64(f.a) / float64(f.b)
}

func (f *Fraction) Add(other *Fraction) Fraction {
	return add(f.a, f.b, other.a, other.b)
}

func (f *Fraction) AddWith(val int) Fraction {
	return add(f.a, f.b, val, 1)
}

func (f *Fraction) Subtract(other Fraction) Fraction {
	return subtract(f.a, f.b, other.a, other.b)
}

func (f *Fraction) SubtractWith(val int) Fraction {
	return subtract(f.a, f.b, val, 1)
}

func (f *Fraction) Multiply(other Fraction) Fraction {
	return multiply(f.a, f.b, other.a, other.b)
}

func (f *Fraction) MultiplyWith(val int) Fraction {
	return multiply(f.a, f.b, val, 1)
}

func (f *Fraction) Divide(other Fraction) (*Fraction, int) {
	return divide(f.a, f.b, other.a, other.b)
}

func (f *Fraction) DivideWith(val int) (*Fraction, int) {
	return divide(f.a, f.b, val, 1)
}

func (f *Fraction) Inc() {
	f.a += f.b
}

func (f *Fraction) Dec() {
	f.a -= f.b
}

func (f *Fraction) Print() {
	fmt.Printf("%d / %d = %f\n", f.a, f.b, f.GetRealValue())
}
