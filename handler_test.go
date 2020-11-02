package lab2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type outputTest struct {
	called bool
}

func (o *outputTest) Write(p []byte) (n int, err error) {
	o.called = true
	return 0, nil
}

func TestComputerHandleCorrect(t *testing.T) {
	expression := "4 2 - 3 * 5 +"
	firstTest := strings.NewReader(expression)
	output := outputTest{}

	handler := ComputeHandler{
		Input:  firstTest,
		Output: &output,
	}

	handler.Compute()
	assert.Equal(t, true, output.called)
}

func TestComputerHandleIncorrect(t *testing.T) {
	expression := ""
	secondTest := strings.NewReader(expression)
	output := outputTest{}

	handler := ComputeHandler{
		Input:  secondTest,
		Output: &output,
	}

	err := handler.Compute()
	assert.NotNil(t, err)
}
