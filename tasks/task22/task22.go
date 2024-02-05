package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Пример значений переменных a и b
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(21), nil) // 2^21
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(22), nil) // 2^22

	// Умножение
	multiplicationResult := new(big.Int).Mul(a, b)
	fmt.Printf("Multiplication: %s\n", multiplicationResult)

	// Деление
	divisionResult := new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b))
	fmt.Printf("Division: %s\n", divisionResult)

	// Сложение
	additionResult := new(big.Int).Add(a, b)
	fmt.Printf("Addition: %s\n", additionResult)

	// Вычитание
	subtractionResult := new(big.Int).Sub(a, b)
	fmt.Printf("Subtraction: %s\n", subtractionResult)

}
