package main

import "strings"

func main() {

}

type tile struct {
	position [2]int
	value    string
}

var adjacentIndicies = []int{-1, 1}

var found = false

func buildBoardGraph(board [][]byte) map[tile][]tile {
	output := make(map[tile][]tile)
	colCount := len(board[0])
	rowCount := len(board)
	for rowIdx, row := range board {
		for colIdx, col := range row {
			newTile := tile{
				position: [2]int{rowIdx, colIdx},
				value:    string(col),
			}

			if _, ok := output[newTile]; !ok {
				output[newTile] = []tile{}
			}

			// Add adjacent Tiles in column
			for _, adjacentIndex := range adjacentIndicies {
				indexToAdd := rowIdx + adjacentIndex
				if indexToAdd >= 0 && indexToAdd < rowCount {
					adjacentTile := tile{
						position: [2]int{indexToAdd, colIdx},
						value:    string(board[indexToAdd][colIdx]),
					}
					output[newTile] = append(output[newTile], adjacentTile)
				}
			}

			// Add adjacent Tiles in row
			for _, adjacentIndex := range adjacentIndicies {
				indexToAdd := colIdx + adjacentIndex
				if indexToAdd >= 0 && indexToAdd < colCount {
					adjacentTile := tile{
						position: [2]int{rowIdx, indexToAdd},
						value:    string(board[rowIdx][indexToAdd]),
					}
					output[newTile] = append(output[newTile], adjacentTile)
				}
			}
		}
	}
	return output
}

func exist(board [][]byte, word string) bool {
	// should probably have a channel that my dfs function can send on
	// and my exist function can receive on to control this variable
	// instead of using a global
	found = false

	if len(word) == 0 {
		return found
	}

	graph := buildBoardGraph(board)
	firstLetter := word[0:1]

	if len(word) == 1 {
		for t := range graph {
			if t.value == firstLetter {
				return true
			}
		}
	}

	for t := range graph {
		if t.value == firstLetter {
			usedTiles := []tile{t}
			dfsTile(t, graph, usedTiles, word)
			if found {
				return found
			}
		}
	}
	return found
}

func dfsTile(t tile, g map[tile][]tile, usedTiles []tile, targetWord string) {
	if found {
		return
	}

	for _, adjacentTile := range g[t] {
		if found {
			return
		}
		usedTilesCopy := usedTiles
		alreadyUsed := false
		for _, usedTile := range usedTilesCopy {
			if usedTile == adjacentTile {
				alreadyUsed = true
				break
			}
		}

		if alreadyUsed {
			continue
		}

		usedTilesCopy = append(usedTilesCopy, adjacentTile)
		foundLetters := ""
		for _, usedTile := range usedTilesCopy {
			foundLetters += usedTile.value
		}

		if foundLetters == targetWord {
			found = true
			return
		}

		if len(usedTilesCopy) >= len(targetWord) {
			continue
		}

		if strings.Index(targetWord, foundLetters) == 0 {
			dfsTile(adjacentTile, g, usedTilesCopy, targetWord)
		}
	}
}
