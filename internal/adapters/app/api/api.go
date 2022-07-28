package api

import "github.com/PrinceDavis/hex/internal/ports"

type Adapter struct {
	db    ports.DBPort
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DBPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetSubstraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "substraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := api.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
