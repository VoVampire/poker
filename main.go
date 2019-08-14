package main

import (
	"fmt"
	"poker/src"
	"time"
)

func main() {
	var matches *src.Matches

	for _, path := range src.MatchSamplesPaths {
		matches = src.MustGetMatchesFromMatchSamples(path)
	}

	beginTime := time.Now()

	matches.PrintCompareResult()

	finishTime := time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}
