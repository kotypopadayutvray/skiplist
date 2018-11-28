package skiplist

import (
	"sort"
	"testing"
)

const maxLevel = 5
const prob = 1.0

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList(maxLevel, prob)
	t.Run("sl is not nil", func(t *testing.T) {
		if sl == nil {
			t.Errorf("New skip list was not created!")
		}
	})
	t.Run("sl.Header is not nil", func(t *testing.T) {
		if sl.Header == nil {
			t.Errorf("sl.Header was not created!")
		}
	})
	t.Run("sl.MaxLevel = maxLevel", func(t *testing.T) {
		if sl.MaxLevel != maxLevel {
			t.Errorf("MaxLevel is incorrect, expected: %d, got: %d", maxLevel, sl.MaxLevel)
		}
	})
	t.Run("sl.Probability = prob", func(t *testing.T) {
		if sl.Probability != prob {
			t.Errorf("TestNewSkipList: Probability value is incorrect, expected: %f, got: %f!", prob, sl.Probability)
		}
	})

}

func TestInsert(t *testing.T) {
	t.Run("should insert all values", func(t *testing.T) {
		valuesToInsert := []int{-4, 3, 2, 10, 51, -78, 6, 0}
		sl := NewSkipList(maxLevel, prob)
		for _, value := range valuesToInsert {
			sl.Insert(value)
		}
		countInsertedElements := 0
		currentElement := sl.Header.Forward[0]
		for currentElement != nil {
			currentElement = currentElement.Forward[0]
			countInsertedElements++
		}
		countValuesToInsert := len(valuesToInsert)
		if countInsertedElements != countValuesToInsert {
			t.Errorf("Wrong count of inserted elements. Expected: %d, got: %d", countValuesToInsert, countInsertedElements)
		}
	})
	t.Run("should insert in correct sequence", func(t *testing.T) {
		valuesToInsert := []int{-4, 3, 2, 10, 51, -78, 6, 0}
		sortedValues := make([]int, len(valuesToInsert))
		copy(sortedValues, valuesToInsert)
		sort.Ints(sortedValues)
		sl := NewSkipList(maxLevel, prob)
		for _, value := range valuesToInsert {
			sl.Insert(value)
		}
		index := 0
		currentElement := sl.Header.Forward[0]
		for currentElement != nil {
			if currentElement.Key != sortedValues[index] {
				t.Errorf("Got wrong value. Expected: %d, got: %d", sortedValues[index], currentElement.Key)
			}
			currentElement = currentElement.Forward[0]
			index++
		}
	})
	t.Run("should not insert existing element", func(t *testing.T) {
		sl := NewSkipList(maxLevel, prob)
		sl.Insert(3)
		// Try to insert existing element
		if sl.Insert(3) {
			t.Errorf("Inserted existing element")
		}
		countElements := 0
		expectedCountElements := 1
		el := sl.Header.Forward[0]
		for el != nil {
			el = el.Forward[0]
			countElements++
		}
		if countElements != expectedCountElements {
			t.Errorf("Count of elements wrong. Expected: %d, got: %d", expectedCountElements, countElements)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete element in the empty list", func(t *testing.T) {
		sl := NewSkipList(maxLevel, prob)
		if sl.Delete(1) {
			t.Errorf("Successfully delete not existing element")
		}
	})
	t.Run("Delete one element in the begin of list", func(t *testing.T) {
		sl := NewSkipList(maxLevel, prob)
		elementsToInsert := [5]int{1, 2, 3, 4, 5}
		for _, key := range elementsToInsert {
			sl.Insert(key)
		}
		// Check deleting
		if !sl.Delete(1) {
			t.Errorf("Don't delete the first element")
		}
		// Check elements
		index := 1
		el := sl.Header.Forward[0]
		for el != nil {
			expectedElement := elementsToInsert[index]
			if el.Key != expectedElement {
				t.Errorf("The sequence after delete element is incorrect, expected: %d, got: %d", expectedElement, el.Key)
			}
			el = el.Forward[0]
			index++
		}
	})
	t.Run("Delete one element in the end of list", func(t *testing.T) {
		sl := NewSkipList(maxLevel, prob)
		elementsToInsert := [5]int{1, 2, 3, 4, 5}
		for _, key := range elementsToInsert {
			sl.Insert(key)
		}
		// Check deleting
		if !sl.Delete(5) {
			t.Errorf("Don't delete the first element")
		}
		// Check elements
		index := 0
		el := sl.Header.Forward[0]
		for el != nil {
			expectedElement := elementsToInsert[index]
			if el.Key != expectedElement {
				t.Errorf("The sequence after delete element is incorrect, expected: %d, got: %d", expectedElement, el.Key)
			}
			el = el.Forward[0]
			index++
		}
	})
	t.Run("Delete one element in the middle of list", func(t *testing.T) {
		sl := NewSkipList(maxLevel, prob)
		elementsToInsert := [5]int{1, 2, 3, 4, 5}
		for _, key := range elementsToInsert {
			sl.Insert(key)
		}
		// Check deleting
		if !sl.Delete(3) {
			t.Errorf("Don't delete the first element")
		}
		// Check elements
		elementForCheck := [4]int{1, 2, 4, 5}
		index := 0
		el := sl.Header.Forward[0]
		for el != nil {
			expectedElement := elementForCheck[index]
			if el.Key != expectedElement {
				t.Errorf("The sequence after delete element is incorrect, expected: %d, got: %d", expectedElement, el.Key)
			}
			el = el.Forward[0]
			index++
		}
	})
	t.Run("Delete elements in the list", func(t *testing.T) {
		sl := NewSkipList(maxLevel, prob)
		elementsToInsert := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, key := range elementsToInsert {
			sl.Insert(key)
		}
		elementsToDelete := [4]int{1, 3, 5, 9}
		for _, key := range elementsToDelete {
			if !sl.Delete(key) {
				t.Errorf("Can't delete existing element %d", key)
			}
		}
		// Check elements
		elementForCheck := [6]int{2, 4, 6, 7, 8, 10}
		index := 0
		el := sl.Header.Forward[0]
		for el != nil {
			expectedElement := elementForCheck[index]
			if el.Key != expectedElement {
				t.Errorf("The sequence after delete element is incorrect, expected: %d, got: %d", expectedElement, el.Key)
			}
			el = el.Forward[0]
			index++
		}
	})
}
