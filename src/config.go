package src

var MatchSamplesPaths = []string{
	"./match_samples/seven_cards_with_ghost.result.json",
	//"./match_samples/seven_cards_with_ghost.json",
	//"./match_samples/five_cards_with_ghost.json",
	//"./match_samples/match.json",
}

// 花色对应编号
var Suits = map[string]int{
	"s": 3,
	"h": 2,
	"d": 1,
	"c": 0,
}

// 牌面对应编号（对应bit位置）
var Faces = map[byte]uint64{
	'A': 1<<12,
	'K': 1<<11,
	'Q': 1<<10,
	'J': 1<<9,
	'T': 1<<8,
	'9': 1<<7,
	'8': 1<<6,
	'7': 1<<5,
	'6': 1<<4,
	'5': 1<<3,
	'4': 1<<2,
	'3': 1<<1,
	'2': 1<<0,
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
