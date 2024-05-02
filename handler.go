package lab2

import (
	"fmt"
	"io"
	"strings"
)

type InfixInterpreter interface {
	ToInfix(expression string) (string, error)
}

type SpyInfixInterpreter struct {
	spyResult string
	spyError  error
}

func (spc *SpyInfixInterpreter) ToInfix(expression string) (string, error) {
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
	result, err := ch.Calculator.ToInfix(expression)
	if err != nil {
		return err
	}
	_, err = ch.Output.Write([]byte(fmt.Sprintf("%s\n", result)))
	return err
}
