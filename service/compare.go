package service

import (
	"poker/model"
	"poker/util"
)

func Compare(pokersOne string, pokersTwo string) int {
	// 初始化牌组对象
	pgOne := model.PokerGroup{}
	pgTwo := model.PokerGroup{}

	pgOne.NewPG(pokersOne)
	pgTwo.NewPG(pokersTwo)

	// 如果牌型（PokerGroup.PokerType）不同则比较牌型，否则比较权值（PokerGroup.Value）
	if result := getCompareResult(pgOne.PokerType, pgTwo.PokerType); result != 0 {
		return result

	} else {
		// 顺子&同花顺存在“A2345”这一特殊情况，此时为最小顺子，需要手动标记（权值value设为0）
		pgOneValue := util.If(pgOne.IsA2345(), 0, pgOne.Value).(uint64)
		pgTwoValue := util.If(pgTwo.IsA2345(), 0, pgTwo.Value).(uint64)

		return getCompareResult(pgOneValue, pgTwoValue)
	}
}

func getCompareResult(a, b uint64) int {
	if a > b {
		return 1
	} else if a < b {
		return 2
	} else {
		return 0
	}
}