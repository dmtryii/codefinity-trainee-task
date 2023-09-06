package util

import "strconv"

func FizzBuzzNum(n int) string {
	result := ""
	if n%3 == 0 {
		result += "Fizz"
	}
	if n%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		return strconv.Itoa(n)
	}
	return result
}
