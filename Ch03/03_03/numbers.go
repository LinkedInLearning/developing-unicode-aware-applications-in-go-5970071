package numbers

import (
	"regexp"
	"strconv"
)

var numRe = regexp.MustCompile(`\pN+`)

// FindNums returns slice of number found in text.
func FindNums(text string) ([]int, error) {
	var nums []int
	for _, s := range numRe.FindAllString(text, -1) {
		n, err := strconv.Atoi(s)
		if err != nil {
			n, err = arAtoi(s)
			if err != nil {
				return nil, err
			}
		}
		nums = append(nums, n)
	}

	return nums, nil
}
