package algorithm

import "math"

func salarySimilarity(candidateSalaries []float64, jobSalary float64) float64 {
	if len(candidateSalaries) == 0 || jobSalary == 0 {
		return 0
	}

	sum := 0.0

	for _, s := range candidateSalaries {
		sum += s
	}

	avg := sum / float64(len(candidateSalaries))

	diff := math.Abs(avg - jobSalary)
	similarity := 1 - (diff / jobSalary)

	if similarity < 0 {
		similarity = 0
	}

	return similarity
}
