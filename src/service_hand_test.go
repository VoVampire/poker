package src

import (
	"testing"
)

func TestGetMaxHands(t *testing.T) {
	assertTemplet(t, "XnJc5dKh3d4d6d", StraightFlush, 62)
	assertTemplet(t, "JdJcJsJh3d4d6d", FourOfAKind, 281509340643856)
	assertTemplet(t, "6dXnAc3c7dAh7s", FullHouse, 274911727648)
	assertTemplet(t, "JdJc5dKh3d4d6d", Flush, 542)
	assertTemplet(t, "9cXnKs2c6dTc7c", Flush, 4513)
	assertTemplet(t, "JdJcKcKh3d4d6d", TwoPair, 20974096)
	assertTemplet(t, "Qs9dAh6hAc6d9c", TwoPair, 34608256)
	assertTemplet(t, "Ad3hQhJc2c8s2s", OnePair, 13825)
	assertTemplet(t, "JdJc5hKh3d4d6d", OnePair, 4196888)
	assertTemplet(t, "Jd5d9cXnKs2c6d", OnePair, 16779920)
}

func assertTemplet(t *testing.T, handStr string, expectCase, expectHand uint64) {
	if maxHand := analyzeHandStr(handStr).getMaxHands(); maxHand.MaxCase != expectCase {
		t.Errorf("MaxCase 错误，应该为 %d，结果为 %d", expectCase, maxHand.MaxCase)
	} else if maxHand.MaxHand != expectHand {
		t.Errorf("MaxHand 错误，应该为 %d，结果为 %d", expectHand, maxHand.MaxHand)
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
