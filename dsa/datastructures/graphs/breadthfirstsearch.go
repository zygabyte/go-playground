// idea is to put a node into a queue, starting from the root node,
// then pop it, then add it's children nodes to the end of the queue and repeat the steps for the child nodes

package main

import "fmt"

type Node struct {
	Value    string
	Children []*Node // children of the graph
	Visited  bool    // to determine if this vertex has been visited
}

func main() {
	graph := Node{
		Value: "root",
		Children: []*Node{
			{Value: "1st"},
			{Value: "2nd", Children: []*Node{
				{Value: "2nd-1"},
				{Value: "2nd-2", Children: []*Node{
					{Value: "2nd2-1"},
					{Value: "2nd2-2"},
				}},
			}},
			{Value: "3rd"},
		},
	}

	fmt.Println(graph.search())
}

// Node pointer receiver to ensure this can only be called by a node
func (node *Node) search() []string {
	// add root element
	node.Visited = true
	queue := []*Node{node}
	array := []string{} // to store all the values of the graphs uring BFS

	array = append(array, node.Value)

	// as long as queue isn't empty keep looping and appending to queue (more children available)
	for len(queue) > 0 {
		current := queue[0] // get first element from queue
		queue = queue[1:]   // assign the remaining items in queue (excluding first item) back to queue (popping)

		// if not visited, add the value of graph node
		if !current.Visited {
			current.Visited = true
			array = append(array, current.Value)
		}

		// if there are children, add them to the queue
		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}

	return array
}
