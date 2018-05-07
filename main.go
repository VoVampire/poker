package main

import (
	"fmt"
	"time"
	"poker/src"
	"strings"
)

var paths = []string{
	"./match_samples/seven_cards_with_ghost.result.json",
	"./match_samples/seven_cards_with_ghost.json",
	"./match_samples/five_cards_with_ghost.json",
	"./match_samples/match.json",
}

func main() {
	timeBegin := time.Now()

	for _, path := range paths {
		checkResult := strings.Contains(path, "result")
		src.MustGetMatchesFromMatchSamples(path).PrintCompareResult(checkResult)
	}

	timeEnd := time.Now()
	fmt.Printf("耗时：%.2f 毫秒\n", timeEnd.Sub(timeBegin).Seconds()*1000)
}
