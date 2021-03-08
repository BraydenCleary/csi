package main

func main() {
}

const (
	unvisited = iota
	visiting  = iota
	visited   = iota
)

func buildGraph(prerequisites [][]int) map[int][]int {
	output := make(map[int][]int)
	for _, pair := range prerequisites {
		output[pair[0]] = append(output[pair[0]], pair[1])
	}
	return output
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(prerequisites)
	nodeStates := make(map[int]int)

	for i := 0; i < numCourses; i++ {
		if nodeStates[i] == unvisited {
			if containsCycle(i, graph, nodeStates) {
				return false
			}
		}
	}

	return true
}

func containsCycle(idx int, graph map[int][]int, nodeStates map[int]int) bool {
	nodeStates[idx] = visiting

	for _, prereq := range graph[idx] {
		if nodeStates[prereq] == visiting {
			return true
		}

		if nodeStates[prereq] == unvisited {
			nodeStates[prereq] = visiting
			if containsCycle(prereq, graph, nodeStates) {
				return true
			}
		}
	}
	nodeStates[idx] = visited
	return false
}

// maybe depth first search and check if you ever see the original key
// We're basically trying to figure out if the graph is acyclic...aka if there are no
// courses that require themselves as prereqs
