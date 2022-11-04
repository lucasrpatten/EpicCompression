package main

import (
	"github.com/lucasrpatten/compression/fileutils"
)

func main() {
	fileutils.ReadFile("enwik9")
	fileutils.CreateFile("./output.txt")
	StringByte(97)
}
