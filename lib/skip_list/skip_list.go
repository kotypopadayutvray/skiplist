package skip_list

import (
	"../node"
	"math/rand"
)

type SkipList struct {
	MaxLevel int
	Header *node.Node
	Probability float64
	CurrentLevel int
}

func NewSkipList(maxLevel int, prob float64) *SkipList {
	// Create the header node
	// Header will contain pointers to inserted nodes
	header := node.NewNode(0, maxLevel)
	return &SkipList {
		MaxLevel: maxLevel,
		Header: header,
		Probability: prob,
		CurrentLevel: 0,
	}
}

func (sl *SkipList) GenerateLevel() int {
	level := 1
	for rand.Float64() < sl.Probability && level <= sl.MaxLevel {
		level++
	}
	if level > sl.MaxLevel {
		level = sl.MaxLevel
	}
	return level
}

func (sl *SkipList) Insert(key int) {
	// Create an update slice
	// Elements of the update will contain pointers to one element from
	// any of all levels. It necessary for inserting element in the future
	update := [] *node.Node { }
	// Fill update array by nul values
	for i := 0; i < sl.MaxLevel; i++ {
		update = append(update, nil)
	}
	// Start search place for inserting from header node
	current_node := sl.Header
	// Go from the highest level
	for level := sl.MaxLevel - 1; level >= 0; level-- {
		// If current node has next element in the same level
		// and this next element has key less than inserted key
		for current_node.Forward[level] != nil && current_node.Forward[level].Key < key {
			// Move through the current level next
			current_node = current_node.Forward[level]
		}
		// Insert found node in the update array
		update[level] = current_node
	}
	// Get the new level of the new element
	new_element_level := sl.GenerateLevel()
	// Create new element
	new_element := node.NewNode(key, new_element_level)
	// Insert new element in the list
	for i := 0; i < new_element_level; i++ {
		new_element.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = new_element
	}
}
