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

	visitedWords := []string{}
	initialWord := wordWithDepth{word: beginWord, depth: 0}

	workQueue, visitedWords, depth := bfsForWord(initialWord, endWord, visitedWords, graph, []wordWithDepth{})
	for len(workQueue) > 0 && depth < 0 && len(visitedWords) < len(wordList)-1 {
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

func bfsForWord(w wordWithDepth, targetWord string, visitedWords []string, graph map[string][]string, workQueue []wordWithDepth) ([]wordWithDepth, []string, int) {
	depth := -1
	for _, w1 := range graph[w.word] {
		if w1 == targetWord {
			depth = w.depth + 1
			break
		}

		shouldAdd := true
		for _, w2 := range workQueue {
			if w2.word == w1 {
				shouldAdd = false
			}
		}
		if shouldAdd {
			wordWithDepthToAdd := wordWithDepth{word: w1, depth: w.depth + 1}
			workQueue = append(workQueue, wordWithDepthToAdd)
		}
	}

	shouldAddToVisited := true
	for _, w3 := range visitedWords {
		if w.word == w3 {
			shouldAddToVisited = false
		}
	}

	if shouldAddToVisited {
		visitedWords = append(visitedWords, w.word)
	}

	return workQueue, visitedWords, depth
}
