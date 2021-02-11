package main

func Sum(v []int) int {
	sum := 0

	for _, n := range v {
		sum += n
	}
	return sum
}

func SumAllTrails(numbers ...[]int) []int {

	var sums []int
	for _, n := range numbers {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(n[1:]))
		}
	}
	return sums
}
