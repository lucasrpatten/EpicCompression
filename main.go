package main

import (
	"github.com/lucasrpatten/compression/fileutils"
	"github.com/lucasrpatten/compression/stageone"
)

func main() {
	file := string(fileutils.ReadFile("enwik9"))
	fileutils.CreateFile("./output.txt")
	stageone.StageOne(file)
}
