package model

import (
	"poker/util"
)

const (
	// 牌型
	//RoyalFlush    = 8
	StraightFlush = 8
	FourOfAKind   = 7
	FullHouse     = 6
	Flush         = 5
	Straight      = 4
	ThreeOfAKind  = 3
	TwoPair       = 2
	OnePair       = 1
	HighCard      = 0

	// 特殊牌型权值// AKQJT98765432 AKQJT98765432 AKQJT98765432 AKQJT98765432
	A2345 = 4111 // 0000000000000 0000000000000 0000000000000 1000000001111

	Level1 = 1 << 0
	Level2 = 1 << 12
	Level3 = 1 << 25
	Level4 = 1 << 38
)

var Face = map[string]uint64{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

type PokerGroup struct {
	Pokers    string
	PokerType uint64
	Value     uint64
}

func (pg *PokerGroup) NewPG(pokers string) {
	color := pokers[1:2] // 记录初始花色
	pg.Pokers = pokers   // KsKhKdKc2c
	isFlush := true
	maxPoker := Face[pokers[0:1]]
	minPoker := Face[pokers[0:1]]

	// todo 遍历并记录牌信息
	var faceValue uint64
	for i := 0; i < len(pokers); i++ {
		if i%2 == 1 {
			// 对牌的花色进行处理
			if isFlush {
				if color != string(pokers[i]) {
					isFlush = false
				}
			}

		} else {
			// 对牌的面值进行处理
			faceValue = Face[string(pokers[i])]
			pg.setValue(faceValue)

			maxPoker = util.If(maxPoker > faceValue, maxPoker, faceValue).(uint64)
			minPoker = util.If(minPoker < faceValue, maxPoker, faceValue).(uint64)
		}
	}

	// todo 判断牌类型
	switch {
	case pg.Value > Level4: // 有四张相同只能是四条
		pg.PokerType = FourOfAKind
	case Level4 > pg.Value && pg.Value > Level3: // 有三张相同可能是葫芦或者三条
		{
			// 若共有三种不同牌则为三条，否则为葫芦（两张不同的牌）
			if util.CountOne(pg.Value&(1<<13)-1) == 3 {
				pg.PokerType = ThreeOfAKind
			} else {
				pg.PokerType = FullHouse
			}
		}
	case Level3 > pg.Value && pg.Value > Level2: // 有两张相同可能是两对或一对
		{
			// 若共有四种不同牌则为一对，否则为两对（三张不同的牌）
			if util.CountOne(pg.Value&(1<<13)-1) == 4 {
				pg.PokerType = OnePair
			} else {
				pg.PokerType = TwoPair
			}
		}
	case Level2 > pg.Value && pg.Value > Level1: // 没有相同牌可能是同花顺、同花、顺子、高牌
		{
			pg.PokerType = HighCard

			if maxPoker-minPoker == 4 || pg.Value == A2345 {
				pg.PokerType = Straight
			}

			if isFlush {
				if pg.PokerType != Straight {
					pg.PokerType = Flush
				} else {
					pg.PokerType = StraightFlush
				}
			}
		}
	}
}

func (pg *PokerGroup) setValue(i uint64) {
	if pg.Value&(1<<i) > 1 {
		pg.setValue(i + 13) // 递归
	} else {
		pg.Value = pg.Value | (1 << i)
	}
}

// 判断是否为“A2345”特殊牌型
func (pg *PokerGroup) IsA2345() bool {
	return (pg.PokerType == Straight || pg.PokerType == StraightFlush) && pg.Value == A2345
}
