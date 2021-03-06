package src

type Hand struct {
	HandStr  string    // 记录原始手牌字符串
	GhostNum uint64    // 鬼牌数量
	Suits    [4]uint64 // 记录手牌中出现过得所有牌的花色
	Faces    [4]uint64 // 记录手牌中出现过得所有牌的出现的次数（数组下标加1即为出现次数，bit位记录手牌牌面）
}

type MaxHand struct {
	MaxCase   uint64 // 记录最大牌型（StraightFlush, FourOfAKind, FullHouse...）
	MaxHand   uint64 // 记录最大五张牌和得分（bit位记录牌，int值表示得分）
	FlushFlag bool   // 记录是否存在同花牌型
	FlushSuit int    // 如果有同花，记录同花的花色编号
}

// 比较两张手牌、支持任意数量手牌及任意数量赖子
func Compare(strA string, strB string) int {
	playerA := analyzeHandStr(strA).getMaxHands()
	playerB := analyzeHandStr(strB).getMaxHands()

	// 比较最大牌型
	if winner := getWinner(playerA.MaxCase, playerB.MaxCase); winner != 0 {
		return winner
	}

	// 顺子&同花顺存在“A2345”这一特殊情况，此时为最小顺子，需要手动标记（权值score设为0）
	scoreA := If(playerA.MaxHand == A2345, uint64(0), playerA.MaxHand).(uint64)
	scoreB := If(playerB.MaxHand == A2345, uint64(0), playerB.MaxHand).(uint64)
	return getWinner(scoreA, scoreB)
}

// 获取获胜者编号
func getWinner(a, b uint64) int {
	return CaseWhen(a == b, 0, a > b, 1, a < b, 2).(int)
}

// 解析手牌字符串
func analyzeHandStr(handStr string) *Hand {
	var hand = Hand{HandStr: handStr}
	var faceValue uint64 // 面值
	for i := 0; i < len(handStr); i++ {
		if handStr[i] == 'X' { // 鬼牌（赖子）直接跳过当前面值和花色
			hand.GhostNum++
			i++
			continue
		}

		if i%2 == 0 {
			faceValue = Faces[handStr[i]]
			// 出现四次的相同面值的牌,更新对应bit位为1
			hand.Faces[3] |= hand.Faces[2] & faceValue
			// 出现三次的相同面值的牌,更新对应bit位为1
			hand.Faces[2] |= hand.Faces[1] & faceValue
			// 出现两次的相同面值的牌,更新对应bit位为1
			hand.Faces[1] |= hand.Faces[0] & faceValue
			// 出现一次的相同面值的牌,更新对应bit位为1
			hand.Faces[0] |= faceValue
		} else {
			// 记录花色
			hand.Suits[Suits[handStr[i]]] |= faceValue
		}
	}
	return &hand
}

// 获取最大手牌
func (hand *Hand) getMaxHands() *MaxHand {
	maxHand := MaxHand{}
	if maxHand.isStraightFlush(hand) {
	} else if maxHand.isFourOfAKind(hand) {
	} else if maxHand.isFullHouse(hand) {
	} else if maxHand.isFlush(hand) {
	} else if maxHand.isStraight(hand) {
	} else if maxHand.isThreeOfAKind(hand) {
	} else if maxHand.isTwoPair(hand) {
	} else if maxHand.isOnePair(hand) {
	} else if maxHand.isHighCard(hand) {
	}
	return &maxHand
}

// 筛选同花顺
func (maxHand *MaxHand) isStraightFlush(hand *Hand) bool {
	var tempValue uint64
	for i := 0; i < len(hand.Suits); i++ {
		// 筛选相同花色牌个数，如果大于（5-赖子）则标记为同花
		if cardNum := countOne(hand.Suits[i]); cardNum >= 5-hand.GhostNum {
			maxHand.FlushFlag = true
			maxHand.FlushSuit = i
			// 再用检查是否有顺子，若有则标记为同花顺
			if tempValue = findStraight(hand.Suits[i], hand.GhostNum); tempValue > 0 {
				if maxHand.MaxHand == 0 {
					maxHand.MaxHand = tempValue
				} else {
					maxHand.MaxHand = If(tempValue > maxHand.MaxHand && tempValue != A2345, tempValue, maxHand.MaxHand).(uint64)
				}
				maxHand.MaxCase = StraightFlush
			}
		}
	}
	return maxHand.MaxCase == StraightFlush
}

// 筛选四条 赖子最多三个，超过三个必为同花顺
func (maxHand *MaxHand) isFourOfAKind(hand *Hand) bool {
	if hand.Faces[3-hand.GhostNum] > 0 {
		maxHand.MaxCase = FourOfAKind
		switch {
		case hand.GhostNum == 0:
			maxHand.MaxHand = leftMoveAndAdd(hand.Faces[3], 4) | getFirstOne(hand.Faces[3]^hand.Faces[0])
		case hand.GhostNum > 0:
			{
				firstOne := getFirstOne(hand.Faces[3-hand.GhostNum])
				// if判断有无多余赖子，如果有多余赖子则直接补为A，否则取原单张牌中的最大牌
				maxHand.MaxHand = leftMoveAndAdd(firstOne, 4) |
					If(firstOne&hand.Faces[(3-hand.GhostNum)+1] > 0, A, getFirstOne(firstOne^hand.Faces[0])).(uint64)
			}
		}
		return true
	}
	return false
}

// 筛选葫芦 赖子最多一个，超过一个必大于等于四条
func (maxHand *MaxHand) isFullHouse(hand *Hand) bool {
	if hand.Faces[2-hand.GhostNum] > 0 && countOne(hand.Faces[1]) >= 2 {
		maxHand.MaxCase = FullHouse
		switch {
		case hand.GhostNum == 0:
			{
				firstOne := hand.Faces[2]
				secondOne := getFirstOne(hand.Faces[2] ^ hand.Faces[1])
				maxHand.MaxHand = leftMoveAndAdd(firstOne, 3) | leftMoveAndAdd(secondOne, 2)
			}
		case hand.GhostNum == 1: // 两对+1赖子 或 三对加1赖子
			{
				firstOne := getFirstOne(hand.Faces[1])
				secondOne := getFirstOne(hand.Faces[1] ^ firstOne)
				maxHand.MaxHand = leftMoveAndAdd(firstOne, 3) | leftMoveAndAdd(secondOne, 2)
			}
		}
		return true
	}
	return false
}

// 筛选同花 到这里赖子最多两个 剩下五张牌最多只能拼出一幅同花
func (maxHand *MaxHand) isFlush(hand *Hand) bool {
	if maxHand.FlushFlag {
		var tempValue uint64
		maxHand.MaxCase = Flush
		tempValue = (hand.Suits[maxHand.FlushSuit] & AKQJT) ^ AKQJT                  // 生成賴子可能放置的位置 例如 01110...
		tempValue = deleteLastOne(tempValue, int(countOne(tempValue)-hand.GhostNum)) // 确认賴子放置的位置 例如 01100...
		tempValue = hand.Suits[maxHand.FlushSuit] | tempValue                        // 拼接賴子
		maxHand.MaxHand = deleteLastOne(tempValue, int(countOne(tempValue)-5))       // 裁剪多余的1
		return true
	}
	return false
}

// 筛选顺子
func (maxHand *MaxHand) isStraight(hand *Hand) bool {
	if maxHand.MaxHand = findStraight(hand.Faces[0], hand.GhostNum); maxHand.MaxHand != 0 {
		maxHand.MaxCase = Straight
		return true
	}
	return false
}

// 筛选三对
func (maxHand *MaxHand) isThreeOfAKind(hand *Hand) bool {
	if hand.Faces[2-hand.GhostNum] > 0 {
		maxHand.MaxCase = ThreeOfAKind
		firstOne := getFirstOne(hand.Faces[2-hand.GhostNum])
		maxHand.MaxHand = leftMoveAndAdd(firstOne, 3) | deleteLastOne(hand.Faces[0]^firstOne, 2)
		return true
	}
	return false
}

// 筛选两对 不可能有赖子
func (maxHand *MaxHand) isTwoPair(hand *Hand) bool {
	if countOne := countOne(hand.Faces[1]); countOne >= 2 {
		var tempValue uint64
		maxHand.MaxCase = TwoPair
		tempValue = deleteLastOne(hand.Faces[1], int(countOne-2)) // 有可能有三对，剔除多余的对子
		maxHand.MaxHand = leftMoveAndAdd(tempValue, 2) | deleteLastOne(hand.Faces[0]^tempValue, int(4-countOne))
		return true
	}
	return false
}

// 筛选一对
func (maxHand *MaxHand) isOnePair(hand *Hand) bool {
	if hand.Faces[1-hand.GhostNum] > 0 {
		maxHand.MaxCase = OnePair
		switch {
		case hand.GhostNum == 0:
			maxHand.MaxHand = leftMoveAndAdd(hand.Faces[1], 2) | deleteLastOne(hand.Faces[0]^hand.Faces[1], 2)

		case hand.GhostNum == 1:
			maxHand.MaxHand = (getFirstOne(hand.Faces[0]) << 13) | deleteLastOne(hand.Faces[0], 2)
		}
		return true
	}
	return false
}

// 筛选高牌 到高牌则说明没有赖子，直接去掉两张最小牌即可
func (maxHand *MaxHand) isHighCard(hand *Hand) bool {
	maxHand.MaxCase = HighCard
	maxHand.MaxHand = deleteLastOne(hand.Faces[0], 2)
	return true
}

//****************************以下为工具代码**********************************

// 查找序列中可能存在的顺子，并返回牌面最大的一个
func findStraight(data uint64, superCardNum uint64) uint64 {
	var cardNum uint64
	var cardMold uint64

	// 定义模板模板,从最大顺子"AKQJT"开始依次与牌面做匹配,例:
	// cardface	0000011011111    0000011011111    		  0000011011111    0000011011111
	// cardMold 1111100000000 -> 0111110000000 -> ... ->  0000011111000 -> 0000000011111
	// superCard											(有1赖子情况)		(无赖子情况)

	cardMold = AKQJT
	for cardMold >= 31 {
		if cardNum = countOne(data & cardMold); cardNum >= 5-superCardNum {
			return cardMold
		}
		cardMold = cardMold >> 1
	}

	// 最后判断"A2345"这一特殊情况
	cardMold = A2345
	if cardNum = countOne(data & cardMold); cardNum >= 5-superCardNum {
		return cardMold
	}
	return 0
}

// 获取整形转二进制后最高位1的值 func(1011) -> 1000
func getFirstOne(data uint64) (result uint64) {
	for data > 0 {
		result = data
		data = data & (data - 1)
	}
	return
}

// 删除整形转二进制后最后n个1,并返回删除后的值 func(1011, 2) -> 1000
func deleteLastOne(data uint64, deleteOneNum int) uint64 {
	if deleteOneNum <= 0 {
		return data
	} else {
		deleteOneNum--
		return deleteLastOne(data&(data-1), deleteOneNum)
	}
}

// 将数值左移后累加 func(100,2) -> 100100  func(100,3) -> 100100100
func leftMoveAndAdd(data uint64, moveCount int) (result uint64) {
	for i := 0; i < moveCount; i++ {
		result |= data << uint(i*13)
	}
	return
}

// 统计二进制中1的个数（最大有效位数为16位）
func countOne(a uint64) uint64 {
	// 这里用了分治思想：先将相邻两个比特位１的个数相加，再将相邻四各比特位值相加...
	a = ((a & 0xAAAA) >> 1) + (a & 0x5555) // 1010101010101010  0101010101010101
	a = ((a & 0xCCCC) >> 2) + (a & 0x3333) // 1100110011001100  0011001100110011
	a = ((a & 0xF0F0) >> 4) + (a & 0x0F0F) // 1111000011110000  0000111100001111
	a = ((a & 0xFF00) >> 8) + (a & 0x00FF) // 1111111100000000  0000000011111111
	return a
}

// 三目表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// Case When Then
func CaseWhen(whenThen ...interface{}) interface{} {
	for i := 0; i < len(whenThen)-1; i += 2 {
		if whenThen[i].(bool) {
			return whenThen[i+1]
		}
	}
	return nil
}
