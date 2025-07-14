package algorithm

import "math"

func cosineSimilarity(candidateSkills, jobSkills []string) float64 {
	candidateSkillsSet := make(map[string]struct{}, len(candidateSkills))

	for _, skill := range candidateSkills {
		candidateSkillsSet[skill] = struct{}{}
	}

	candidateVector := make([]float64, len(candidateSkillsSet))
	jobVector := make([]float64, len(candidateSkillsSet))

	for i, skill := range jobSkills {
		if _, ok := candidateSkillsSet[skill]; ok {
			candidateVector[i] = 1

		} else {
			candidateVector[i] = 0
		}
		jobVector[i] = 1
	}

	dotProduct := 0.0
	magnitudeCandidate := 0.0
	magnitudeJob := 0.0

	for i := 0; i < len(candidateVector); i++ {
		dotProduct += candidateVector[i] * jobVector[i]
		magnitudeCandidate += candidateVector[i] * candidateVector[i]
		magnitudeJob += jobVector[i] * jobVector[i]
	}

	magnitudeCandidate = math.Sqrt(magnitudeCandidate)
	magnitudeJob = math.Sqrt(magnitudeJob)

	if magnitudeCandidate == 0 || magnitudeJob == 0 {
		return 0
	}

	return dotProduct / (magnitudeCandidate * magnitudeJob)
}
