package lab2

import (
	"bytes"
	"testing"
)

func TestComputeHandler(t *testing.T) {
	// Тест успішного обчислення
	input := bytes.NewBufferString("+ 2 2")
	output := &bytes.Buffer{}
	handler := ComputeHandler{Input: input, Output: output}
	err := handler.Compute()
	if err != nil {
		t.Errorf("Неочікувана помилка: %v", err)
	}
	if output.String() != "4\n" {
		t.Errorf("Невірний результат: %s", output.String())
	}

	// Тест помилки синтаксису
	input = bytes.NewBufferString(" + 2 2 ")
	output = &bytes.Buffer{}
	handler = ComputeHandler{Input: input, Output: output}
	err = handler.Compute()
	if err == nil {
		t.Error("Очікувалась помилка синтаксису")
	}
}
