package pkg

import "strconv"

func FizzBuzz() (result []string) {

	for i := 1 ; i <= 100; i++ {
		if i % 3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
