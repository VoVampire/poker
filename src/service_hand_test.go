package src

import (
	"testing"
)

func TestCompare(t *testing.T) {

	if maxHand := analyzeHandsStr("Qs9dAh6hAc6d9c").getMaxHands(); maxHand.MaxCase != TwoPair {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", TwoPair, maxHand.MaxCase)
	} else if maxHand.MaxHand != 34608256 {
		t.Errorf("Value 错误，应该为\n%b，结果为 \n%b", 34608256, maxHand.MaxHand)
	}

	if maxHand := analyzeHandsStr("Jd5d9cXnKs2c6d").getMaxHands(); maxHand.MaxCase != OnePair {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", OnePair, maxHand.MaxCase)
	} else if maxHand.MaxHand != 16779920 {
		t.Errorf("Value 错误，应该为\n%b，结果为 \n%b", 16779920, maxHand.MaxHand)
	}

	if maxHand := analyzeHandsStr("9cXnKs2c6dTc7c").getMaxHands(); maxHand.MaxCase != Flush {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", Flush, maxHand.MaxCase)
	} else if maxHand.MaxHand != 4513 {
		t.Errorf("Value 错误，应该为\n%b，结果为 \n%b", 4513, maxHand.MaxHand)
	}

	if maxHand := analyzeHandsStr("Ad3hQhJc2c8s2s").getMaxHands(); maxHand.MaxCase != OnePair { // 一对 11011000000001
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", OnePair, maxHand.MaxCase)
	} else if maxHand.MaxHand != 13825 {
		t.Errorf("Value 错误，应该为\n%b，结果为 \n%b", 13825, maxHand.MaxHand)
	}

	// 葫蘆 100000000000010000001000001000000100000
	if maxHand := analyzeHandsStr("6dXnAc3c7dAh7s").getMaxHands(); maxHand.MaxCase != FullHouse {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", FullHouse, maxHand.MaxCase)
	} else if maxHand.MaxHand != 274911727648 {
		t.Errorf("Value 错误，应该为\n%b，结果为 \n%b", 274911727648, maxHand.MaxHand)
	}

	// 四條 0001000000000000100000000000010000000000001000010000
	if maxHand := analyzeHandsStr("JdJcJsJh3d4d6d").getMaxHands(); maxHand.MaxCase != FourOfAKind {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", FourOfAKind, maxHand.MaxCase)
	} else if maxHand.MaxHand != 281509340643856 {
		t.Errorf("Value 错误，应该为\n%b，结果为 \n%b", 281509340643856, maxHand.MaxHand)
	}

	// 一对
	if maxHand := analyzeHandsStr("JdJc5hKh3d4d6d").getMaxHands(); maxHand.MaxCase != OnePair {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", OnePair, maxHand.MaxCase)
	} else if maxHand.MaxHand != 4196888 {
		t.Errorf("Value 错误，应该为 %d，结果为 %d", 4196888, maxHand.MaxHand)
	}

	// 同花顺 111110
	if maxHand := analyzeHandsStr("XnJc5dKh3d4d6d").getMaxHands(); maxHand.MaxCase != StraightFlush {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", StraightFlush, maxHand.MaxCase)
	} else if maxHand.MaxHand != 62 {
		t.Errorf("Value 错误，应该为 %d，结果为 %d", 62, maxHand.MaxHand)
	}

	// 同花 1000011110
	if maxHand := analyzeHandsStr("JdJc5dKh3d4d6d").getMaxHands(); maxHand.MaxCase != Flush {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", Flush, maxHand.MaxCase)
	} else if maxHand.MaxHand != 542 {
		t.Errorf("Value 错误，应该为 %d，结果为 %d", 542, maxHand.MaxHand)
	}

	// 两对 1010000000000101000010000
	if maxHand := analyzeHandsStr("JdJcKcKh3d4d6d").getMaxHands(); maxHand.MaxCase != TwoPair {
		t.Errorf("PokerType 错误，应该为 %d，结果为 %d", TwoPair, maxHand.MaxCase)
	} else if maxHand.MaxHand != 20974096 {
		t.Errorf("Value 错误，应该为 %d，结果为 %d", 20974096, maxHand.MaxHand)
	}
}

func TestFindStraight(t *testing.T) {
	if tempValue := findStraight(30, 1); tempValue != 62 {
		t.Errorf("Value 错误，应该为 %d，结果为 %d", 62, tempValue)
	}
}

func TestGetFirstOne(t *testing.T) {
	if tempValue := getFirstOne(22); tempValue != 16 {
		t.Errorf("Value 错误，应该为 %d，结果为 %d", 16, tempValue)
	}
}

func TestMoveAndAdd(t *testing.T) {
	if tempValue := leftMoveAndAdd(4, 3); tempValue != 268468228 {
		t.Errorf("Value 错误，应该为\n %64b，结果为\n %64b\n", 268468228, tempValue)
	}
}
