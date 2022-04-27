package service

import "fmt"

func IsMutant(dna []string) bool {
	var all []string
	all = addHorizontalAndVerticalCombinations(dna, all)
	all = addObliqueCombinations(dna, all)
	return findMutant(all)
}

func findMutant(allDnaChains []string) bool {
	mutantChains := 0
	for i := range allDnaChains {
		if allDnaChains[i] == "AAAA" || allDnaChains[i] == "TTTT" || allDnaChains[i] == "CCCC" || allDnaChains[i] == "GGGG" {
			mutantChains++
			if mutantChains == 2 {
				return true
			}
		}
	}
	return false
}

func addHorizontalAndVerticalCombinations(dna []string, allDnaChains []string) []string {
	for i := 0; i < 6; i++ {
		for j := 0; j < 3; j++ {
			allDnaChains = append(allDnaChains, dna[i][0:4])
			chain := fmt.Sprintf("%s%s%s%s", dna[j][i:i+1], dna[1+j][i:i+1], dna[2+j][i:i+1], dna[3+j][i:i+1])
			allDnaChains = append(allDnaChains, chain)
		}
	}

	return allDnaChains
}

func addObliqueCombinations(dna []string, allDnaChains []string) []string {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 0 || j == 0 {
				chain := fmt.Sprintf("%s%s%s%s", dna[j][i:i+1], dna[1+j][1+i:i+2], dna[2+j][2+i:i+3], dna[3+j][3+i:4+i])
				allDnaChains = append(allDnaChains, chain)
			}
		}
	}
	return allDnaChains
}
