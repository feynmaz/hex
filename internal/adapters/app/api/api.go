package api

import (
	"github.com/feynmaz/hex/internal/ports"
)

// Adapter is an implementation of API interface port
type Adapter struct {
	db         ports.DB
	arithmetic ports.Arithmetic
}

func NewAdapter(db ports.DB, arithmetic ports.Arithmetic) *Adapter {
	return &Adapter{db: db, arithmetic: arithmetic}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Addition(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "addition"); err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "subtraction"); err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "multiplication"); err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arithmetic.Division(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "division"); err != nil {
		return 0, err
	}

	return answer, nil
}
