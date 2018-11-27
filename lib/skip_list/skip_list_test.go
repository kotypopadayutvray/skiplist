package skip_list

import (
	"sort"
	"testing"
)

const max_level = 5
const prob = 1.0

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList(max_level, prob)
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
	t.Run("sl.MaxLevel = max_level", func(t *testing.T) {
		if sl.MaxLevel != max_level {
			t.Errorf("MaxLevel is incorrect, expected: %d, got: %d", max_level, sl.MaxLevel)
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
		values_to_insert := []int { -4, 3, 2, 10, 51, -78, 6, 0 }
		sl := NewSkipList(max_level, prob)
		for _, value := range values_to_insert {
			sl.Insert(value)
		}
		count_inserted_elements := 0
		current_element := sl.Header.Forward[0]
		for current_element != nil {
			current_element = current_element.Forward[0]
			count_inserted_elements++
		}
		count_values_to_insert := len(values_to_insert)
		if count_inserted_elements != count_values_to_insert {
			t.Errorf("Wrong count of inserted elements. Expected: %d, got: %d", count_values_to_insert, count_inserted_elements)
		}
	})
	t.Run("should insert in correct sequence", func(t *testing.T) {
		values_to_insert := []int { -4, 3, 2, 10, 51, -78, 6, 0 }
		sorted_values := make([]int, len(values_to_insert))
		copy(sorted_values, values_to_insert)
		sort.Ints(sorted_values)
		sl := NewSkipList(max_level, prob)
		for _, value := range values_to_insert {
			sl.Insert(value)
		}
		index := 0
		current_element := sl.Header.Forward[0]
		for current_element != nil {
			if current_element.Key != sorted_values[index] {
				t.Errorf("Got wrong value. Expected: %d, got: %d", sorted_values[index], current_element.Key)
			}
			current_element = current_element.Forward[0]
			index++
		}
	})
}
