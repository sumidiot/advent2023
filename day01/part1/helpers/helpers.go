package helpers

import (
	"fmt"
	"regexp"
	"strconv"
)

func Calval(line string) int {
	digits := regexp.MustCompile("^[^0-9]*([0-9]).*?([0-9]?)[^0-9]*$")
	first := digits.ReplaceAllString(line, "$1")
	second := digits.ReplaceAllString(line, "$2")
	last := first
	if len(second) > 0 {
		last = second
	}
	ret, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
	return ret
}
