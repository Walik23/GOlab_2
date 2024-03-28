package lab2

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_Computes(t *testing.T) {
	type fields struct {
		Input  io.Reader
		Output io.Writer
	}

	var writer bytes.Buffer
	mockErrorMessage := "mockError"

	tests := []struct {
		name    string
		fields  fields
		calc    SpyPostfixCalculator
		wantErr bool
	}{
		{"HappyPathCase", fields{strings.NewReader("7 2 -"), &writer}, SpyPostfixCalculator{7, nil}, false},
		{"HappyPathCase", fields{strings.NewReader("7 2 - - -"), &writer}, SpyPostfixCalculator{0, errors.New(mockErrorMessage)}, true},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			ch := &ComputeHandler{
				Input:      testCase.fields.Input,
				Output:     testCase.fields.Output,
				Calculator: &testCase.calc,
			}
			err := ch.Compute()

			if testCase.wantErr {
				assert.Error(t, err, mockErrorMessage)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
