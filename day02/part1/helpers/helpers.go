package helpers

import (
	"strings"
	"strconv"
)

type GameSamples struct {
	Id int
	Samples []map[string]int
}

func ParseSample(line string) GameSamples {
	idSamples := strings.Split(line, ": ")
	id, _ := strconv.Atoi(strings.Split(idSamples[0], " ")[1])
	samplesStrings := strings.Split(strings.Trim(idSamples[1], " "), ";")
	samples := make([]map[string]int, len(samplesStrings))
	for idx, sampleString := range(samplesStrings) {
		pieces := strings.Split(strings.Trim(sampleString, " "), ",")
		sample := make(map[string]int, len(pieces))
		for _, piece := range(pieces) {
			numColPieces := strings.Split(strings.Trim(piece, " "), " ")
			num, _ := strconv.Atoi(strings.Trim(numColPieces[0], " "))
			col := numColPieces[1]
			sample[col] = num
		}
		samples[idx] = sample
	}
	return GameSamples {
		id,
		samples,
	}
}

func SupportedBy (gs GameSamples, limits map[string]int) bool {
	for col, num := range(limits) {
		for _, sample := range(gs.Samples) {
			sc, ok := sample[col]
			if (ok && sc > num) {
				return false
			}
		}
	}
	return true
}