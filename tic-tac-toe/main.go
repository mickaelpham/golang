package main

import "fmt"

func main() {
	fmt.Println("Welcome to Tic-Tac-Toe")
	const numPlayers = 2
	var input, row, col int
	grid := [][]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	currentPlayer := 0
	for nextTurn(grid) {
		printGrid(grid)
		fmt.Printf("[Player %d] Place token at: ", currentPlayer+1)
		fmt.Scanf("%d", &input)
		if input <= 0 || input > len(grid)*len(grid) {
			fmt.Println("Invalid input! Try again.")
			continue
		}
		row = (input - 1) / len(grid)
		col = (input - 1) % len(grid)
		if grid[row][col] != -1 {
			fmt.Println("This cell is already occupied")
			continue
		}
		grid[row][col] = currentPlayer
		currentPlayer = (currentPlayer + 1) % numPlayers
	}
}

func printGrid(grid [][]int) {
	characters := []rune{'X', 'O'}
	fmt.Println("+---+---+---+")
	for row, cells := range grid {
		fmt.Print("|")
		for col, cell := range cells {
			if cell == -1 {
				fmt.Printf(" %d |", row*len(grid)+col+1)
			} else {
				fmt.Printf(" %c |", characters[cell])
			}
		}
		fmt.Println("\n+---+---+---+")
	}
}

func nextTurn(grid [][]int) bool {
	return !(victory(grid) || gridFull(grid))
}

func gridFull(grid [][]int) bool {
	for _, row := range grid {
		for _, cell := range row {
			if cell == -1 {
				return false
			}
		}
	}
	fmt.Println("The grid is full. DRAW!!")
	printGrid(grid)
	return true
}

func victory(grid [][]int) bool {
	win := row(grid) || col(grid) || diagonal(grid) || antiDiagonal(grid)
	if win {
		fmt.Println("Victory!")
		printGrid(grid)
	}
	return win
}

func row(grid [][]int) bool {
	for _, row := range grid {
		firstCell := row[0]
		if firstCell == -1 {
			continue
		}
		uniq := true
		for _, cell := range row {
			if cell != firstCell {
				uniq = false
			}
		}
		if uniq {
			return true
		}
	}
	return false
}

func col(grid [][]int) bool {
	for col := 0; col < len(grid[0]); col++ {
		firstCell := grid[0][col]
		if firstCell == -1 {
			continue
		}
		uniq := true
		for _, row := range grid {
			if row[col] != firstCell {
				uniq = false
			}
		}
		if uniq {
			return true
		}
	}
	return false
}

// This method assumes a square matrix
func diagonal(grid [][]int) bool {
	firstCell := grid[0][0]
	if firstCell == -1 {
		return false
	}
	for i := 0; i < len(grid); i++ {
		if firstCell != grid[i][i] {
			return false
		}
	}
	return true
}

func antiDiagonal(grid [][]int) bool {
	firstCell := grid[len(grid)-1][0]
	if firstCell == -1 {
		return false
	}
	for row, col := len(grid)-1, 0; row >= 0; row, col = row-1, col+1 {
		if firstCell != grid[row][col] {
			return false
		}
	}
	return true
}
