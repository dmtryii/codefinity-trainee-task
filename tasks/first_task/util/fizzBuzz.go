package util

func FizzBuzz(n int) []string {
	result := make([]string, n+1)
	for i := 1; i <= n; i++ {
		result[i] = FizzBuzzNum(i)
	}
	return result
}
