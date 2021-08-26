package soundex

import (
	"fmt"
	"regexp"
	"strings"
)

var alpha2digit = map[byte]string{
	'a': "0", 'e': "0", 'i': "0", 'o': "0", 'u': "0",
	'y': "9", 'h': "9", 'w': "9",
	'b': "1", 'f': "1", 'p': "1", 'v': "1",
	'c': "2", 'g': "2", 'j': "2", 'k': "2", 'q': "2", 's': "2", 'x': "2", 'z': "2",
	'd': "3", 't': "3",
	'l': "4",
	'm': "5", 'n': "5",
	'r': "6",
}

// Value 计算给定字符串的soundex
func Value(word string) (result string, err error) {
	// 判断字符串合法性，只支持英文单词
	matched, err := regexp.MatchString(`^[A-Za-z]+$`, word)
	if err != nil {
		return
	}
	if !matched {
		err = fmt.Errorf("must match ^[A-Za-z]+$")
		return
	}

	// 把单词转换为小写，并替换为对应数字
	lowerWord := strings.ToLower(word)
	digits := alpha2digit[lowerWord[0]]
	for i := 1; i < len(lowerWord); i++ {
		digits += alpha2digit[lowerWord[i]]
	}

	// 删除重复数字
	digits2 := strings.ReplaceAll(digits, "9", "")
	digits3 := ""
	for i := 0; i < len(digits2); i++ {
		digit := digits2[i]
		l := len(digits3)
		if i == 0 {
			digits3 = string(digit)
		} else {
			if digits3[l-1] == digit {
				continue
			}
			digits3 += string(digit)
		}
	}

	// 删除元音字母
	digits4 := strings.ReplaceAll(digits3, "0", "")

	// 首字母可能被删除了，需要补上
	digits5 := digits4
	if alpha2digit[lowerWord[0]] == "0" || alpha2digit[lowerWord[0]] == "9" {
		digits5 = "x" + digits4
	}

	// 为防止长度不够，统一填充3个0
	digits6 := digits5 + strings.Repeat("0", 3)

	// firstChar :=
	// if 'A' <= firstChar && firstChar <= 'Z' {
	// 	firstChar += 20
	// }

	// 保留首字母，并转换为大写
	result = strings.ToUpper(string(word[0])) + digits6[1:4]

	return
}
