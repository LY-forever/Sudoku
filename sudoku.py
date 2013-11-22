#coding: utf-8


import sys
import time
import copy
from random import choice


def cal(f):
    def wrapped(*args, **kwargs):
        start = time.time()
        f(*args,**kwargs)
        print  "cost %s seconds" %(time.time() - start)
    return wrapped

class Sudoku:

    @cal
    def create(self):
        line = [0 for i in range(1,10)]
        self.sudoku_init = [copy.deepcopy(line) for i in range(1,10)]

        while(not self._do_create(self.sudoku_init)):
            pass
        self._create_output(self.sudoku)

    @cal
    def resolve(self,sudoku):
        self.sudoku_init = sudoku
        self.sudoku = copy.deepcopy(sudoku)
        self._do_resolve(0)

    def _do_create(self,sudoku):
        self.sudoku = copy.deepcopy(sudoku)

        for i in range(0,9):
            for j in range(0,9):
                if self.sudoku[i][j] == 0:
                    tmp = self._left(i,j)
                    if len(tmp) == 0:
                        return False
                    self.sudoku[i][j] = choice(tmp)
        return True

    def _do_resolve(self,pos):
        x = pos / 9
        y = pos % 9

        if pos > 80:
            self._resolve_output(self.sudoku)
            return 

        if self.sudoku[x][y] != 0:
            self._do_resolve(pos+1)
        else:
            left = self._left(x,y)
            for i in left:
                self.sudoku[x][y] = i
                self._do_resolve(pos+1)

            self.sudoku[x][y] = 0

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

    def _create_output(self,sudoku):
        print 
        for i in range(1,10):
            for j in range(1,10):
                print "%d" %sudoku[i-1][j-1], 
                print '|' if j%3 == 0 else '',
            print
            if i%3 == 0:
                print 

    def _resolve_output(self,sudoku):
        print
        for i in range(1,10):
            for j in range(1,10):
                out = "%d" if self.sudoku_init[i-1][j-1] != 0 else "\033[1;31;40m%d\033[0m"
                print out %sudoku[i-1][j-1], 
                print '|' if j%3 == 0 else '',
            print
            if i%3 == 0:
                print 



if __name__ == "__main__":
    sudoku = Sudoku()

    sudoku.create()

    aSudoku = [ [0,3,6,8,0,0,0,0,2],
                [9,0,0,0,5,0,0,3,0],
                [0,0,5,0,0,6,0,0,0],
                [0,2,0,0,0,0,1,0,0],
                [0,0,3,0,8,0,0,5,0],
                [0,1,9,0,0,0,0,0,0],
                [1,0,0,0,0,0,0,9,5],
                [0,0,0,0,0,2,0,0,8],
                [0,0,0,3,9,0,7,0,0] ]

    sudoku.resolve(aSudoku)
