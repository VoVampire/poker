package main

import (
	"fmt"
	"time"
	"poker/src"
)

func main() {
	timeBegin := time.Now()
	// 从样例中获取比赛牌组
	matches, err := src.GetMatchesFromMatchSamples("./match_samples/seven_cards_with_ghost.result.json")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	// 循环数组比较各组牌大小
	for _, v := range matches.Matches {
		playerA := v.PlayerA
		playerB := v.PlayerB

		if result := src.Compare(playerA, playerB); result != v.Result {
			fmt.Printf("error result:%s, %s ,%d, should be %d\n", playerA, playerB, result, v.Result)
		} else {
			// 打印执行结果
			fmt.Printf("%s, %s ,%d\n", playerA, playerB, result)
		}
	}
	timeEnd := time.Now()
	fmt.Printf("总计： %d 条, %.2f 毫秒\n", len(matches.Matches), timeEnd.Sub(timeBegin).Seconds()*1000)
}
