package api

import (
	"github.com/feynmaz/hex/internal/ports"
)

// Adapter is an implementation of API interface port
type Adapter struct {
	arithmetic ports.Arithmetic
}

func NewAdapter(arithmetic ports.Arithmetic) *Adapter {
	return &Adapter{arithmetic: arithmetic}
}

func (api Adapter) GetAddition(a, b int32) (int32, error) {
	return api.arithmetic.Addition(a, b)
}

func (api Adapter) GetSubtraction(a, b int32) (int32, error) {
	return api.arithmetic.Subtraction(a, b)
}

func (api Adapter) GetMultiplication(a, b int32) (int32, error) {
	return api.arithmetic.Multiplication(a, b)
}

func (api Adapter) GetDivision(a, b int32) (int32, error) {
	return api.arithmetic.Division(a, b)
}
