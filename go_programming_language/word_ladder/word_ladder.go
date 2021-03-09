package main

func main() {

}

func buildGraph(wordList []string) map[string][]string {
	output := make(map[string][]string)
	for _, word := range wordList {
		for _, otherWord := range wordList {
			if word == otherWord {
				continue
			}

			skip := false
			for _, w := range output[word] {
				if w == otherWord {
					skip = true
				}
			}

			if skip {
				continue
			}

			characterCountDifference := 0
			shouldAdd := true
			w := []rune(word)
			ow := []rune(otherWord)
			for idx, r := range w {
				if ow[idx] != r {
					characterCountDifference++
				}

				if characterCountDifference > 1 {
					shouldAdd = false
					break
				}
			}
			if shouldAdd {
				output[word] = append(output[word], otherWord)
			}
		}
	}
	return output
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordListContainsEndWord := false
	for _, word := range wordList {
		if word == endWord {
			wordListContainsEndWord = true
		}
	}

	if !wordListContainsEndWord {
		return 0
	}

	fullWordList := append(wordList, beginWord)
	fullWordList = append(fullWordList, endWord)
	graph := buildGraph(fullWordList)

	visitedWords := make(map[string]bool)
	initialWord := wordWithDepth{word: beginWord, depth: 0}

	workQueue, visitedWords, depth := bfsForWord(initialWord, endWord, visitedWords, graph, []wordWithDepth{})
	for len(workQueue) > 0 && depth < 0 {
		nextWord := workQueue[0]
		workQueue = workQueue[1:]
		workQueue, visitedWords, depth = bfsForWord(nextWord, endWord, visitedWords, graph, workQueue)
	}

	if depth < 0 {
		return 0
	}

	return depth + 1
}

type wordWithDepth struct {
	word  string
	depth int
}

func bfsForWord(w wordWithDepth, targetWord string, visitedWords map[string]bool, graph map[string][]string, workQueue []wordWithDepth) ([]wordWithDepth, map[string]bool, int) {
	depth := -1
	if _, ok := visitedWords[w.word]; !ok {
		visitedWords[w.word] = true
	}
	for _, w1 := range graph[w.word] {
		if w1 == targetWord {
			depth = w.depth + 1
			break
		}

		if _, ok := visitedWords[w1]; !ok {
			visitedWords[w1] = true
			wordWithDepthToAdd := wordWithDepth{word: w1, depth: w.depth + 1}
			workQueue = append(workQueue, wordWithDepthToAdd)
		}
	}

	return workQueue, visitedWords, depth
}
