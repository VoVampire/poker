package main

import (
	"flag"
	"fmt"
	"poker/src"
	"time"
)

var filePath = flag.String("file", "./match_samples/seven_cards_with_ghost.result.json", "file path")

func main() {
	flag.Parse()

	matches := src.MustGetMatchesFromMatchSamples(*filePath)

	beginTime := time.Now()

	matches.PrintCompareResult()

	finishTime := time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}
