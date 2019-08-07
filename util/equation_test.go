package util

import ( 
	"testing" 
)

const a, b = -12.98, 67.111

func TestAdd(t *testing.T) {
	want := a + b
	if got := Add(a, b); got != want {
		t.Errorf("Add(%f, %f) = %f want %f", a, b, got, want)
	}
}

func TestSubtract(t *testing.T) {
	want := a - b
	if got := Subtract(a, b); got != want {
		t.Errorf("Sub(%f, %f) = %f want %f", a, b, got, want)
	}
}

func TestMultiply(t *testing.T) {
	want := a * b
	if got := Multiply(a, b); got != want {
		t.Errorf("Mul(%f, %f) = %f want %f", a, b, got, want)
	}
}

func TestDivide(t *testing.T) {
	want := a / b
	if got := Divide(a, b); got != want {
		t.Errorf("Div(%f, %f) = %f want %f", a, b, got, want)
	}
}
