package main

import (
	"fmt"
	"math/rand"
)

type Sudoku [9][9]int
type Left []int

func create() Sudoku {
	sudoku := Sudoku{}
	for {
		sudokuResult, isSucc := doCreate(sudoku)
		if isSucc == true {
			return sudokuResult
		}
	}
}

func doCreate(sudoku Sudoku) (Sudoku, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			tmp := left(sudoku, i, j)
			tmpLen := len(tmp)
			if tmpLen == 0 {
				return sudoku, false
			}
			x := rand.Intn(tmpLen)
			sudoku[i][j] = tmp[x]
		}
	}
	return sudoku, true
}

func createOutput(sudoku Sudoku) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf(" %d ", sudoku[i][j])
			if j%3 == 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i%3 == 2 {
			fmt.Println()
		}
	}
}

func solve(sudoku Sudoku) Sudoku {
	c := make(chan Sudoku)
	go doSolve(sudoku, 0, c)
	solved := <-c
	return solved
}

func doSolve(sudoku Sudoku, pos int, c chan Sudoku) bool {

	x := pos / 9
	y := pos % 9

	if pos > 80 {
		c <- sudoku
		close(c)
		return true
	}

	if sudoku[x][y] != 0 {
		doSolve(sudoku, pos+1, c)
	} else {
		left := left(sudoku, x, y)
		if len(left) == 0 {
			sudoku[x][y] = 0
		} else {
			for _, i := range left {
				sudoku[x][y] = i
				doSolve(sudoku, pos+1, c)
			}
		}
	}
	return false
}

func solveOutput(sudoku, solved Sudoku) {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var out string
			if sudoku[i][j] == 0 {
				out = " \033[1;31;40m%d\033[0m "
			} else {
				out = " %d "
			}
			fmt.Printf(out, solved[i][j])
			if j%3 == 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i%3 == 2 {
			fmt.Println()
		}
	}
}

func left(sudoku Sudoku, x, y int) Left {
	left := Left{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//horizontal check
	for i := 0; i < 9; i++ {
		if isInList(sudoku[x][i], left) {
			left = removeFromList(sudoku[x][i], left)
		}
	}

	//vertical check
	for j := 0; j < 9; j++ {
		if isInList(sudoku[j][y], left) {
			left = removeFromList(sudoku[j][y], left)
		}
	}

	//block check
	x_block := x / 3
	y_block := y / 3
	x_start := x_block * 3
	y_start := y_block * 3

	for i := x_start; i < x_start+3; i++ {
		for j := y_start; j < y_start+3; j++ {
			if isInList(sudoku[i][j], left) {
				left = removeFromList(sudoku[i][j], left)
			}
		}
	}

	return left
}

func isInList(i int, list Left) bool {
	for _, v := range list {
		if i == v {
			return true
		}
	}
	return false
}

func removeFromList(i int, list Left) Left {
	newList := Left{}
	for _, v := range list {
		if v != i {
			newList = append(newList, v)
		}
	}
	return newList
}

func main() {
	s := create()
	createOutput(s)

	fmt.Println()

	aSudoku := Sudoku{
		{0, 3, 6, 8, 0, 0, 0, 0, 2},
		{9, 0, 0, 0, 5, 0, 0, 3, 0},
		{0, 0, 5, 0, 0, 6, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 3, 0, 8, 0, 0, 5, 0},
		{0, 1, 9, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 9, 5},
		{0, 0, 0, 0, 0, 2, 0, 0, 8},
		{0, 0, 0, 3, 9, 0, 7, 0, 0},
	}
	solved := solve(aSudoku)
	solveOutput(aSudoku, solved)
}
