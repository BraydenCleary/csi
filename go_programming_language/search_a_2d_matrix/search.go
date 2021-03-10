package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	numRows := len(matrix)
	numCols := len(matrix[0])

	if numRows == 1 && numCols == 1 {
		return matrix[0][0] == target
	}

	if target < matrix[0][0] {
		return false
	}

	doRowSearch := true
	if numRows == 1 {
		doRowSearch = false
	}

	row := 0
	if doRowSearch {
		// binary search to find row
		firstColumn := []int{}
		for i := 0; i < numRows; i++ {
			firstColumn = append(firstColumn, matrix[i][0])
		}

		left := 0
		right := numRows - 1
		middle := (left + right) / 2
		for {
			val := firstColumn[middle]
			if val == target {
				return true
			}

			if val < target && middle+1 <= numRows-1 {
				if firstColumn[middle+1] > target {
					row = middle
					break
				}

				if firstColumn[middle+1] == target {
					return true
				}
			}

			if val > target && middle-1 >= 0 {
				if firstColumn[middle-1] < target {
					row = middle - 1
					break
				}

				if firstColumn[middle-1] == target {
					return true
				}
			}

			if val > target {
				right = middle
				if right-left == 1 {
					row = left
					break
				}
			} else {
				left = middle
				if right-left == 1 {
					row = right
					break
				}
			}
			middle = (left + right) / 2
		}
	}

	if matrix[row][0] > target {
		return false
	}

	// then binary search to find tile
	left := 0
	right := numCols - 1
	middle := (left + right) / 2

	if left == right {
		return matrix[row][left] == target
	}

	for {
		val := matrix[row][middle]

		if val == target {
			return true
		}

		if val > target {
			right = middle
			if right-left == 1 {
				return matrix[row][left] == target
			}
		} else {
			left = middle
			if right-left == 1 {
				return matrix[row][right] == target
			}
		}

		middle = (left + right) / 2
	}

	return false
}
