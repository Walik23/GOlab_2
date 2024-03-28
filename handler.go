package lab2

import (
	"fmt"
	"io"
	"strings"
)

type InfixInterpreter interface {
	EvaluateInfix(expression string) (int, error)
}

type SpyInfixInterpreter struct {
	spyResult int
	spyError  error
}

func (spc *SpyInfixInterpreter) EvaluateInfix(expression string) (int, error) {
	return spc.spyResult, spc.spyError
}

type ComputeHandler struct {
	Input      io.Reader
	Output     io.Writer
	Calculator InfixInterpreter
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}
	expression := strings.TrimSpace(string(data))
	result, err := ch.Calculator.EvaluateInfix(expression)
	if err != nil {
		return err
	}
	_, err = ch.Output.Write([]byte(fmt.Sprintf("%d\n", result)))
	return err
}
