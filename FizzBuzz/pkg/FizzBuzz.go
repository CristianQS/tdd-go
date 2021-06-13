package pkg

import (
	"strconv"
	"strings"
)


var (
	result []string
)

func Transform() [] string {
	for number := 1 ; number <= 100; number++ {
		if isFizzBuzz(number) {
			result = append(result, "FizzBuzz")
		} else if isFizz(number) {
			result = append(result, "Fizz")
		} else if isBuzz(number) {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(number))
		}
	}
	return result
}

func isFizzBuzz(i int) bool {
	return isFizz(i) && isBuzz(i)
}

func isBuzz(number int) bool {
	return number % 5 == 0 || strings.Contains(strconv.Itoa(number), "5")
}

func isFizz(number int) bool {
	return number % 3 == 0 || strings.Contains(strconv.Itoa(number), "3")
}

func GetNumber(number int) string {
	return result[number-1]
}

