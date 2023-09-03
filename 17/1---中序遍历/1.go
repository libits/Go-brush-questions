package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	node := parseAndBuildTree([]byte(s))
	printTree(node)
}

type Node struct {
	v     byte
	left  *Node
	right *Node
}

func parseAndBuildTree(bytes []byte) *Node {
	if len(bytes) == 0 {
		return nil
	}
	node := &Node{
		v: bytes[0],
	}
	if len(bytes) == 1 {
		return node
	}
	cnt := 0
	for i := 2; i < len(bytes)-1; i++ {
		b := bytes[i]
		if b == ',' && cnt == 0 {
			node.left = parseAndBuildTree(bytes[2:i])
			node.right = parseAndBuildTree(bytes[i+1 : len(bytes)-1])
			return node

		}
		if b == '{' {
			cnt += 1
		} else if b == '}' {
			cnt -= 1
		}
	}
	node.left = parseAndBuildTree(bytes[2 : len(bytes)-1])
	return node
}
func printTree(node *Node) {
	if node == nil {
		return
	}
	printTree(node.left)
	fmt.Printf("%s", []byte{node.v})
	printTree(node.right)
}

/*a{b{d,e{g,h{,i}}},c{f}}
dbgehiafc
*/
