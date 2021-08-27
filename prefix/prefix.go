package prefix

import "fmt"

// Common 共同前缀
func Common(list []string) (prefix string) {
	if len(list) == 0 {
		return
	}

	// 按utf8格式分离每个字符串的字符
	min := 0
	table := make([][]rune, len(list))
	for i, str := range list {
		if len(str) == 0 {
			return
		}

		row := make([]rune, 0)
		for _, char := range str {
			row = append(row, char)
		}
		table[i] = row

		if min == 0 || min > len(row) {
			min = len(row)
		}
	}

	// 逐一比较每列字符是否一致
	for j := 0; j < min; j++ {
		char := rune(0)
		for i := 0; i < len(table); i++ {
			if i == 0 {
				char = table[i][j]
			} else {
				if char != table[i][j] {
					char = rune(0)
					break
				}
			}
		}

		if char != 0 {
			prefix += fmt.Sprintf("%c", char)
		} else {
			return
		}
	}

	return
}
