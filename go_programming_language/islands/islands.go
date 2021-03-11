package main

import "fmt"

func main() {
}

type tile struct {
	position [2]int
	value    byte
}

func buildBoard(grid [][]byte) [][]tile {
	board := [][]tile{}
	for rowIdx, row := range grid {
		board = append(board, []tile{})
		for colIdx, val := range row {
			tileToAdd := tile{
				position: [2]int{rowIdx, colIdx},
				value:    val,
			}
			board[rowIdx] = append(board[rowIdx], tileToAdd)
		}
	}
	return board
}

func buildGraph(board [][]tile) map[tile][]tile {
	output := make(map[tile][]tile)
	for rowIdx, tileRow := range board {
		for colIdx, t := range tileRow {
			// add adacent tiles in column
			for i := -1; i < 2; i += 2 {
				if rowIdx+i >= 0 && rowIdx+i < len(board) {
					output[t] = append(output[t], board[rowIdx+i][colIdx])
				}
			}

			// add adacent tiles in row
			for i := -1; i < 2; i += 2 {
				if colIdx+i >= 0 && colIdx+i < len(board[0]) {
					output[t] = append(output[t], board[rowIdx][colIdx+i])
				}
			}
		}
	}
	return output
}

func numIslands(grid [][]byte) int {
	board := buildBoard(grid)
	graph := buildGraph(board)
	islands := [][]tile{}
	searchedTiles := make(map[tile]bool)
	for _, tileRow := range board {
		for _, t := range tileRow {
			if t.value == '1' && !searchedTiles[t] {
				fmt.Println(t)
				island := bfsForIsland(t, graph)
				fmt.Println(island)
				islands = append(islands, island)
				for _, st := range island {
					searchedTiles[st] = true
				}
			}
		}
	}
	return len(islands)
}

func bfsForIsland(t tile, graph map[tile][]tile) []tile {
	island := []tile{t}
	queue := []tile{t}
	for len(queue) > 0 {
		currentTile := queue[0]
		queue = queue[1:]
		adjacentTiles := graph[currentTile]
		for _, adjacentTile := range adjacentTiles {
			if adjacentTile.value == '1' {
				addToQueue := true
				for _, existingTile := range island {
					if existingTile == adjacentTile {
						addToQueue = false
					}
				}
				if addToQueue {
					queue = append(queue, adjacentTile)
				}
				island = append(island, adjacentTile)
			}
		}
	}

	return island
}
