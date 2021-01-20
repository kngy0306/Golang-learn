package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][columns]Cell

// 発生する可能性のあるエラー群
var (
	ErrBounds = errors.New("範囲外の数値が入力されました。")
	ErrDigit  = errors.New("invaliid digit")
	// ErrInRow = errors.New("digit already present in this row")
	// ErrInCol = errors.New("digit already present in this column")
	// ErrInRegion = errors.New("digit already present in this region")
	// ErrFixedDigit = errors.New("intitial digits cannot be overwritten")
)

func NewSudoku(digit [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digit[r][c]
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}

	return &grid
}

func (g *Grid) Show() {
	for c := 0; c < rows; c++ {
		for r := 0; r < columns; r++ {
			switch {
			case g[c][r].digit == empty:
				fmt.Printf("  ")
			default:
				fmt.Printf("%v ", g[c][r].digit)
			}
			if r == 2 || r == 5 || r == 8 {
				fmt.Print("| ")
			}

		}
		if c == 2 || c == 5 || c == 8 {
			fmt.Printf("\n- - - - - - - - - - - -  ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Grid) Set(row, column int, digit int8) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	}

	g[row][column].digit = digit
	return nil
}

func inBounds(r, c int) bool {
	if r > 8 || r < 0 || c > 8 || c < 0 {
		return false
	}
	return true
}

func validDigit(d int8) bool {
	return d >= 0 && d <= 8
}

func main() {
	// var sudoku *Grid
	sudoku := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	sudoku.Show()

	err := sudoku.Set(1, 1, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sudoku.Show()
}
