package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInfix(t *testing.T) {
	calculator := DefaultInfixInterpreter{}

	tests := []struct {
		name       string
		expression string
		expected   string
		wantErr    bool
	}{
		{"SimpleAddition", "+ 1 4", "(1+4", false},
		{"SimplePow", "^ 3 4", "(3^4)", false},
		{"SimpleExpression", "/ - 3 4 + 2 7", "((3-4)/(2+7))", false},
		{"EmptyString", " ", "expression format is incorrect", true},
		{"InvalidSymbol", "$ 3 4", "invalid character: $", true},
		{"OneOperand", "5", "5", false},
		{"ComplexExpression", "+ - 1 2 / 3 4 ", "((1-2)+(3/4))", false},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := calculator.ToInfix(testCase.expression)

			if testCase.wantErr {
				assert.Error(t, err, "ToInfix() should return an error")
			} else {
				assert.NoError(t, err, "ToInfix() should not return an error")
				assert.Equal(t, testCase.expected, got, "ToInfix() returned incorrect result")
			}
		})
	}
}

func ExampleDefaultInfixInterpreter_ToInfix() {
	calculator := DefaultInfixInterpreter{}

	// Example 1: SimpleAddition
	result, err := calculator.ToInfix("+ 1 4")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)

	// Example 2: SimplePow
	result, err = calculator.ToInfix("+ - 1 2 / 3 4 ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)

	// Example 3: SimpleExpression
	result, err = calculator.ToInfix("/ - 3 4 + 2 7 ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)

	// Example 4: EmptyString
	result, err = calculator.ToInfix(" ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example 5: InvalidSymbol
	result, err = calculator.ToInfix("$ 3 4")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example 6: OneOperand
	result, err = calculator.ToInfix("5")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example : ComplexExpression
	result, err = calculator.ToInfix("+ - 1 2 / 3 4")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)

	// Output:
	// (1+4)
	// ((1-2)+(3/4))
	// ((3-4)/(2+7))
	// Error: expression format is incorrect
	// Error: invalid character: $
	// ((1-2)+(3/4))
}

func BenchmarkToInfix(b *testing.B) {
	calculator := DefaultInfixInterpreter{}
	expression := "+ 1 4" // Задайте тут ваш вираз для перевірки швидкості
	for i := 0; i < b.N; i++ {
		_, _ = calculator.ToInfix(expression)
	}
}
