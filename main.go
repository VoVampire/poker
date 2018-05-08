package main

import (
	"fmt"
	"time"
	"poker/src"
)

func main() {
	beginTime := time.Now()

	for _, path := range src.MatchSamplesPaths {
		src.MustGetMatchesFromMatchSamples(path).PrintCompareResult()
	}

	finishTime := time.Now()
	fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}
