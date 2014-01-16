package main

import (
	"fmt"
)

type Sudoku [9][9]int

func create() Sudoku {
	s := Sudoku{}
	return s
}

func doCreate() {
}

func createOutput() {
}

func main() {
	sudoku := create()
	for k, v := range sudoku {
		fmt.Printf("key:%d, value:%d\n", k, v)
	}
}
