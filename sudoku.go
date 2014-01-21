package main

import (
	"fmt"
//	"math/rand"
)

type Sudoku [9][9]int

func create() Sudoku {
	s := Sudoku{}
	return s
}

func doCreate(s Sudoku) {
}

func createOutput(sudoku Sudoku) {
	for _, v := range sudoku {
		fmt.Printf("%d\n", v)
	}
}

func solve() {
    c := make(chan Sudoku,10)
    //sudoku := create()
    go doSolve(c) 
    for i := range c{
        createOutput(i)
    }
}

func doSolve(c chan Sudoku) {
    sudoku := create()
    c <- sudoku
    close(c)
}

func _left(sudoku Sudoku, x,y int) {
}

func main() {
}
