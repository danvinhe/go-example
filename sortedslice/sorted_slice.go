package sortedslice

import "fmt"

// SortedSlice 有序切片
type SortedSlice struct {
	data []interface{}
	less func(a, b interface{}) bool
}

// New 初始化有序切片
func New(less func(a, b interface{}) bool) *SortedSlice {
	return &SortedSlice{data: make([]interface{}, 0), less: less}
}

// Len 元素个数
func (s *SortedSlice) Len() int {
	return len(s.data)
}

// Add 添加元素
func (s *SortedSlice) Add(v interface{}) int {
	size := s.Len()
	index := s.bisectLeft(v)
	if index == 0 {
		newData := make([]interface{}, 0)
		newData = append(newData, v)
		newData = append(newData, s.data...)
		s.data = newData
	} else if index == size {
		s.data = append(s.data, v)
	} else {
		newData := s.data[:index]
		newData = append(newData, v)
		newData = append(newData, s.data[index+1:]...)
		s.data = newData
	}
	return index
}

// Remove 移除元素
func (s *SortedSlice) Remove(index int) bool {
	if index < 0 || index >= s.Len() {
		return false
	}

	if index == 0 {
		s.data = s.data[1:]
	} else if index == s.Len()-1 {
		s.data = s.data[:index]
	} else {
		s.data = append(s.data[:index], s.data[index+1:]...)
	}

	return true
}

// Clear 清空
func (s *SortedSlice) Clear() {
	s.data = make([]interface{}, 0)
}

// At 获取索引上的值
func (s *SortedSlice) At(i int) interface{} {
	return s.data[i]
}

func (s *SortedSlice) String() string {
	return fmt.Sprintf("%v", s.data)
}

func (s *SortedSlice) bisectLeft(x interface{}) int {
	left, right := 0, s.Len()
	for left < right {
		middle := (left + right) / 2
		if s.less(s.data[middle], x) {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}
