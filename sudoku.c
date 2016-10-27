//
// Created by Neo on 2016/10/24.
//

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

#define TRUE 1
#define FALSE 0

#ifndef SLEN
#define SLEN 9
#endif

#define timesub(start, end, result) \
            do {    \
                (result)->tv_sec = (end)->tv_sec - (start)->tv_sec; \
                (result)->tv_usec = (end)->tv_usec - (start)->tv_usec;  \
            }while(0)

typedef int Sudoku[SLEN][SLEN];
typedef struct {
    int list[SLEN];
    int size;
} Left;

static Sudoku sudokuToSolve;
static struct timeval StartTime;

void timeCost() {
    struct timeval EndTime, result;

    gettimeofday(&EndTime, NULL);
    timesub(&StartTime, &EndTime, &result);

    double diff = result.tv_sec + (1.0 * result.tv_usec)/1000000;

    printf("Cost: %f s\n", diff);
}

void solvedShow(Sudoku solvedSudoku)
{
    printf("\n");
    for (int i=0; i<SLEN; i++) {
        for (int j=0; j<SLEN; j++) {
            if (sudokuToSolve[i][j] == 0)
                printf(" \033[1;31;40m%d\033[0m ", solvedSudoku[i][j]);
            else
                printf(" %d ", solvedSudoku[i][j]);
            if (j%3 == 2 && j != SLEN-1)
                printf("|");
        }
        if (i%3 == 2 && i != SLEN-1)
                printf("\n");
        printf("\n");
    }
    printf("\n");
}

void removeIfIn(int x, Left *left) {

    int *ptr = left->list;
    int exist = FALSE;
    int existAt = 0;
    for (int i = 0; i < left->size; i++) {
        if (*(ptr + i) == x) {
            exist = TRUE;
            existAt = i;
        }
    }
    if (exist) {
        for (int i = existAt; i < left->size; i++) {
            *(ptr + i) = *(ptr + i + 1);
        }
        left->size--;
    }
}

int randomFromLeft(Left *left)
{
    int randomInt;
    randomInt = rand() % left->size;
    return *(left->list + randomInt);
}

Left getLeft(Sudoku sudoku, int x, int y)
{
    Left left = {{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9};

    //horizontal check
    for (int i=0; i<SLEN; i++) {
        removeIfIn(sudoku[x][i], &left);
    }

    //vertical check
    for (int j=0; j<SLEN; j++) {
        removeIfIn(sudoku[j][y], &left);
    }

    //block check
    int x_block = x / 3;
    int y_block = y / 3;
    int x_start =  x_block * 3;
    int y_start = y_block * 3;

    for (int i=x_start; i<x_start+3; i++) {
        for (int j=y_start; j<y_start+3; j++) {
            removeIfIn(sudoku[i][j], &left);
        }
    }

    return left;
}


int solve(Sudoku *sudoku, int pos)
{
    Sudoku toSolve;
    memcpy(toSolve, sudoku, sizeof(Sudoku));

    int x = pos / 9;
    int y = pos % 9;
    Left left;

    if (pos > 80) {
        solvedShow(toSolve);
        return TRUE;
    }

    if (toSolve[x][y] != 0)
        solve(&toSolve, pos+1);
    else {
        left = getLeft(toSolve, x, y);
        if (left.size == 0)
            return FALSE;
        else {
            for (int i=0; i<left.size; i++) {
                toSolve[x][y] = *(left.list+i);
                solve(&toSolve, pos+1);
            }
        }
    }
    return FALSE;
}

int main()
{
    gettimeofday(&StartTime, NULL);
    atexit(timeCost);
    Sudoku toSolve = {
        {0, 3, 6, 8, 0, 0, 0, 0, 2},
        {9, 0, 0, 0, 5, 0, 0, 3, 0},
        {0, 0, 5, 0, 0, 6, 0, 0, 0},
        {0, 2, 0, 0, 0, 0, 1, 0, 0},
        {0, 0, 3, 0, 8, 0, 0, 5, 0},
        {0, 1, 9, 0, 0, 0, 0, 0, 0},
        {1, 0, 0, 0, 0, 0, 0, 9, 5},
        {0, 0, 0, 0, 0, 2, 0, 0, 8},
        {0, 0, 0, 3, 9, 0, 7, 0, 0}
    };
    memcpy(sudokuToSolve, toSolve, sizeof(Sudoku));

    solve(&toSolve, 0);

    return 0;
}

