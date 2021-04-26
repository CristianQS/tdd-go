package pkg

import "strconv"


var (
	result []string
)

func Transform() [] string {
	for i := 1 ; i <= 100; i++ {
		if i % 3 == 0 {
			result = append(result, "Fizz")
		} else if i % 5 == 0 {
			result = append(result, "Buzz")
		} else{
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func GetNumber(number int) string {
	return result[number-1]
}

