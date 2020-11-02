package lab2


import (
	"fmt"
	"io"
	"os"
	"log"
)

type ComputeHandler struct {
	Input io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	expression := ""
	p := make([]byte, 8)
	for {
		n, err := (ch.Input).Read(p)
		if err != nil{
		    if err == io.EOF {
			fmt.Println(string(p[:n]))
			break
		    }
		    fmt.Println(err)
		    os.Exit(1)
		}
		expression += string(p[:n])
	}

	res, err := EvaluatePostfix(expression)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	if ch.Output == nil {
		l := log.New(os.Stdout, "", 0)
		l.Println("Result: ", res)
	} else {
		resString := fmt.Sprintf("%f", res)
		(ch.Output).Write([]byte(resString))
	}
	return nil
}
