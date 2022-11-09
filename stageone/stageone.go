package stageone

import (
	"fmt"
	"strings"
)

func StageOne(s string) {
	splitstring := strings.Split(s, " ")
	for i := 0; i < len(splitstring); i++ {
		var word string =
		if splitstring[i] == "killed" {
			fmt.Print(splitstring[i])
		}
	}
}
