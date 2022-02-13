package slice

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}

	return sums
}
