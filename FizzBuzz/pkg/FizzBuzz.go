package pkg

import "strconv"


var (
	result []string
)

func Transform() [] string {
	for i := 1 ; i <= 100; i++ {
		if isFizz(i) {
			result = append(result, "Fizz")
		} else if isBuzz(i) {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func isBuzz(i int) bool {
	return i%5 == 0
}

func isFizz(i int) bool {
	return i%3 == 0
}

func GetNumber(number int) string {
	return result[number-1]
}

