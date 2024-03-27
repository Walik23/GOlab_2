package lab2

import (
	"io"
)

// Its Compute() method should read the expression from input and write the computed result to the output.

// ComputeHandler обробляє обчислення виразу
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}
// Compute обчислює вираз та записує результат
func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	return nil
}
