package util

import (
	"errors"
	"fmt"
	"reflect"
)

//Equation is a representation of an equation.
type Equation struct {
	A, B   float64
	Fix    string
}

//SetField Copied code from the internet, have no idea how it works, debugging didnt work and I have to read more on relfection later.
func SetField(strct interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(strct).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

//FillStruct fills a struct with a map, matching the struct's feld names with the map's keys.
func FillStruct(s *Equation, m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

//Calculate calculates the answer. EZ
func Calculate(eq Equation) (float64) {
	var ans float64
	a, b := eq.A, eq.B
	switch eq.Fix {
	case "+":
		ans = Add(a, b)
	case "-":
		ans = Subtract(a, b)
	case "*":
		ans = Multiply(a, b)
	case "/":
		ans = Divide(a, b)
	}
	return ans
}

func Add(a float64, b float64) float64 {
	return a + b
}

func Subtract(a float64, b float64) float64 {
	return a - b
}

func Multiply(a float64, b float64) float64 {
	return a * b
}

func Divide(a float64, b float64) float64 {
	return a / b
}