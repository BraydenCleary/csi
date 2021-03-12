package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {}

const startingPosition = "0000"

func neighborNumber(r rune, direction int32) rune {
	if r == '0' {
		if direction == -1 {
			return '9'
		} else {
			return '1'
		}
	}

	if r == '9' {
		if direction == 1 {
			return '0'
		} else {
			return '8'
		}
	}

	return r + direction
}

func findLockNeighbors(val string, deadends []string) []string {
	valAsRunes := []rune(val)
	output := []string{}
	for idx, r := range valAsRunes {
		for i := -1; i < 2; i += 2 {
			v := []rune(string(valAsRunes))
			v[idx] = neighborNumber(r, int32(i))
			shouldAdd := true
			for _, d := range deadends {
				if d == string(v) {
					shouldAdd = false
				}
			}

			if shouldAdd {
				output = append(output, string(v))
			}
		}
	}
	return output
}

func buildGraph(deadends []string) map[string][]string {
	output := make(map[string][]string)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for m := 0; m < 10; m++ {
					lockValue := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(m)
					shouldAdd := true
					for _, d := range deadends {
						if d == lockValue {
							shouldAdd = false
						}
					}

					if !shouldAdd {
						continue
					}

					output[lockValue] = findLockNeighbors(lockValue, deadends)
				}
			}
		}
	}
	return output
}

func openLock(deadends []string, target string) int {
	graph := buildGraph(deadends)

	if len(graph[target]) == 0 {
		return -1
	}

	if len(graph[startingPosition]) == 0 {
		return -1
	}

	return bfsForLock(startingPosition, target, graph)
}

func bfsForLock(startingPosition string, target string, graph map[string][]string) int {
	queue := [][]string{
		{startingPosition},
	}
	i := 0
	for len(queue) > 0 {
		fmt.Println(len(queue))
		currentPath := queue[0]
		fmt.Println(currentPath)
		queue = queue[1:]

		lastInCurrentPath := currentPath[len(currentPath)-1]

		for _, position := range graph[lastInCurrentPath] {
			shouldAddToPath := true
			for _, p := range currentPath {
				if position == p {
					shouldAddToPath = false
				}
			}
			if position == target {
				return len(currentPath)
			}

			if shouldAddToPath {
				shouldAddToQueue := true
				for _, q := range queue {
					if reflect.DeepEqual(q, append(currentPath, position)) {
						shouldAddToQueue = false
					}
				}

				if shouldAddToQueue {
					pathToAdd := append(currentPath, position)
					queue = append(queue, pathToAdd)
				}
			}
		}
		i++
	}

	return -1
}
