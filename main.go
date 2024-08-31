package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read tetrominoes from the file
func readTetrominoes(filePath string) ([][][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tetrominoes [][][]string
	var currentTetromino [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentTetromino) > 0 {
				tetrominoes = append(tetrominoes, currentTetromino)
				currentTetromino = nil
			}
		} else {
			currentTetromino = append(currentTetromino, strings.Split(line, ""))
		}
	}

	if len(currentTetromino) > 0 {
		tetrominoes = append(tetrominoes, currentTetromino)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tetrominoes, nil
}

// Validate if each tetromino is 4x4
func validateTetrominoes(tetrominoes [][][]string) bool {
	for _, tetromino := range tetrominoes {
		if len(tetromino) != 4 {
			return false
		}
		for _, row := range tetromino {
			if len(row) != 4 {
				return false
			}
		}
	}
	return true
}

// Initialize a grid
func initializeGrid(size int) [][]string {
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	return grid
}

// Place tetrominoes into the grid (simple placement)
func placeTetrominoes(tetrominoes [][][]string) [][]string {
	gridSize := 10 // Example grid size, adjust as needed
	grid := initializeGrid(gridSize)

	// Simple placement logic
	tetrominoIndex := 'A'
	rowOffset := 0
	colOffset := 0
	for _, tetromino := range tetrominoes {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if tetromino[i][j] == "#" {
					if rowOffset+i < gridSize && colOffset+j < gridSize {
						grid[rowOffset+i][colOffset+j] = string(tetrominoIndex)
					}
				}
			}
		}
		tetrominoIndex++
		colOffset += 5 // Move to next column (example spacing, adjust as needed)
		if colOffset >= gridSize {
			colOffset = 0
			rowOffset += 5 // Move to next row (example spacing, adjust as needed)
		}
	}

	return grid
}

// Print the grid
func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}techniques
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR")
		return
	}

	filePath := os.Args[1]
	tetrominoes, err := readTetrominoes(filePath)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	if !validateTetrominoes(tetrominoes) {
		fmt.Println("ERROR")
		return
	}

	grid := placeTetrominoes(tetrominoes)
	printGrid(grid)
}
