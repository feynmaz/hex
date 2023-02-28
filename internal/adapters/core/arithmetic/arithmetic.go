package arithmetic

import "errors"

var (
	ErrDivideByZero = errors.New("division by zero")
)

// Adapter is an implementation of Arithmetic port
type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arithmetic Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arithmetic Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arithmetic Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arithmetic Adapter) Division(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
