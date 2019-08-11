package main

import "testing"

func TestRow(t *testing.T) {
	emptyGrid := [][]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	if row(emptyGrid) {
		t.Error("Empty grid doesn't count as a completed row")
	}
	filledRowDifferent := [][]int{
		{-1, -1, -1},
		{0, 1, 0},
		{-1, -1, -1},
	}
	if row(filledRowDifferent) {
		t.Error("Filled row with different tokens is not complete")
	}
	filledRow := [][]int{
		{-1, -1, -1},
		{0, 0, 0},
		{-1, -1, -1},
	}
	if !row(filledRow) {
		t.Error("A filled row returns true")
	}
}

func TestCol(t *testing.T) {
	emptyGrid := [][]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	if col(emptyGrid) {
		t.Error("An empty grid cannot have a filled column")
	}
	filledColDifferent := [][]int{
		{-1, 0, -1},
		{-1, 1, -1},
		{-1, 0, -1},
	}
	if col(filledColDifferent) {
		t.Error("Filled column with different token is not complete")
	}
	filledCol := [][]int{
		{-1, 1, -1},
		{-1, 1, -1},
		{-1, 1, -1},
	}
	if !col(filledCol) {
		t.Error("A filled column returns true")
	}
}

func TestDiagonal(t *testing.T) {
	emptyGrid := [][]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	if diagonal(emptyGrid) {
		t.Error("Empty grid doesn't have a filled diagonal")
	}
	diagonalFilledDifferent := [][]int{
		{1, -1, -1},
		{-1, 0, -1},
		{-1, -1, 1},
	}
	if diagonal(diagonalFilledDifferent) {
		t.Error("Diagonal filled with different tokens returns false")
	}
	diagonalFilled := [][]int{
		{0, -1, -1},
		{-1, 0, -1},
		{-1, -1, 0},
	}
	if !(diagonal(diagonalFilled)) {
		t.Error("Filled diagonal returns true")
	}

}

func TestAntiDiagonal(t *testing.T) {
	emptyGrid := [][]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	if antiDiagonal(emptyGrid) {
		t.Error("Empty grid doesn't have a filled anti diagonal")
	}
	diagonalFilledDifferent := [][]int{
		{-1, -1, 0},
		{-1, 1, -1},
		{0, -1, -1},
	}
	if antiDiagonal(diagonalFilledDifferent) {
		t.Error("Anti diagonal filled with different tokens returns " +
			"false")
	}
	diagonalFilled := [][]int{
		{-1, -1, 0},
		{-1, 0, -1},
		{0, -1, -1},
	}
	if !(antiDiagonal(diagonalFilled)) {
		t.Error("Filled anti diagonal returns true")
	}

}
