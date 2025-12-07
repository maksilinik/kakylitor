package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator представляет простой калькулятор
type Calculator struct {
	Result float64
}

// NewCalculator создает новый калькулятор
func NewCalculator() *Calculator {
	return &Calculator{Result: 0}
}

// Calculate выполняет операцию
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

// Clear сбрасывает результат
func (c *Calculator) Clear() {
	c.Result = 0
}

// GetResult возвращает текущий результат
func (c *Calculator) GetResult() float64 {
	return c.Result
}

func main() {
	calculator := NewCalculator()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== ПРОСТОЙ КАЛЬКУЛЯТОР ===")
	fmt.Println("Доступные операции: +, -, *, /")
	fmt.Println("Введите 'clear' для сброса")
	fmt.Println("Введите 'exit' для выхода")
	fmt.Println("============================")

	for {
		fmt.Printf("\nТекущий результат: %.2f\n", calculator.GetResult())
		fmt.Print("Введите выражение (например: 5 + 3): ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		// Проверка на выход
		if strings.ToLower(input) == "exit" {
			fmt.Println("До свидания!")
			break
		}

		// Проверка на сброс
		if strings.ToLower(input) == "clear" {
			calculator.Clear()
			fmt.Println("Результат сброшен до 0")
			continue
		}

		// Парсинг ввода
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Ошибка: введите выражение в формате 'число оператор число'")
			continue
		}

		// Парсинг первого числа
		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Printf("Ошибка парсинга первого числа: %v\n", err)
			continue
		}

		// Парсинг оператора
		operator := parts[1]
		if !isValidOperator(operator) {
			fmt.Println("Ошибка: неверный оператор. Используйте +, -, *, /")
			continue
		}

		// Парсинг второго числа
		num2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Printf("Ошибка парсинга второго числа: %v\n", err)
			continue
		}

		// Вычисление
		result, err := calculator.Calculate(num1, num2, operator)
		if err != nil {
			fmt.Printf("Ошибка вычисления: %v\n", err)
			continue
		}

		// Обновление результата калькулятора
		calculator.Clear()
		calculator.Result = result

		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
	}
}

// isValidOperator проверяет валидность оператора
func isValidOperator(op string) bool {
	validOps := []string{"+", "-", "*", "/"}
	for _, validOp := range validOps {
		if op == validOp {
			return true
		}
	}
	return false
}
