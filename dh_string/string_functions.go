package dh_string

import (
	"fmt"
	"strconv"
)

func GetCPPArrayFormatString(array_name string, index int, field_name string, value interface{}) string {
	ret := ""

	var value_str string
	switch v := value.(type) {
	case int:
		value_str = strconv.Itoa(v)
	case float32:
	case float64:
		value_str = fmt.Sprintf("%g", v)
	case string:
		value_str = ""
		value_str += "\""
		value_str += v
		value_str += "\""
	case bool:
		if v == true {
			value_str = "true"
		} else {
			value_str = "false"
		}
	default:
		value_str = ""
	}

	ret += array_name
	ret += "["
	ret += strconv.Itoa(index)
	ret += "]."
	ret += field_name
	ret += "="
	ret += value_str
	ret += ";"

	return ret
}

func GetFirstNonSpaceCharacter(str string) rune {
	var ret rune = 0

	for _, c := range str {
		if c != ' ' {
			ret = c
			break
		}
	}

	return ret
}

func GetFirstStringInBetween(str string, first_character rune, second_character rune) string {
	first_pos := -1
	second_pos := -1

	for index, c := range str {
		if first_pos < 0 {
			if c == first_character {
				first_pos = index
				continue
			}
		}
		if second_pos < 0 {
			if c == second_character {
				second_pos = index
				break
			}
		}
	}

	if first_pos < 0 || second_pos < 0 {
		return str
	}

	rs := []rune(str)
	substr := make([]rune, second_pos-first_pos-1)
	for i := first_pos + 1; i < second_pos; i++ {
		substr[i-(first_pos+1)] = rs[i]
	}

	return string(substr)
}
