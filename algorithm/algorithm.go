package algorithm

type Candidate struct {
	Skills          []string
	PretendSalaries []float64
}

type Job struct {
	Skills []string
	Salary float64
}

func AlgorithmRecommendation(candidate Candidate, job Job, weighSkills, weighSalary float64) float64 {

	skillsSimilarity := cosineSimilarity(candidate.Skills, job.Skills)
	salarySimilarity := salarySimilarity(candidate.PretendSalaries, job.Salary)

	score := (weighSkills * skillsSimilarity) + (weighSalary * salarySimilarity)

	if score > 1 {
		score = 1
	}

	if score < 0 {
		score = 0
	}

	return score
}
