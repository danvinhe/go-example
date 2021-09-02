package fibonacci

// Sequence 获取最大值小于等于max的数列
func Sequence(max int) []int {
	if max < 0 {
		panic("max must be >= 0")
	}

	arr := make([]int, 0)
	if max == 0 {
		arr = append(arr, 0)
		return arr
	}

	first, second := 0, 1
	arr = append(arr, first)
	for {
		first, second = second, first+second
		if first > max {
			break
		}
		arr = append(arr, first)
	}

	return arr
}

// At 获取数列的第n个数值
func At(index int) int {
	if index < 0 {
		panic("index must be >= 0")
	}

	if index < 2 {
		return index
	}

	return At(index-1) + At(index-2)
}
