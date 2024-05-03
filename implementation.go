package lab2

import (
	"errors"
	"strconv"
	"strings"
)



type DefaultInfixInterpreter struct {
}

func (dpc *DefaultInfixInterpreter) isOperator(op string) bool {
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}
	return operators[op]
}

func (dpc *DefaultInfixInterpreter) ToInfix(prefix string) (string, error) {
	if prefix == "" {
		return "", errors.New("empty string")
	}

	tokens := strings.Fields(prefix)
	var stack []string
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if dpc.isOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("expression format is incorrect")
			}
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			infixExpr := "(" + operand1 + token + operand2 + ")"
			stack = append(stack, infixExpr)
		} else {
			_, err := strconv.Atoi(token)
			if err != nil {
				return "", errors.New("invalid character: " + token)
			}
			stack = append(stack, token)
		}
	}
	if len(stack) != 1 {
		return "", errors.New("expression format is incorrect")
	}
	return stack[0], nil
}
