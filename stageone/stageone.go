package stageone

import (
	"strings"
)

func StageOne(s string) {
	var words []string = strings.Split(s, " ")
	var histogram map[string]int = make(map[string]int)
	for _, str := range words {
		histogram[str]++
	}

	MapSort(histogram)

}

func MapSort(arr map[string]int) map[string]int {

}
