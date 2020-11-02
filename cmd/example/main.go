package main

import (
	"flag"
	"os"
	"io"
	"log"
	"strings"
	lab2 "github.com/KPI-KMD/Labochka_2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f","","File with expression")
	resultFile = flag.String("o","","File with calculation result")
)

func main() {
	flag.Parse()
	var ans io.Reader
	var wr io.Writer
	var err error
	
	if (*inputFile == "" && *inputExpression == "") {
		l := log.New(os.Stderr, "", 0)
		l.Println("ERROR: USE -e to enter an expression to compute; \n\t   -f to enter location of a file containing an expression.")
		os.Exit(3)
	}

	if *inputFile != "" {
		ans, err = os.Open(*inputFile);
		if err != nil {
			l := log.New(os.Stderr, "", 0)
			l.Println("ERROR: ", err)
			os.Exit(3)
		}
	} else {
		ans = strings.NewReader(*inputExpression)
	}

	if(*resultFile != "") {
		wr, err = os.Create(*resultFile);
		if err != nil {
			l := log.New(os.Stderr, "", 0)
			l.Println("ERROR: ", err)
			os.Exit(3)
		}
	} 


	handler := &lab2.ComputeHandler{
	    Input: ans,
	    Output: wr,	
	}
	err = handler.Compute()	
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Println("ERROR: ", err)
		os.Exit(3)
	}
}
