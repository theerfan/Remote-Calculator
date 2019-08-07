package util

import (
	"errors"
	"fmt"
	"reflect"
)

//Equation is a representation of an equation.
type Equation struct {
	A   float64
	Fix string
	B   float64
}

//SetField 1212121212  dsdsd
func SetField(strct interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(strct).Elem()
	structFieldValue := structValue.FieldByName(name)

	// Put breakpoint here!

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

//FillStruct 123
func FillStruct(s *Equation, m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

//Calculate calculates
func Calculate(eq Equation) (ans float64) {
	switch eq.Fix {
	case "+":
		ans = eq.A + eq.B
	case "-":
		ans = eq.A - eq.B
	case "*":
		ans = eq.A * eq.B
	case "/":
		ans = eq.A / eq.B
	}
	return
}
