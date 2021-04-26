package pkg

import "strconv"


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

func isBuzz(i int) bool {
	return i%5 == 0
}

func isFizz(i int) bool {
	return i%3 == 0
}

func GetNumber(number int) string {
	return result[number-1]
}

