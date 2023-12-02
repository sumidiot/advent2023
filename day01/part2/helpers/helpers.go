package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func startDigit (str string) int {
	res, _ := regexp.MatchString("[0-9]", str[0:1])
	if res {
		res, _ := strconv.Atoi(str[0:1])
		return res
	} else {
		switch {
			case strings.HasPrefix(str, "zero"):   return 0
			case strings.HasPrefix(str, "one"):    return 1
			case strings.HasPrefix(str, "two"):    return 2
			case strings.HasPrefix(str, "three"):  return 3
			case strings.HasPrefix(str, "four"):   return 4
			case strings.HasPrefix(str, "five"):   return 5
			case strings.HasPrefix(str, "six"):    return 6
			case strings.HasPrefix(str, "seven"):  return 7
			case strings.HasPrefix(str, "eight"):  return 8
			case strings.HasPrefix(str, "nine"):   return 9
			default:                               return -1
		}
	}
}

func Calval(str string) int {
	first := -1
	latest := -1
	for chridx := 0; chridx < len(str); chridx++ {
		digit := startDigit(str[chridx:])
		if digit >= 0 {
			latest = digit
			if first == -1 {
				first = digit
			}
		}
	}
	calval, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, latest))
	return calval
}
