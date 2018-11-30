package node

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	key := 0
	level := 0
	node := NewNode(key, level)
	if node == nil {
		t.Errorf("Node was not created!")
	}
	if node.Key != key {
		t.Errorf("Key value is incorrect, expected: %d, got: %d", key, node.Key)
	}
	if len(node.Forward) != level {
		t.Errorf("Count elements in forward is incorrect, expected: %d, got: %d", level, len(node.Forward))
	}
}

func TestNewNodeWithLevel(t *testing.T) {
	key := 13
	level := 13
	node := NewNode(key, level)
	if node == nil {
		t.Errorf("TestNewNodeWithLevel: node was not created!")
	}
	if node.Key != key {
		t.Errorf("TestNewNodeWithLevel: key value is incorrect, expected: %d, got: %d", key, node.Key)
	}
	if len(node.Forward) != level {
		t.Errorf("TestNewNodeWithLevel: count elements in forward is incorrect, expected: %d, got: %d", level, len(node.Forward))
	}
}
