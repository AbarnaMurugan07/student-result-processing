package main

type Student struct {
	Name       string
	RegisterNo string
	Marks      []int
}

func calculateTotal(marks []int) int {
	total := 0

	for _, mark := range marks {
		total += mark
	}

	return total
}

func calculateAverage(total int, count int) float64 {
	return float64(total) / float64(count)
}

func calculateGrade(average float64) string {
	if average >= 90 {
		return "A"
	} else if average >= 80 {
		return "B"
	} else if average >= 70 {
		return "C"
	} else if average >= 50 {
		return "D"
	}

	return "F"
}
