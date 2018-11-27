package skip_list

import (
	"../node"
	"math/rand"
)

type SkipList struct {
	MaxLevel int
	Header *node.Node
	Probability float64
}

func NewSkipList(maxLevel int, prob float64) *SkipList {
	// Create the header node
	// Header will contain pointers to inserted nodes
	header := node.NewNode(0, maxLevel)
	return &SkipList {
		MaxLevel: maxLevel,
		Header: header,
		Probability: prob,
	}
}

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

func (sl *SkipList) path(key int) [] *node.Node {
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
	return update
}

func (sl *SkipList) Search(key int) bool {
	path := sl.path(key)
	low_level := 0
	return path[low_level].Forward[low_level].Key == key
}

func (sl *SkipList) Insert(key int) {
	// Create the path array
	path := sl.path(key)
	// Check if we have the same key
	// For this purpose get the low level element
	low_level_element := path[0].Forward[0]
	// And compare its key with inserted key
	if low_level_element != nil && low_level_element.Key == key {
		// If keys are equal don't insert
		return
	}
	// Get the new level of the new element
	new_element_level := sl.generateLevel()
	// Create new element
	new_element := node.NewNode(key, new_element_level)
	// Insert new element in the list
	for i := 0; i < new_element_level; i++ {
		// Update pointer for every level
		new_element.Forward[i] = path[i].Forward[i]
		path[i].Forward[i] = new_element
	}
}

func (sl *SkipList) Delete(key int) {
	path := sl.path(key)
	low_level := 0
	deleting_element := path[low_level].Forward[low_level]
	if deleting_element.Key == key {
		for level := 0; level < len(path); level++ {
			if path[level].Forward[level] != deleting_element {
				continue
			}
			path[level].Forward[level] = deleting_element.Forward[level]
		}
	}
}
