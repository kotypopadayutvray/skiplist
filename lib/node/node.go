package node

type Node struct {
	Key     int
	Forward []*Node // Slice
}

func NewNode(key, level int) *Node {
	// Create the empty slice
	forward := []*Node{}
	// Fill the slice by the nil values
	for i := 0; i < level; i++ {
		forward = append(forward, nil)
	}
	return &Node{
		Key:     key,
		Forward: forward,
	}
}
