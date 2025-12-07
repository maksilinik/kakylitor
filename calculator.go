package calc

import (
	"fmt"
	"strconv"
)

// Calculator представляет простой калькулятор
type Calculator struct {
	Result float64
}

// NewCalculator создает новый калькулятор
func NewCalculator() *Calculator {
	return &Calculator{Result: 0}
}

// Add складывает числа
func (c *Calculator) Add(x float64) {
	c.Result += x
}

// Subtract вычитает числа
func (c *Calculator) Subtract(x float64) {
	c.Result -= x
}

// Multiply умножает числа
func (c *Calculator) Multiply(x float64) {
	c.Result *= x
}

// Divide делит числа
func (c *Calculator) Divide(x float64) error {
	if x == 0 {
		return fmt.Errorf("ошибка: деление на ноль")
	}
	c.Result /= x
	return nil
}

// Clear сбрасывает результат
func (c *Calculator) Clear() {
	c.Result = 0
}

// GetResult возвращает текущий результат
func (c *Calculator) GetResult() float64 {
	return c.Result
}

// Calculate выполняет операцию на основе ввода пользователя
func (c *Calculator) Calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("ошибка: деление на ноль")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("неизвестная операция: %s", operator)
	}
}

// ParseInput парсит ввод пользователя
func ParseInput(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}