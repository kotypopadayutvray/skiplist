package skiplist

import (
	"math/rand"

	"../node"
)

// SkipList structure
type SkipList struct {
	MaxLevel    int
	Header      *node.Node
	Probability float64
}

// NewSkipList function create a new SkipList instance
func NewSkipList(maxLevel int, prob float64) *SkipList {
	// Create the header node
	// Header will contain pointers to inserted nodes
	header := node.NewNode(0, maxLevel)
	return &SkipList{
		MaxLevel:    maxLevel,
		Header:      header,
		Probability: prob,
	}
}

// Generate level of the new inserted element
func (sl *SkipList) generateLevel() int {
	level := 1
	for rand.Float64() < sl.Probability && level <= sl.MaxLevel {
		level++
	}
	if level > sl.MaxLevel {
		level = sl.MaxLevel
	}
	return level
}

func (sl *SkipList) path(key int) []*node.Node {
	// Create an update slice
	// Elements of the update will contain pointers to one element from
	// any of all levels. It necessary for inserting element in the future
	update := []*node.Node{}
	// Fill update array by nul values
	for i := 0; i < sl.MaxLevel; i++ {
		update = append(update, nil)
	}
	// Start search place for inserting from header node
	currentNode := sl.Header
	// Go from the highest level
	for level := sl.MaxLevel - 1; level >= 0; level-- {
		// If current node has next element in the same level
		// and this next element has key less than inserted key
		for currentNode.Forward[level] != nil && currentNode.Forward[level].Key < key {
			// Move through the current level next
			currentNode = currentNode.Forward[level]
		}
		// Insert found node in the update array
		update[level] = currentNode
	}
	return update
}

// Search element by key in the SkipList instance
func (sl *SkipList) Search(key int) bool {
	path := sl.path(key)
	lowLevel := 0
	return path[lowLevel].Forward[lowLevel].Key == key
}

// Insert the new element with key in the SkipList instance
func (sl *SkipList) Insert(key int) bool {
	// Create the path array
	path := sl.path(key)
	// Check if we have the same key
	// For this purpose get the low level element
	lowLevelElement := path[0].Forward[0]
	// And compare its key with inserted key
	if lowLevelElement != nil && lowLevelElement.Key == key {
		// If keys are equal don't insert
		return false
	}
	// Get the new level of the new element
	newElementLevel := sl.generateLevel()
	// Create new element
	newElement := node.NewNode(key, newElementLevel)
	// Insert new element in the list
	for i := 0; i < newElementLevel; i++ {
		// Update pointer for every level
		newElement.Forward[i] = path[i].Forward[i]
		path[i].Forward[i] = newElement
	}
	return true
}

// Delete element in the SkipList instance
func (sl *SkipList) Delete(key int) bool {
	path := sl.path(key)
	lowLevel := 0
	deletingElement := path[lowLevel].Forward[lowLevel]
	if deletingElement != nil && deletingElement.Key == key {
		for level := 0; level < len(path); level++ {
			if path[level].Forward[level] != deletingElement {
				continue
			}
			path[level].Forward[level] = deletingElement.Forward[level]
		}
		return true
	}
	return false
}
