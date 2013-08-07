#coding: utf-8


import sys, time, copy
from random import choice

class Sudoku:

    def __init__(self):
        line = [0 for i in range(1,10)]
        self.sudoku_init = [copy.deepcopy(line) for i in range(1,10)]

        while(not self.create()):
            pass
        self.out_put(self.sudoku)

    def create(self):
        self.sudoku = copy.deepcopy(self.sudoku_init)

        for i in range(0,9):
            for j in range(0,9):
                if self.sudoku[i][j] == 0:
                    tmp = self._left(i,j)
                    if len(tmp) == 0:
                        return False
                    self.sudoku[i][j] = choice(tmp)
        return True

    def _left(self,x,y):
        left = [i for i in range(1,10)]

        #horizontal check 
        for i in range(0,9):
            if self.sudoku[x][i] in left:
                left.remove(self.sudoku[x][i]) 

        #vertical check 
        for i in range(0,9):
            if self.sudoku[i][y] in left:
                left.remove(self.sudoku[i][y]) 

        #block check
        x_block = x / 3
        y_block = y / 3
        x_start = x_block * 3
        y_start = y_block * 3

        for i in range(x_start,x_start+3):
            for j in range(y_start,y_start+3):
                if self.sudoku[i][j] in left:
                    left.remove(self.sudoku[i][j]) 

        return left

    def is_sudoku(self,sudoku):
        pass

    def out_put(self,sudoku):
        for i in range(1,10):
            for j in range(1,10):
                print "%d" %sudoku[i-1][j-1], 
                print '|' if j%3 == 0 else '',
            print
            if i%3 == 0:
                print 


sudoku = Sudoku()

'''
sudoku = [[0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0],
            [0,0,0,0,0,0,0,0,0]]

'''
