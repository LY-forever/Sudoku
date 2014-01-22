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
        sudokuResult,isSucc := doCreate(sudoku)
        if isSucc == true {
            return sudokuResult
        }
    }
}

func doCreate(sudoku Sudoku) (Sudoku, bool) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            tmp := left(sudoku,i,j)
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

func left(sudoku Sudoku, x,y int) Left {
    left := Left {1,2,3,4,5,6,7,8,9}

    //horizontal check 
    for i := 0; i < 9; i++ {
        if isInList(sudoku[x][i], left) {
            left = removeFromList(sudoku[x][i],left)
        }
    }

    //vertical check 
    for j := 0; j < 9; j++ {
        if isInList(sudoku[j][y], left) {
            left = removeFromList(sudoku[j][y],left)
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
                left = removeFromList(sudoku[i][j],left)
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
            newList = append(newList,v)
        }
    }
    return newList
}

func main() {
    s := create()
    createOutput(s)
}



