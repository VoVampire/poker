package src

type Hand struct {
	HandStr  string    // 记录原始数据字符串 //例:AdAsAcAh2h2h2d
	GhostNum uint64    // 鬼牌数量
	Suits    [4]uint64 // Color[0]:记录所有黑桃's',Color[1]:记录所有红心'h',Color[2]:记录所有方块'd',Color[3]:记录所有草花'c'
	Faces    [4]uint64 // Face[0]:记录所有出现过的牌,Face[1]:记录至少出现过两次的牌,Face[2]:记录至少出现过三次的牌...
}

type MaxHand struct {
	MaxCase   uint64 // 记录最大牌型 例: StraightFlush, FourOfAKind, FullHouse...
	MaxHand   uint64 // 记录最大五张牌和得分（bit位记录牌，int值表示得分）
	FlushFlag bool   // 是否存在同花牌型
	FlushSuit int    // 如果有同花，则记录同花的花色编号
}

var Suits = map[string]int{
	"s": 3,
	"h": 2,
	"d": 1,
	"c": 0,
}

var Faces = map[string]uint64{
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

const (
	StraightFlush = 8 // 皇家同花顺&同花顺
	FourOfAKind   = 7 // 四条
	FullHouse     = 6 // 葫芦
	Flush         = 5 // 同花
	Straight      = 4 // 顺子
	ThreeOfAKind  = 3 // 三条
	TwoPair       = 2 // 两对
	OnePair       = 1 // 一对
	HighCard      = 0 // 散牌
)

const (
	// 特殊值        AKQJT98765432
	A2345 = 4111 // 1000000001111
	AKQJT = 7936 // 1111100000000
	A     = 4096 // 1000000000000
)
