package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePostfixFormIncorrect(t *testing.T) {

	_, err := EvaluatePostfix("")
	assert.NotNil(t, err)

	_, err = EvaluatePostfix("dsad")
	assert.NotNil(t, err)

	_, err = EvaluatePostfix("2 ** 2")
	assert.NotNil(t, err)

}

func TestEvaluatePostfixFormCorrect(t *testing.T) {

	res, err := EvaluatePostfix("1488 2 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 1490.0, res)
	}

	res, err = EvaluatePostfix("1337 3 ^ 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 2389979758.0, res)
	}

	res, err = EvaluatePostfix("228 3 - 5 + 5 /")
	if assert.Nil(t, err) {
		assert.Equal(t, 46.0, res)
	}

	res, err = EvaluatePostfix("47 3 + 23 - 3 * 2 ^ 3 / 255 -")
	if assert.Nil(t, err) {
		assert.Equal(t, 1932.0, res)
	}

	res, err = EvaluatePostfix("47 3 + 23 - 3 * 2 ^ 3 / 255 - 5 /")
	if assert.Nil(t, err) {
		assert.Equal(t, 386.4, res)
	}

}

func ExampleEvaluatePostfix() {

	res, _ := EvaluatePostfix("1 2 /")
	fmt.Println(res)

	// Output:
	// 0.5
}
