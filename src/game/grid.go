package game

import (
	"fmt"
	"strings"
)

type Grid struct {
	grid [][]bool
	width int
	height int
}

func (g *Grid) initGrid(width int, height int) {
	g.grid = make([][]bool, height)
	for i := range g.grid {
		g.grid[i] = make([]bool, width)
	}
	g.width = width
	g.height = height
}

func (g Grid) printGrid() {
	// Print array
	for i := range g.grid {
		var row strings.Builder

		for j := range g.grid[i] {
			if g.grid[i][j] {
				row.WriteString("1 ")
			} else {
				row.WriteString("0 ")
			}
		}

		fmt.Println(row.String())
	}
	fmt.Println("======================")
}

func (g *Grid) toggleCell(row int, col int) {
	if g.isInBounds(row, col) == false {
		return
	}
	g.grid[row][col] = !g.grid[row][col]
}

func (g *Grid) isInBounds(row int, col int) bool {
	if row < 0 || row >= g.height || col >= g.width || col < 0{
		return false
	}
	return true
}

func (g *Grid) countNeighbors(row int, col int) int {
	// given a cell position at row and col, we need to count the number of alive cells in the 8 surrounding cells
	if g.isInBounds(row, col) == false {
		return 0
	}
	neighborsCount := 0
	for i := row - 1; i <= row + 1; i++ {
		for j := col - 1; j <= col + 1; j++ {
			if i == row && j == col {
				continue
			}
			if g.isInBounds(i, j)  {
				if g.grid[i][j] {
					neighborsCount += 1
				}
			}
			
		}
	}
	return neighborsCount
}

func (g *Grid) copyGrid(source Grid) {
	for i := range source.grid {
		copy(g.grid[i], source.grid[i])
	}
}


func (g *Grid) incrementStep() {
	newGrid := Grid{}
	newGrid.initGrid(g.width, g.height)
	newGrid.copyGrid(*g)
	// Iterate over grid
	for row := range g.grid {
		for col := range g.grid[row] {

			// live cell logic
			if g.grid[row][col] {
				// for each cell, calculate number of neighbors and return true for alive or false for death
				if g.countNeighbors(row, col) < 2 || g.countNeighbors(row, col) > 3{
					newGrid.grid[row][col] = false // dead because less than 2 neighbors (underpopulation) or more than 3 neighbors (overpopulation)
				}
			} else {
				// if cell is dead,
				if g.countNeighbors(row, col) >= 3 {
					newGrid.grid[row][col] = true // cell comes alive from reproduction
				}
			}
		}
	}
	g.copyGrid(newGrid)
	g.printGrid()
}
